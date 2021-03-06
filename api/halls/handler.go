package halls

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"go.uber.org/zap"

	_ "github.com/darkjedidj/cinema-service/docs"
	"github.com/darkjedidj/cinema-service/internal"
	repo "github.com/darkjedidj/cinema-service/internal/repository/halls"
	service "github.com/darkjedidj/cinema-service/internal/service/halls"
)

type Handler struct {
	s   internal.Service // Allows use service features
	log *zap.Logger
}

func Init(db *sql.DB, l *zap.Logger) *Handler {

	service := service.Init(db, l)

	return &Handler{
		s:   service,
		log: l,
	}
}

// HandleID handles all endpoints on this route
func (h *Handler) HandleID(response http.ResponseWriter, request *http.Request) {

	switch request.Method {
	case http.MethodGet:
		h.Get(response, request) // GET BASE_URL/v1/halls/{id}
	case http.MethodDelete:
		h.Delete(response, request) // DELETE BASE_URL/v1/halls/{id}
	default:
		response.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// Handle handles all endpoints on this route
func (h *Handler) Handle(response http.ResponseWriter, request *http.Request) {

	switch request.Method {
	case http.MethodGet:
		h.GetAll(response, request) // GET BASE_URL/v1/halls
	case http.MethodPost:
		h.Create(response, request) // GET BASE_URL/v1/halls
	default:
		response.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// Create get json and creates new Hall
// Create godoc
// @Security     ApiKeyAuth
// @Summary      Create hall
// @Description  Creates hall and returns created object
// @Tags         Halls
// @Param        Body  body  repo.Resource  true  "The body to create a hall"
// @Accept       json
// @Produce      json
// @Success      200  {object}  repo.Resource
// @Failure      400
// @Failure      422
// @Failure      500
// @Failure      401
// @Router       /halls [post]
func (h *Handler) Create(response http.ResponseWriter, request *http.Request) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var hall repo.Resource

	response.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(request.Body).Decode(&hall)
	if err != nil {
		h.log.Info("Failed to decode hall json.",
			zap.Error(err),
		)

		response.WriteHeader(http.StatusBadRequest)
		return
	}

	resource, err := h.s.Create(&hall, ctx)
	if err != nil {
		response.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	body, err := json.Marshal(resource)
	if err != nil {
		h.log.Info("Failed to marshall hall structure.",
			zap.Error(err),
		)

		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = response.Write(body)
	if err != nil {
		h.log.Info("Failed to write hall response.",
			zap.Error(err),
		)

		response.WriteHeader(http.StatusInternalServerError)
		return
	}

}

// Delete get ID and deletes Hall with the same ID
// Delete godoc
// @Security     ApiKeyAuth
// @Summary      Delete hall
// @Description  Deletes hall
// @Param        id  path  integer  true  "Hall ID"
// @Tags         Halls
// @Accept       json
// @Produce      json
// @Success      200
// @Failure      400
// @Failure      422
// @Failure      500
// @Failure      401
// @Router       /halls/{id} [delete]
func (h *Handler) Delete(response http.ResponseWriter, request *http.Request) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	vars := mux.Vars(request)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		h.log.Info("Failed to parse hall id.",
			zap.Error(err),
		)

		response.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.s.Delete(int64(id), ctx)
	if err != nil {
		response.WriteHeader(http.StatusUnprocessableEntity)
	}

	response.WriteHeader(http.StatusOK)
}

// Get ID and selects Hall with the same ID
// Get godoc
// @Security     ApiKeyAuth
// @Summary      Get hall
// @Description  Gets hall
// @Param        id  path  integer  true  "Hall ID"
// @Tags         Halls
// @Accept       json
// @Produce      json
// @Success      200  {object}  repo.Resource
// @Failure      400
// @Failure      422
// @Failure      500
// @Failure      401
// @Router       /halls/{id} [get]
func (h *Handler) Get(response http.ResponseWriter, request *http.Request) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	response.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(request)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		h.log.Info("Failed to parse hall id.",
			zap.Error(err),
		)

		response.WriteHeader(http.StatusBadRequest)
		return
	}

	resource, err := h.s.Retrieve(int64(id), ctx)
	if err != nil {
		response.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	if resource == nil {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	body, err := json.Marshal(resource)
	if err != nil {
		h.log.Info("Failed to marshall hall structure.",
			zap.Error(err),
		)

		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = response.Write(body)
	if err != nil {
		h.log.Info("Failed to write hall response.",
			zap.Error(err),
		)

		response.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// GetAll selects all Halls
// GetAll godoc
// @Security     ApiKeyAuth
// @Summary      List halls
// @Description  get halls
// @Tags         Halls
// @Accept       json
// @Produce      json
// @Success      200  {array}  []repo.Resource
// @Failure      400
// @Failure      422
// @Failure      500
// @Failure      401
// @Router       /halls [get]
func (h *Handler) GetAll(response http.ResponseWriter, request *http.Request) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	response.Header().Set("Content-Type", "application/json")

	resource, err := h.s.RetrieveAll(ctx)
	if err != nil {
		response.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	if resource == nil {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	body, err := json.Marshal(resource)
	if err != nil {
		h.log.Info("Failed to marshall hall structure.",
			zap.Error(err),
		)

		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = response.Write(body)
	if err != nil {
		h.log.Info("Failed to write hall response.",
			zap.Error(err),
		)

		response.WriteHeader(http.StatusInternalServerError)
		return
	}
}
