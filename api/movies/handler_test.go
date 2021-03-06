package movies

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	"github.com/darkjedidj/cinema-service/internal"
	movie "github.com/darkjedidj/cinema-service/internal/repository/movies"
	"github.com/darkjedidj/cinema-service/test"
)

func TestCreate(t *testing.T) {
	testCreateCases := []struct {
		name           string
		mockService    *test.MockService
		body           string
		expectedStatus int
	}{
		{
			name: "failure: empty body",
			mockService: &test.MockService{
				ExpectedResult: &movie.Resource{ID: 15, Name: "Harry Potter", Duration: "2h15m"},
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "success",
			mockService: &test.MockService{
				ExpectedResult: &movie.Resource{ID: 15, Name: "Harry Potter", Duration: "2h15m"},
			},
			body:           `{"Name": "Harry Potter", "Duration": "2h15m"}`,
			expectedStatus: http.StatusOK,
		},
		{
			name: "failure: DB error",
			mockService: &test.MockService{
				ExpectedError: internal.ErrInternalFailure,
			},
			body:           `{"Name": "Harry Potter", "Duration": "2h15m"}`,
			expectedStatus: http.StatusUnprocessableEntity,
		},
	}
	for _, tc := range testCreateCases {

		t.Run(tc.name, func(t *testing.T) {

			logger, err := zap.NewProduction()
			if err != nil {
				log.Fatalf("can't initialize zap logger: %v", err)
			}

			defer func() {
				if err := logger.Sync(); err != nil {
					fmt.Println(err)
				}
			}()

			w := httptest.NewRecorder()

			r := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(tc.body))
			r.Header.Set("Content-Type", "application/json")

			(&Handler{s: tc.mockService, log: logger}).Handle(w, r)

			assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
			assert.Equal(t, tc.expectedStatus, w.Code)

		})
	}
}

func TestRetrieve(t *testing.T) {
	testRetrieveCases := []struct {
		name           string
		mockService    *test.MockService
		id             int64
		expectedStatus int
	}{
		{
			name: "failure: no rows",
			mockService: &test.MockService{
				ExpectedResult: nil,
			},
			id:             20,
			expectedStatus: http.StatusNotFound,
		},
		{
			name: "success",
			mockService: &test.MockService{
				ExpectedResult: &movie.Resource{ID: 15, Name: "Harry Potter", Duration: "2h15m"},
			},
			id:             15,
			expectedStatus: http.StatusOK,
		},
		{
			name: "failure: DB error",
			mockService: &test.MockService{
				ExpectedError: internal.ErrInternalFailure,
			},
			id:             15,
			expectedStatus: http.StatusUnprocessableEntity,
		},
	}
	for _, tc := range testRetrieveCases {

		logger, err := zap.NewProduction()
		if err != nil {
			log.Fatalf("can't initialize zap logger: %v", err)
		}

		defer func() {
			if err := logger.Sync(); err != nil {
				fmt.Println(err)
			}
		}()

		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()

			vars := map[string]string{
				"id": fmt.Sprintf("%d", tc.id),
			}

			r := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/%d", tc.id), nil)

			r = mux.SetURLVars(r, vars)

			r.Header.Set("Content-Type", "application/json")

			(&Handler{s: tc.mockService}).HandleID(w, r)

			assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
			assert.Equal(t, tc.expectedStatus, w.Code)

		})
	}
}

func TestRetrieveAll(t *testing.T) {
	testRetrieveAllCases := []struct {
		name           string
		mockService    *test.MockService
		expectedStatus int
	}{
		{
			name: "failure: no rows",
			mockService: &test.MockService{
				ExpectedArray: nil,
			},
			expectedStatus: http.StatusNotFound,
		},
		{
			name: "success",
			mockService: &test.MockService{
				ExpectedArray: []internal.Identifiable{&movie.Resource{ID: 15, Name: "Harry Potter", Duration: "2h15m"}},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "failure: DB error",
			mockService: &test.MockService{
				ExpectedError: internal.ErrInternalFailure,
			},
			expectedStatus: http.StatusUnprocessableEntity,
		},
	}
	for _, tc := range testRetrieveAllCases {

		logger, err := zap.NewProduction()
		if err != nil {
			log.Fatalf("can't initialize zap logger: %v", err)
		}

		defer func() {
			if err := logger.Sync(); err != nil {
				fmt.Println(err)
			}
		}()

		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()

			r := httptest.NewRequest(http.MethodGet, "/", nil)
			r.Header.Set("Content-Type", "application/json")

			(&Handler{s: tc.mockService}).Handle(w, r)

			assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
			assert.Equal(t, tc.expectedStatus, w.Code)

		})
	}
}

func TestDelete(t *testing.T) {
	testDeleteCases := []struct {
		name           string
		mockService    *test.MockService
		id             int64
		expectedStatus int
		prepare        func() *zap.Logger
	}{
		{
			name: "success",
			mockService: &test.MockService{
				ExpectedResult: nil,
			},
			id:             15,
			expectedStatus: http.StatusOK,
		},
		{
			name: "failure: DB error",
			mockService: &test.MockService{
				ExpectedError: internal.ErrInternalFailure,
			},
			id:             15,
			expectedStatus: http.StatusUnprocessableEntity,
		},
	}
	for _, tc := range testDeleteCases {
		t.Run(tc.name, func(t *testing.T) {

			logger, err := zap.NewProduction()
			if err != nil {
				log.Fatalf("can't initialize zap logger: %v", err)
			}

			defer func() {
				if err := logger.Sync(); err != nil {
					fmt.Println(err)
				}
			}()

			w := httptest.NewRecorder()

			vars := map[string]string{
				"id": fmt.Sprintf("%d", tc.id),
			}

			r := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/%d", tc.id), nil)

			r = mux.SetURLVars(r, vars)

			r.Header.Set("Content-Type", "application/json")

			(&Handler{s: tc.mockService}).HandleID(w, r)

			assert.Equal(t, tc.expectedStatus, w.Code)

		})
	}
}
