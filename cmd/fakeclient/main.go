package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/mikolajsemeniuk/nord/pkg/scootin"
)

const baseURL = "http://localhost:8080/api/v1"
const apiKey = "1234"

// small envelope to keep scooters queryable around their start point
const maxDelta = 0.02 // ~2.2 km at equator
const stepMax = 0.001 // per tick step ~110m at equator
const tick = 3 * time.Second

func main() {
	for i := 0; i < 3; i++ {
		go fakeClient(i + 1)
	}

	select {}
}

func fakeClient(id int) {
	var (
		hasHome          bool
		homeLat, homeLng float64
	)

	// Główna pętla klienta
	for {
		var scooters []scootin.Scooter
		var err error

		if !hasHome {
			// first search: global wide bbox to find any free scooter
			scooters, err = findScootersGlobal("free")
		} else {
			// subsequent searches: tight bbox around the last known home
			scooters, err = findScootersIn(homeLat-maxDelta, homeLng-maxDelta, homeLat+maxDelta, homeLng+maxDelta, "free")
		}
		if err != nil {
			log.Printf("[client %d] find scooters error: %v", id, err)
			time.Sleep(2 * time.Second)
			continue
		}

		// pick first free
		var chosen scootin.Scooter
		for _, s := range scooters {
			if s.Status == "free" {
				chosen = s
				break
			}
		}

		if chosen.ID == uuid.Nil {
			log.Printf("[client %d] no free scooters, waiting...", id)
			time.Sleep(3 * time.Second)
			continue
		}

		// Nowa sekcja: Pętla ponawiania próby rezerwacji
	reserveLoop:
		for {
			// establish/refresh "home" around the chosen scooter position
			homeLat, homeLng = chosen.Latitude, chosen.Longitude
			hasHome = true

			log.Printf("[client %d] trying to reserve scooter %s (%s) at (%.5f, %.5f)",
				id, chosen.ID, chosen.City, chosen.Latitude, chosen.Longitude)

			// trip start
			err := updateScooter(scootin.UpdateScooterInput{
				ID:        chosen.ID,
				Status:    "occupied",
				Timestamp: time.Now().UTC(),
				Latitude:  chosen.Latitude,
				Longitude: chosen.Longitude,
			})

			// Obsługa błędu
			if err != nil {
				// Jeśli rezerwacja zakończyła się błędem 409 Conflict, spróbuj ponownie z inną hulajnogą
				if err.Error() == "update failed: 409 Conflict" {
					log.Printf("[client %d] conflict: scooter %s already occupied. Retrying...", id, chosen.ID)
					time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond) // Krótkie opóźnienie, aby uniknąć kolejnego konfliktu
					break reserveLoop                                            // Przerwij wewnętrzną pętlę i przejdź do następnej iteracji pętli głównej
				}
				log.Printf("[client %d] update scooter error: %v", id, err)
				time.Sleep(2 * time.Second)
				break reserveLoop // Przerwij wewnętrzną pętlę i przejdź do następnej iteracji pętli głównej
			}

			// Jeśli rezerwacja powiodła się, wyjdź z pętli rezerwacji
			log.Printf("[client %d] successfully reserved scooter %s", id, chosen.ID)
			break
		}

		// Jeśli udało się zarezerwować, kontynuuj jazdę
		if err == nil {
			// travel for 10–15s, update every 3s
			travelFor := time.Duration(10+rand.Intn(6)) * time.Second
			deadline := time.Now().Add(travelFor)

			for time.Now().Before(deadline) {
				// move with tiny random jitter, then clamp to envelope
				chosen.Latitude += randStep()
				chosen.Longitude += randStep()
				chosen.Latitude = clamp(chosen.Latitude, homeLat-maxDelta, homeLat+maxDelta)
				chosen.Longitude = clamp(chosen.Longitude, homeLng-maxDelta, homeLng+maxDelta)

				_ = updateScooter(scootin.UpdateScooterInput{
					ID:        chosen.ID,
					Status:    "occupied",
					Timestamp: time.Now().UTC(),
					Latitude:  chosen.Latitude,
					Longitude: chosen.Longitude,
				})

				time.Sleep(tick)
			}

			// trip end
			_ = updateScooter(scootin.UpdateScooterInput{
				ID:        chosen.ID,
				Status:    "free",
				Timestamp: time.Now().UTC(),
				Latitude:  chosen.Latitude,
				Longitude: chosen.Longitude,
			})
		}

		// rest 2–5s
		rest := time.Duration(2+rand.Intn(4)) * time.Second
		log.Printf("[client %d] resting for %v...", id, rest)
		time.Sleep(rest)
	}
}

func findScootersGlobal(status string) ([]scootin.Scooter, error) {
	return findScootersIn(-100, -100, 100, 100, status)
}

func findScootersIn(lat1, lng1, lat2, lng2 float64, status string) ([]scootin.Scooter, error) {
	loLat, hiLat := min(lat1, lat2), max(lat1, lat2)
	loLng, hiLng := min(lng1, lng2), max(lng1, lng2)

	url := fmt.Sprintf("%s/scooters?latitude1=%f&longitude1=%f&latitude2=%f&longitude2=%f", baseURL, loLat, loLng, hiLat, hiLng)
	if status != "" {
		url += "&status=" + status
	}

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("X-API-Key", apiKey)
	req.Header.Set("Accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode/100 != 2 {
		return nil, fmt.Errorf("unexpected status %s", res.Status)
	}

	var scooters []scootin.Scooter
	if err := json.NewDecoder(res.Body).Decode(&scooters); err != nil {
		return nil, err
	}
	return scooters, nil
}

func updateScooter(in scootin.UpdateScooterInput) error {
	body, _ := json.Marshal(in)
	url := fmt.Sprintf("%s/scooters/%s", baseURL, in.ID)

	req, _ := http.NewRequest("POST", url, bytes.NewReader(body))
	req.Header.Set("X-API-Key", apiKey)
	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode/100 != 2 {
		return fmt.Errorf("update failed: %s", res.Status)
	}

	return nil
}

func clamp(v, lo, hi float64) float64 {
	if v < lo {
		return lo
	}
	if v > hi {
		return hi
	}
	return v
}

func randStep() float64 {
	return (rand.Float64()*2 - 1) * stepMax
}
