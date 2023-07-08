package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth"
	"github.com/gsouza97/go-expert-api/internal/dto"
	"github.com/gsouza97/go-expert-api/internal/entity"
	"github.com/gsouza97/go-expert-api/internal/infra/database"
)

type UserHandler struct {
	UserDB       database.UserDBInterface
	Jwt          *jwtauth.JWTAuth // Esse JWT já faz parte do middleware do próprio chi
	JwtExpiresIn int              // Tempo de expiração do token
}

type Error struct {
	Message string `json:"message"`
}

func NewUserHandler(userDB database.UserDBInterface, jwt *jwtauth.JWTAuth, JWTExpiresIn int) *UserHandler {
	return &UserHandler{
		UserDB:       userDB,
		Jwt:          jwt,
		JwtExpiresIn: JWTExpiresIn,
	}
}

// Get user JWT godoc
// @Summary      Get user JWT
// @Description  Get user JWT
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        request    body     dto.GetJWTInput  true  "user credentials"
// @Success      200  {object}  dto.GetJWTOutput
// @Failure      401  {object}  Error
// @Failure      500  {object}  Error
// @Router       /users/getToken [post]
func (h *UserHandler) GetJWT(w http.ResponseWriter, r *http.Request) {
	var user dto.GetJWTInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil || user.Email == "" || user.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u, err := h.UserDB.FindByEmail(user.Email)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if !u.ValidatePassword(user.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	m := map[string]interface{}{"sub": u.ID.String(), "exp": time.Now().Add(time.Hour * time.Duration(h.JwtExpiresIn)).Unix()}
	_, tokenString, _ := h.Jwt.Encode(m)

	accessToken := dto.GetJWTOutput{AccessToken: tokenString}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)
}

// Create User godoc
// @Summary      Create user
// @Description  Create user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        request    body     dto.CreateUserInput  true  "user request"
// @Success      201
// @Failure      500  {object}  Error
// @Router       /users [post]
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil || user.Name == "" || user.Email == "" || user.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u, err := entity.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.UserDB.CreateUser(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
