package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/CaiqueRibeiro/4-api/internal/dto"
	"github.com/CaiqueRibeiro/4-api/internal/entity"
	"github.com/CaiqueRibeiro/4-api/internal/infra/database"
	"github.com/go-chi/jwtauth"
)

type UserHandler struct {
	UserDB database.UserInterface
}

func NewUserHandler(db database.UserInterface) *UserHandler {
	return &UserHandler{
		UserDB: db,
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	u, err := entity.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.UserDB.Create(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func (h *UserHandler) GetJWT(w http.ResponseWriter, r *http.Request) {
	jwt := r.Context().Value("jwt").(*jwtauth.JWTAuth) // Pegando do middleware e fazendo cast do formato esperado
	jwtExpiresIn := r.Context().Value("jwtExpiresIn").(int)
	var dto dto.GetJWTInput
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	user, err := h.UserDB.FindByEmail(dto.Email)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	isValid := user.ValidatePassword(dto.Password)
	if !isValid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	_, token, _ := jwt.Encode(map[string]interface{}{
		"sub": user.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(jwtExpiresIn)).Unix(),
	})
	accessToken := struct {
		AccessToken string `json:"acess_token"`
	}{
		AccessToken: token,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)
}
