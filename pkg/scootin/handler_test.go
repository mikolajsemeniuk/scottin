package scootin

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type mockStorage struct {
	findScooters  func(ctx context.Context, lat1, lng1, lat2, lng2 float64, status string) ([]Scooter, error)
	findScooter   func(ctx context.Context, id uuid.UUID) (*Scooter, error)
	updateScooter func(ctx context.Context, scooter Scooter) error
	close         func() error
}

func (m *mockStorage) FindScooters(ctx context.Context, lat1, lng1, lat2, lng2 float64, status string) ([]Scooter, error) {
	return m.findScooters(ctx, lat1, lng1, lat2, lng2, status)
}

func (m *mockStorage) FindScooter(ctx context.Context, id uuid.UUID) (*Scooter, error) {
	return m.findScooter(ctx, id)
}

func (m *mockStorage) UpdateScooter(ctx context.Context, scooter Scooter) error {
	return m.updateScooter(ctx, scooter)
}

func (m *mockStorage) Close() error {
	return m.close()
}

func TestFindScooters(t *testing.T) {
	t.Parallel()

	serve := func(svc *mockStorage, lat1, lng1, lat2, lng2 float64, status string) *http.Response {
		handler, err := NewHTTPHandler(svc)
		require.NoError(t, err)

		url := "/api/v1/scooters"
		if lat1 != 0 || lng1 != 0 || lat2 != 0 || lng2 != 0 {
			url += "?latitude1=-100&longitude1=-100&latitude2=100&longitude2=100"
			if status != "" {
				url += "&status=" + status
			}
		}

		r := httptest.NewRequest(http.MethodGet, url, nil)
		w := httptest.NewRecorder()
		handler.ServeHTTP(w, r)

		return w.Result()
	}

	t.Run("InvalidQueryParameters", func(t *testing.T) {
		t.Parallel()

		svc := &mockStorage{}

		res := serve(svc, 0, 0, 0, 0, "")
		defer res.Body.Close()
		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})

	t.Run("StorageFailure", func(t *testing.T) {
		t.Parallel()

		svc := &mockStorage{
			findScooters: func(_ context.Context, _, _, _, _ float64, _ string) ([]Scooter, error) {
				return nil, ErrScooterNotFound
			},
		}

		res := serve(svc, -100, -100, 100, 100, "")
		defer res.Body.Close()
		assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	})

	t.Run("Success", func(t *testing.T) {
		t.Parallel()

		expectedScooters := []Scooter{
			{
				ID:        uuid.New(),
				City:      "Ottawa",
				Status:    "free",
				Latitude:  45.4215,
				Longitude: -75.6972,
				ClientID:  "",
			},
		}

		svc := &mockStorage{
			findScooters: func(_ context.Context, _, _, _, _ float64, _ string) ([]Scooter, error) {
				return expectedScooters, nil
			},
		}

		res := serve(svc, -100, -100, 100, 100, "")
		defer res.Body.Close()
		assert.Equal(t, http.StatusOK, res.StatusCode)

		var scooters []Scooter
		err := json.NewDecoder(res.Body).Decode(&scooters)
		require.NoError(t, err)
		assert.Len(t, scooters, 1)
		assert.Equal(t, expectedScooters[0].ID, scooters[0].ID)
	})

	t.Run("SuccessWithStatusFilter", func(t *testing.T) {
		t.Parallel()

		expectedScooters := []Scooter{
			{
				ID:        uuid.New(),
				City:      "Montreal",
				Status:    "occupied",
				Latitude:  45.5017,
				Longitude: -73.5673,
				ClientID:  "client-123",
			},
		}

		svc := &mockStorage{
			findScooters: func(_ context.Context, _, _, _, _ float64, status string) ([]Scooter, error) {
				assert.Equal(t, "occupied", status)
				return expectedScooters, nil
			},
		}

		res := serve(svc, -100, -100, 100, 100, "occupied")
		defer res.Body.Close()
		assert.Equal(t, http.StatusOK, res.StatusCode)

		var scooters []Scooter
		err := json.NewDecoder(res.Body).Decode(&scooters)
		require.NoError(t, err)
		assert.Len(t, scooters, 1)
		assert.Equal(t, expectedScooters[0].Status, scooters[0].Status)
	})
}

func TestUpdateScooter(t *testing.T) { //nolint:funlen
	t.Parallel()

	serve := func(svc *mockStorage, payload UpdateScooterInput) *http.Response {
		handler, err := NewHTTPHandler(svc)
		require.NoError(t, err)

		body, err := json.Marshal(payload)
		require.NoError(t, err)

		url := "/api/v1/scooters/" + payload.ID.String()
		r := httptest.NewRequest(http.MethodPost, url, bytes.NewReader(body))
		r.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		handler.ServeHTTP(w, r)

		return w.Result()
	}

	t.Run("InvalidScooterID", func(t *testing.T) {
		t.Parallel()

		svc := &mockStorage{}

		payload := UpdateScooterInput{
			Status: StatusFree,
		}

		handler, err := NewHTTPHandler(svc)
		require.NoError(t, err)

		body, err := json.Marshal(payload)
		require.NoError(t, err)

		r := httptest.NewRequest(http.MethodPost, "/api/v1/scooters/invalid-uuid", bytes.NewReader(body))
		r.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		handler.ServeHTTP(w, r)

		res := w.Result()
		defer res.Body.Close()
		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})

	t.Run("InvalidJSON", func(t *testing.T) {
		t.Parallel()

		svc := &mockStorage{}

		handler, err := NewHTTPHandler(svc)
		require.NoError(t, err)

		r := httptest.NewRequest(http.MethodPost, "/api/v1/scooters/"+uuid.New().String(), bytes.NewReader([]byte("invalid json")))
		r.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		handler.ServeHTTP(w, r)

		res := w.Result()
		defer res.Body.Close()
		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})

	t.Run("InvalidStatus", func(t *testing.T) {
		t.Parallel()

		svc := &mockStorage{}

		payload := UpdateScooterInput{
			ID:     uuid.New(),
			Status: "invalid",
		}

		res := serve(svc, payload)
		defer res.Body.Close()
		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})

	t.Run("ScooterNotFound", func(t *testing.T) {
		t.Parallel()

		svc := &mockStorage{
			updateScooter: func(_ context.Context, _ Scooter) error {
				return ErrScooterNotFound
			},
		}

		payload := UpdateScooterInput{
			ID:        uuid.New(),
			Status:    StatusOccupied,
			Latitude:  45.4215,
			Longitude: -75.6972,
			ClientID:  "client-123",
		}

		res := serve(svc, payload)
		defer res.Body.Close()
		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	})

	t.Run("ClientIDRequired", func(t *testing.T) {
		t.Parallel()

		svc := &mockStorage{
			updateScooter: func(_ context.Context, _ Scooter) error {
				return ErrClientIDRequired
			},
		}

		payload := UpdateScooterInput{
			ID:        uuid.New(),
			Status:    StatusOccupied,
			Latitude:  45.4215,
			Longitude: -75.6972,
			ClientID:  "",
		}

		res := serve(svc, payload)
		defer res.Body.Close()
		assert.Equal(t, http.StatusConflict, res.StatusCode)
	})

	t.Run("ScooterAlreadyOccupied", func(t *testing.T) {
		t.Parallel()

		svc := &mockStorage{
			updateScooter: func(_ context.Context, _ Scooter) error {
				return ErrScooterAlreadyOccupied
			},
		}

		payload := UpdateScooterInput{
			ID:        uuid.New(),
			Status:    StatusOccupied,
			Latitude:  45.4215,
			Longitude: -75.6972,
			ClientID:  "client-123",
		}

		res := serve(svc, payload)
		defer res.Body.Close()
		assert.Equal(t, http.StatusConflict, res.StatusCode)
	})

	t.Run("OnlyOccupyingClientCanRelease", func(t *testing.T) {
		t.Parallel()

		svc := &mockStorage{
			updateScooter: func(_ context.Context, _ Scooter) error {
				return ErrOnlyOccupyingClientCanRelease
			},
		}

		payload := UpdateScooterInput{
			ID:        uuid.New(),
			Status:    StatusFree,
			Latitude:  45.4215,
			Longitude: -75.6972,
			ClientID:  "wrong-client",
		}

		res := serve(svc, payload)
		defer res.Body.Close()
		assert.Equal(t, http.StatusConflict, res.StatusCode)
	})

	t.Run("OnlyOccupyingClientCanUpdate", func(t *testing.T) {
		t.Parallel()

		svc := &mockStorage{
			updateScooter: func(_ context.Context, _ Scooter) error {
				return ErrOnlyOccupyingClientCanUpdate
			},
		}

		payload := UpdateScooterInput{
			ID:        uuid.New(),
			Status:    StatusOccupied,
			Latitude:  45.4215,
			Longitude: -75.6972,
			ClientID:  "wrong-client",
		}

		res := serve(svc, payload)
		defer res.Body.Close()
		assert.Equal(t, http.StatusConflict, res.StatusCode)
	})

	t.Run("StorageFailure", func(t *testing.T) {
		t.Parallel()

		svc := &mockStorage{
			updateScooter: func(_ context.Context, _ Scooter) error {
				return ErrScooterStatusChanged
			},
		}

		payload := UpdateScooterInput{
			ID:        uuid.New(),
			Status:    StatusOccupied,
			Latitude:  45.4215,
			Longitude: -75.6972,
			ClientID:  "client-123",
		}

		res := serve(svc, payload)
		defer res.Body.Close()
		assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	})

	t.Run("SuccessOccupied", func(t *testing.T) {
		t.Parallel()

		svc := &mockStorage{
			updateScooter: func(_ context.Context, scooter Scooter) error {
				assert.Equal(t, string(StatusOccupied), scooter.Status)
				assert.Equal(t, "client-123", scooter.ClientID)
				assert.Equal(t, 45.4215, scooter.Latitude)
				assert.Equal(t, -75.6972, scooter.Longitude)
				return nil
			},
		}

		payload := UpdateScooterInput{
			ID:        uuid.New(),
			Status:    StatusOccupied,
			Latitude:  45.4215,
			Longitude: -75.6972,
			ClientID:  "client-123",
		}

		res := serve(svc, payload)
		defer res.Body.Close()
		assert.Equal(t, http.StatusOK, res.StatusCode)
	})

	t.Run("SuccessFree", func(t *testing.T) {
		t.Parallel()

		svc := &mockStorage{
			updateScooter: func(_ context.Context, scooter Scooter) error {
				assert.Equal(t, string(StatusFree), scooter.Status)
				assert.Equal(t, "client-123", scooter.ClientID)
				assert.Equal(t, 45.4215, scooter.Latitude)
				assert.Equal(t, -75.6972, scooter.Longitude)
				return nil
			},
		}

		payload := UpdateScooterInput{
			ID:        uuid.New(),
			Status:    StatusFree,
			Latitude:  45.4215,
			Longitude: -75.6972,
			ClientID:  "client-123",
		}

		res := serve(svc, payload)
		defer res.Body.Close()
		assert.Equal(t, http.StatusOK, res.StatusCode)
	})
}
