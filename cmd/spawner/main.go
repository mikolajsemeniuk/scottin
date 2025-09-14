package main

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/kelseyhightower/envconfig"
	"github.com/mikolajsemeniuk/nord/pkg/scootin"
)

type Config struct {
	APIBaseURL   string        `default:"http://localhost:8080" envconfig:"API_BASE_URL"`
	NumClients   int           `default:"3"                     envconfig:"NUM_CLIENTS"`
	RideDuration time.Duration `default:"10s"                   envconfig:"RIDE_DURATION"`
	WaitDuration time.Duration `default:"5s"                    envconfig:"WAIT_DURATION"`
	MaxRetries   int           `default:"3"                     envconfig:"MAX_RETRIES"`
	RetryDelay   time.Duration `default:"1s"                    envconfig:"RETRY_DELAY"`
}

func randomFloat64() float64 {
	var buf [8]byte
	if _, err := rand.Read(buf[:]); err != nil {
		return 0
	}
	bits := binary.LittleEndian.Uint64(buf[:])

	return float64(int64(bits)%1000-500) / 1000.0 //nolint:gosec
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
	for range cfg.NumClients {
		wg.Add(1)
		go func() {
			defer wg.Done()

			client := &TrafficClient{
				clientID:   uuid.New().String(),
				apiBaseURL: cfg.APIBaseURL,
				httpClient: http.DefaultClient,
				logger:     log.New(log.Writer(), "", log.LstdFlags),
			}
			client.RunClient(ctx, cfg)
		}()
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

	url := tc.apiBaseURL + "/api/v1/scooters?latitude1=-100&longitude1=-100&latitude2=100&longitude2=100&status=free"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
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
		return nil, fmt.Errorf("%w: status %d: %s", scootin.ErrAPIRequestFailed, res.StatusCode, string(body))
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

	url := tc.apiBaseURL + "/api/v1/scooters/" + s.ID.String()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(data))
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
		return scootin.ErrScooterAlreadyReserved
	}

	if res.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(res.Body)
		return fmt.Errorf("%w: status %d: %s", scootin.ErrAPIRequestFailed, res.StatusCode, string(body))
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

	url := tc.apiBaseURL + "/api/v1/scooters/" + s.ID.String()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(data))
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
		return fmt.Errorf("%w: status %d: %s", scootin.ErrAPIRequestFailed, res.StatusCode, string(body))
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

	url := tc.apiBaseURL + "/api/v1/scooters/" + s.ID.String()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(data))
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
		return fmt.Errorf("%w: status %d: %s", scootin.ErrAPIRequestFailed, res.StatusCode, string(body))
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
		return scootin.ErrNoFreeScooters
	}

	var reserved *scootin.Scooter
	for i := range cfg.MaxRetries {
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
		return fmt.Errorf("%w: after %d attempts", scootin.ErrReservationFailed, cfg.MaxRetries)
	}

	tc.logger.Printf("%s: 🏍️  Riding scooter %s for %v...", tc.clientID, reserved.ID, cfg.RideDuration)

	currentLat := reserved.Latitude
	currentLng := reserved.Longitude
	rideStart := time.Now()

	for time.Since(rideStart) < cfg.RideDuration {
		latDelta := randomFloat64() * 0.001
		lngDelta := randomFloat64() * 0.001

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
