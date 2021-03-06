package sessions

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"go.uber.org/zap"

	"github.com/darkjedidj/cinema-service/internal"
	repo "github.com/darkjedidj/cinema-service/internal/repository/sessions"
	service "github.com/darkjedidj/cinema-service/internal/service/sessions"
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
		h.Get(response, request) // GET BASE_URL/v1/sessions/{id}
	case http.MethodDelete:
		h.Delete(response, request) // DELETE BASE_URL/v1/sessions/{id}
	default:
		response.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// Handle handles all endpoints on this route
func (h *Handler) Handle(response http.ResponseWriter, request *http.Request) {

	switch request.Method {
	case http.MethodGet:
		h.GetAll(response, request) // GET BASE_URL/v1/sessions
	default:
		response.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// Create get json and creates new session
// Create godoc
// @Security     ApiKeyAuth
// @Summary      Create session
// @Description  Creates session and returns created object
// @Tags         Sessions
// @Param        id    path  integer        true  "Session ID"
// @Param        Body  body  repo.Resource  true  "The body to create a session"
// @Accept       json
// @Produce      json
// @Success      200  {object}  repo.Resource
// @Failure      400
// @Failure      422
// @Failure      500
// @Failure      401
// @Router       /halls/{id}/sessions [post]
func (h *Handler) Create(response http.ResponseWriter, request *http.Request) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	vars := mux.Vars(request)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		h.log.Info("Failed to parse session id.",
			zap.Error(err),
		)

		response.WriteHeader(http.StatusBadRequest)
		return
	}

	var session repo.Resource

	response.Header().Set("Content-Type", "application/json")

	err = json.NewDecoder(request.Body).Decode(&session)
	if err != nil {
		h.log.Info("Failed to decode session json.",
			zap.Error(err),
		)

		response.WriteHeader(http.StatusBadRequest)
		return
	}

	session.Hall_id = int64(id)
	resource, err := h.s.Create(&session, ctx)
	if err != nil {
		if errors.Is(err, internal.ErrValidationFailed) {
			response.WriteHeader(http.StatusBadRequest)

			_, err = response.Write([]byte(err.Error()))
			if err != nil {
				h.log.Info("Failed to write session response.",
					zap.Error(err),
				)

				response.WriteHeader(http.StatusInternalServerError)
				return
			}
			return
		}

		response.WriteHeader(http.StatusUnprocessableEntity)

		return
	}

	body, err := json.Marshal(resource)
	if err != nil {
		h.log.Info("Failed to marshall session structure.",
			zap.Error(err),
		)

		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = response.Write(body)
	if err != nil {
		h.log.Info("Failed to write session response.",
			zap.Error(err),
		)

		response.WriteHeader(http.StatusInternalServerError)
		return
	}

}

// Delete get ID and deletes session with the same ID
// Delete godoc
// @Security     ApiKeyAuth
// @Summary      Delete session
// @Description  Deletes session
// @Param        id  path  integer  true  "Session ID"
// @Tags         Sessions
// @Accept       json
// @Produce      json
// @Success      200
// @Failure      400
// @Failure      422
// @Failure      500
// @Failure      401
// @Router       /sessions/{id} [delete]
func (h *Handler) Delete(response http.ResponseWriter, request *http.Request) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	vars := mux.Vars(request)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		h.log.Info("Failed to parse session id.",
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

// Get ID and selects session with the same ID
// Get godoc
// @Security     ApiKeyAuth
// @Summary      Get session
// @Description  Gets session
// @Param        id  path  integer  true  "Session ID"
// @Tags         Sessions
// @Accept       json
// @Produce      json
// @Success      200  {object}  repo.Resource
// @Failure      400
// @Failure      422
// @Failure      500
// @Failure      401
// @Router       /sessions/{id} [get]
func (h *Handler) Get(response http.ResponseWriter, request *http.Request) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	response.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(request)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		h.log.Info("Failed to parse session id.",
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
		h.log.Info("Failed to marshall session structure.",
			zap.Error(err),
		)

		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = response.Write(body)
	if err != nil {
		h.log.Info("Failed to write session response.",
			zap.Error(err),
		)

		response.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// GetAll selects all sessions
// GetAll godoc
// @Security     ApiKeyAuth
// @Summary      List session
// @Description  get sessions
// @Tags         Sessions
// @Accept       json
// @Produce      json
// @Success      200  {array}  []repo.Resource
// @Failure      400
// @Failure      422
// @Failure      500
// @Failure      401
// @Router       /sessions [get]
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
		h.log.Info("Failed to marshall session structure.",
			zap.Error(err),
		)

		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = response.Write(body)
	if err != nil {
		h.log.Info("Failed to write session response.",
			zap.Error(err),
		)

		response.WriteHeader(http.StatusInternalServerError)
		return
	}
}
