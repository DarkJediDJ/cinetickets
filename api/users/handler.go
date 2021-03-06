package users

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"

	"github.com/darkjedidj/cinema-service/internal"
	repo "github.com/darkjedidj/cinema-service/internal/repository/users"
	user "github.com/darkjedidj/cinema-service/internal/service/user"
	e "github.com/darkjedidj/cinema-service/package"
	tkn "github.com/darkjedidj/cinema-service/package/jwt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

type Handler struct {
	s   user.Service // Allows use service features
	log *zap.Logger
}

func Init(db *sql.DB, l *zap.Logger) *Handler {

	service := user.Init(db, l)

	return &Handler{
		s:   *service,
		log: l,
	}
}

// CheckPrivileges of user
func (h *Handler) CheckPrivileges(route string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if len(header) == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			_, err := w.Write([]byte("Missing Authorization Header"))
			if err != nil {
				h.log.Info("Failed to write ticket response.",
					zap.Error(err),
				)

				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			return
		}

		header = strings.Replace(header, "Bearer ", "", 1)

		claims, err := tkn.VerifyToken(header)
		if err != nil {
			h.log.Info("Failed to verify token.",
				zap.Error(err),
			)
		}

		id := claims.(jwt.MapClaims)["ID"].(float64)
		privileges, err := h.s.RetrievePrivileges(int64(id))
		if err != nil {
			h.log.Info("Failed to get privileges.",
				zap.Error(err),
			)
		}

		for _, s := range privileges {
			if s == route {
				next(w, r)
			}
		}

		w.WriteHeader(http.StatusUnauthorized)
	}
}

// CheckTicket to download for user
func (h *Handler) CheckTicket(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)

		ticket, err := strconv.Atoi(vars["id"])
		if err != nil {
			h.log.Info("Failed to parse ticket id.",
				zap.Error(err),
			)

			w.WriteHeader(http.StatusBadRequest)
			return
		}

		header := r.Header.Get("Authorization")
		if len(header) == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			_, err := w.Write([]byte("Missing Authorization Header"))
			if err != nil {
				h.log.Info("Failed to write ticket response.",
					zap.Error(err),
				)

				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			return
		}

		header = strings.Replace(header, "Bearer ", "", 1)

		claims, err := tkn.VerifyToken(header)
		if err != nil {
			h.log.Info("Failed to verify token.",
				zap.Error(err),
			)
		}

		id := claims.(jwt.MapClaims)["ID"].(float64)

		privileges, err := h.s.RetrieveTickets(int64(ticket), int64(id))
		if err != nil {
			h.log.Info("Failed to get privileges.",
				zap.Error(err),
			)
		}

		if privileges {
			next(w, r)
		}

		w.WriteHeader(http.StatusUnauthorized)
	}
}

// Signup
// Signup godoc
// @Summary      Signup
// @Description  Signup
// @Tags         Users
// @Param        Body  body  repo.Resource  true  "Email and password"
// @Accept       json
// @Produce      json
// @Success      201
// @Failure      400
// @Failure      422
// @Failure      500
// @Router       /signup [post]
func (h *Handler) Signup(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var user repo.Resource

	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		h.log.Info("Failed to decode user json.",
			zap.Error(err),
		)

		return
	}

	user.Password = e.GetHash([]byte(user.Password))

	err = h.s.Create(&user, ctx)
	if err != nil {
		if errors.Is(err, internal.ErrWrongEmail) {
			response.WriteHeader(http.StatusUnprocessableEntity)

			_, err = response.Write([]byte(`{"message":"Email format inccorect"}`))
			if err != nil {
				h.log.Info("Failed to write user response.",
					zap.Error(err),
				)

				response.WriteHeader(http.StatusInternalServerError)
				return
			}
			return
		}
		if errors.Is(err, internal.ErrValidationFailed) {
			response.WriteHeader(http.StatusUnprocessableEntity)

			_, err = response.Write([]byte(`{"message":"This email is already in use"}`))
			if err != nil {
				h.log.Info("Failed to write user response.",
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

	response.WriteHeader(http.StatusCreated)
}

// Signin
// Signin godoc
// @Summary      Signin
// @Description  Signin
// @Tags         Users
// @Param        Body  body  repo.Resource  true  "Email and password"
// @Accept       json
// @Produce      json
// @Success      200  {object}  string
// @Failure      400
// @Failure      422
// @Failure      500
// @Router       /signin [post]
func (h *Handler) Signin(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var user repo.Resource

	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		h.log.Info("Failed to decode user json.",
			zap.Error(err),
		)
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	resource, err := h.s.Retrieve(user.EMail, ctx)
	if err != nil {
		if errors.Is(err, internal.ErrValidationFailed) {
			response.WriteHeader(http.StatusBadRequest)

			_, err = response.Write([]byte(err.Error()))
			if err != nil {
				h.log.Info("Failed to write response.",
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

	if resource == nil {
		response.WriteHeader(http.StatusUnauthorized)
		_, err = response.Write([]byte(`{"message":"Wrong email or password"}`))
		if err != nil {
			h.log.Info("Failed to write user response.",
				zap.Error(err),
			)

			response.WriteHeader(http.StatusInternalServerError)
			return
		}
		return
	}

	userDB, ok := resource.(*repo.Resource)
	if !ok {
		h.log.Info("Failed to assert user object.",
			zap.Bool("ok", ok),
		)

		response.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(userDB.Password), []byte(user.Password))
	if err != nil {
		response.WriteHeader(http.StatusUnauthorized)
		_, err = response.Write([]byte(`{"message":"Wrong email or password"}`))
		if err != nil {
			h.log.Info("Failed to write user response.",
				zap.Error(err),
			)

			response.WriteHeader(http.StatusInternalServerError)
			return
		}
		return
	}

	jwtToken, err := tkn.GenerateJWT(resource.GID())
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = response.Write([]byte(`{"token":"` + jwtToken + `"}`))
	if err != nil {
		h.log.Info("Failed to write user response.",
			zap.Error(err),
		)

		response.WriteHeader(http.StatusInternalServerError)
		return
	}
}
