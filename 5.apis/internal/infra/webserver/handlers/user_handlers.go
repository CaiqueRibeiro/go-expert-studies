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

type Error struct {
	Message string `json:"message"`
}

type UserHandler struct {
	UserDB database.UserInterface
}

func NewUserHandler(db database.UserInterface) *UserHandler {
	return &UserHandler{
		UserDB: db,
	}
}

// Create user 	godoc
// @Summary 	Create a new user
// @Description Create a new user
// @Tags 		users
// @Accept  	json
// @Produce  	json
// @Param 		request body dto.CreateUserInput true "user request"
// @Success 	201
// @Failure 	500 {object} Error
// @Router 		/users [post]
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
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	err = h.UserDB.Create(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

// GetJWT       godoc
// @Summary 	Get a user JWT
// @Description Get a user JWT
// @Tags 		users
// @Accept  	json
// @Produce  	json
// @Param 		request body dto.GetJWTInput true "user credentials"
// @Success 	200 {object} dto.GetJWTOutput
// @Failure 	400 {object} Error
// @Failure 	404 {object} Error
// @Failure 	401
// @Router 		/users/generate_token [post]
func (h *UserHandler) GetJWT(w http.ResponseWriter, r *http.Request) {
	jwt := r.Context().Value("jwt").(*jwtauth.JWTAuth) // Pegando do middleware e fazendo cast do formato esperado
	jwtExpiresIn := r.Context().Value("jwtExpiresIn").(int)
	var u dto.GetJWTInput
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	user, err := h.UserDB.FindByEmail(u.Email)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	isValid := user.ValidatePassword(u.Password)
	if !isValid {
		w.WriteHeader(http.StatusUnauthorized)
		err := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(err)
		return
	}
	_, token, _ := jwt.Encode(map[string]interface{}{
		"sub": user.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(jwtExpiresIn)).Unix(),
	})
	accessToken := dto.GetJWTOutput{AccessToken: token}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)
}
