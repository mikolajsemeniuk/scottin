package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/kelseyhightower/envconfig"
	"github.com/mikolajsemeniuk/nord/pkg/scootin"
)

type Config struct {
	APIBaseURL   string        `envconfig:"API_BASE_URL"  default:"http://localhost:8080"`
	NumClients   int           `envconfig:"NUM_CLIENTS"   default:"3"`
	RideDuration time.Duration `envconfig:"RIDE_DURATION" default:"10s"`
	WaitDuration time.Duration `envconfig:"WAIT_DURATION" default:"5s"`
	MaxRetries   int           `envconfig:"MAX_RETRIES"   default:"3"`
	RetryDelay   time.Duration `envconfig:"RETRY_DELAY"   default:"1s"`
}

func main() {
	var cfg Config
	if err := envconfig.Process("TRAFFIC", &cfg); err != nil {
		log.Fatalf("Failed to process config: %v", err)
	}

	log.Printf("🚦 Traffic simulation starting with %d clients", cfg.NumClients)
	log.Printf("📊 Configuration:")
	log.Printf("   API Base URL: %s", cfg.APIBaseURL)
	log.Printf("   Ride Duration: %v", cfg.RideDuration)
	log.Printf("   Wait Duration: %v", cfg.WaitDuration)
	log.Printf("   Max Retries: %d", cfg.MaxRetries)
	log.Printf("   Retry Delay: %v", cfg.RetryDelay)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	for i := 0; i < cfg.NumClients; i++ {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()

			client := &TrafficClient{
				clientID:   uuid.New().String(),
				apiBaseURL: cfg.APIBaseURL,
				httpClient: http.DefaultClient,
				logger:     log.New(log.Writer(), "", log.LstdFlags),
			}
			client.RunClient(ctx, cfg)
		}(i)
	}

	wg.Wait()
	log.Printf("🏁 All clients finished")
}

type TrafficClient struct {
	clientID   string
	apiBaseURL string
	httpClient *http.Client
	logger     *log.Logger
}

func (tc *TrafficClient) FindFreeScooters(ctx context.Context) ([]scootin.Scooter, error) {
	tc.logger.Printf("%s: 🔍 Searching for free scooters...", tc.clientID)

	url := fmt.Sprintf("%s/api/v1/scooters?latitude1=-100&longitude1=-100&latitude2=100&longitude2=100&status=free", tc.apiBaseURL)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	res, err := tc.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(res.Body)
		return nil, fmt.Errorf("API returned status %d: %s", res.StatusCode, string(body))
	}

	var scooters []scootin.Scooter
	if err := json.NewDecoder(res.Body).Decode(&scooters); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	tc.logger.Printf("%s: ✅ Found %d free scooters", tc.clientID, len(scooters))
	return scooters, nil
}

func (tc *TrafficClient) ReserveScooter(ctx context.Context, s scootin.Scooter) error {
	tc.logger.Printf("%s: 🚀 Attempting to reserve scooter %s", tc.clientID, s.ID)

	input := scootin.UpdateScooterInput{
		Status:    "occupied",
		Latitude:  s.Latitude,
		Longitude: s.Longitude,
		ClientID:  tc.clientID,
		Timestamp: time.Now().UTC(),
	}

	data, err := json.Marshal(input)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}

	url := fmt.Sprintf("%s/api/v1/scooters/%s", tc.apiBaseURL, s.ID)
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(data))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := tc.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to make request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusConflict {
		body, _ := io.ReadAll(res.Body)
		tc.logger.Printf("%s: ⚠️  Scooter %s is already reserved: %s", tc.clientID, s.ID, string(body))
		return fmt.Errorf("scooter already reserved")
	}

	if res.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(res.Body)
		return fmt.Errorf("API returned status %d: %s", res.StatusCode, string(body))
	}

	tc.logger.Printf("%s: ✅ Successfully reserved scooter %s", tc.clientID, s.ID)
	return nil
}

func (tc *TrafficClient) ReleaseScooter(ctx context.Context, s scootin.Scooter) error {
	tc.logger.Printf("%s: 🔓 Releasing scooter %s", tc.clientID, s.ID)

	input := scootin.UpdateScooterInput{
		Status:    "free",
		Latitude:  s.Latitude,
		Longitude: s.Longitude,
		ClientID:  tc.clientID,
		Timestamp: time.Now().UTC(),
	}

	data, err := json.Marshal(input)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}

	url := fmt.Sprintf("%s/api/v1/scooters/%s", tc.apiBaseURL, s.ID)
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(data))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := tc.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to make request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(res.Body)
		return fmt.Errorf("API returned status %d: %s", res.StatusCode, string(body))
	}

	tc.logger.Printf("%s: ✅ Successfully released scooter %s", tc.clientID, s.ID)
	return nil
}

func (tc *TrafficClient) UpdateScooterPosition(ctx context.Context, s scootin.Scooter, newLat, newLng float64) error {
	tc.logger.Printf("%s: 📍 Updating scooter %s position to (%.6f, %.6f)", tc.clientID, s.ID, newLat, newLng)

	input := scootin.UpdateScooterInput{
		Status:    "occupied",
		Latitude:  newLat,
		Longitude: newLng,
		ClientID:  tc.clientID,
		Timestamp: time.Now().UTC(),
	}

	data, err := json.Marshal(input)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}

	url := fmt.Sprintf("%s/api/v1/scooters/%s", tc.apiBaseURL, s.ID)
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(data))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := tc.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to make request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(res.Body)
		return fmt.Errorf("API returned status %d: %s", res.StatusCode, string(body))
	}

	return nil
}

func (tc *TrafficClient) SimulateRide(ctx context.Context, cfg Config) error {
	scooters, err := tc.FindFreeScooters(ctx)
	if err != nil {
		return fmt.Errorf("failed to find free scooters: %w", err)
	}

	if len(scooters) == 0 {
		tc.logger.Printf("%s: 😞 No free scooters available, waiting...", tc.clientID)
		return fmt.Errorf("no free scooters available")
	}

	var reserved *scootin.Scooter
	for i := 0; i < cfg.MaxRetries; i++ {
		scooter := scooters[i%len(scooters)]

		if err := tc.ReserveScooter(ctx, scooter); err == nil {
			reserved = &scooter
			break
		}

		tc.logger.Printf("%s: ⚠️  Reservation attempt %d failed, scooter %s is already occupied", tc.clientID, i+1, scooter.ID)
		if i < cfg.MaxRetries-1 {
			tc.logger.Printf("⏳ Retrying in %v...", cfg.RetryDelay)
			time.Sleep(cfg.RetryDelay)
		}
	}

	if reserved == nil {
		tc.logger.Printf("❌ Failed to reserve any scooter after %d attempts", cfg.MaxRetries)
		return fmt.Errorf("failed to reserve scooter after %d attempts", cfg.MaxRetries)
	}

	tc.logger.Printf("%s: 🏍️  Riding scooter %s for %v...", tc.clientID, reserved.ID, cfg.RideDuration)

	currentLat := reserved.Latitude
	currentLng := reserved.Longitude
	rideStart := time.Now()

	for time.Since(rideStart) < cfg.RideDuration {
		latDelta := (rand.Float64() - 0.5) * 0.001
		lngDelta := (rand.Float64() - 0.5) * 0.001

		currentLat += latDelta
		currentLng += lngDelta

		if err := tc.UpdateScooterPosition(ctx, *reserved, currentLat, currentLng); err != nil {
			tc.logger.Printf("%s: ⚠️  Failed to update position: %v", tc.clientID, err)
		}

		time.Sleep(3 * time.Second)
	}

	reserved.Latitude = currentLat
	reserved.Longitude = currentLng

	if err := tc.ReleaseScooter(ctx, *reserved); err != nil {
		return fmt.Errorf("failed to release scooter: %w", err)
	}

	tc.logger.Printf("%s: 🎉 Ride completed successfully!", tc.clientID)
	return nil
}

func (tc *TrafficClient) RunClient(ctx context.Context, cfg Config) {
	for {
		select {
		case <-ctx.Done():
			tc.logger.Printf("🛑 Client %s stopping...", tc.clientID)
			return
		default:
			if err := tc.SimulateRide(ctx, cfg); err != nil {
				tc.logger.Printf("%s: ❌ Ride simulation failed: %v", tc.clientID, err)
			}

			tc.logger.Printf("%s: 😴 Waiting %v before next ride...", tc.clientID, cfg.WaitDuration)
			time.Sleep(cfg.WaitDuration)
		}
	}
}
