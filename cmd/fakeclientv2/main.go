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

// Zmieniono stałe, aby zawsze szukać na szerokim obszarze
const minLat = -100.0
const maxLat = 100.0
const minLng = -100.0
const maxLng = 100.0
const tick = 3 * time.Second

func main() {
	for i := 0; i < 3; i++ {
		go fakeClient(i + 1)
	}

	select {}
}

func fakeClient(id int) {
	for {
		var scooters []scootin.Scooter
		var err error

		// Sekcja: Zawsze szukaj na szerokim, stałym obszarze
		scooters, err = findScootersIn(minLat, minLng, maxLat, maxLng, "free")

		if err != nil {
			log.Printf("[client %d] find scooters error: %v", id, err)
			time.Sleep(2 * time.Second)
			continue
		}

		// Wybierz losową hulajnogę z listy
		var chosen scootin.Scooter
		if len(scooters) > 0 {
			chosen = scooters[rand.Intn(len(scooters))]
		}

		if chosen.ID == uuid.Nil {
			log.Printf("[client %d] no free scooters, waiting...", id)
			time.Sleep(3 * time.Second)
			continue
		}

		// Sekcja: Próba rezerwacji
		for {
			log.Printf("[client %d] attempting to reserve scooter %s", id, chosen.ID)

			err := updateScooter(scootin.UpdateScooterInput{
				ID:        chosen.ID,
				Status:    "occupied",
				Timestamp: time.Now().UTC(),
				Latitude:  chosen.Latitude,
				Longitude: chosen.Longitude,
			})

			if err != nil {
				if err.Error() == "update failed: 409 Conflict" {
					log.Printf("[client %d] conflict: scooter %s already occupied. Trying again...", id, chosen.ID)
					time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
					break
				}
				log.Printf("[client %d] update scooter error: %v", id, err)
				time.Sleep(2 * time.Second)
				break
			}

			// Sekcja: Udany przejazd
			log.Printf("[client %d] successfully reserved scooter %s. Starting trip...", id, chosen.ID)

			// Brak ruchu, po prostu czekaj
			time.Sleep(tick)

			// Sekcja: Zakończenie przejazdu i oddanie hulajnogi
			log.Printf("[client %d] trip finished for scooter %s. Releasing...", id, chosen.ID)
			_ = updateScooter(scootin.UpdateScooterInput{
				ID:        chosen.ID,
				Status:    "free",
				Timestamp: time.Now().UTC(),
				Latitude:  chosen.Latitude,
				Longitude: chosen.Longitude,
			})

			// Po oddaniu hulajnogi, klient odpoczywa
			rest := time.Duration(2+rand.Intn(4)) * time.Second
			log.Printf("[client %d] resting for %v...", id, rest)
			time.Sleep(rest)
			break
		}
	}
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
