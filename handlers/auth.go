package handlers

import (
	authdto "dumbflix/dto/auth"
	dto "dumbflix/dto/result"
	"dumbflix/models"
	"dumbflix/pkg/bcrypt"
	jwtToken "dumbflix/pkg/jwt"
	"dumbflix/repositories"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	// "os/user"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
)

type handlerAuth struct {
	AuthRepository repositories.AuthRepository
}

func HandlerAuth(AuthRepository repositories.AuthRepository) *handlerAuth {
	return &handlerAuth{AuthRepository}
}

func (h *handlerAuth) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(authdto.CreateUserRegister)
	if err := json.NewDecoder(r.Body).Decode(&request);
	 err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {

		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	password, err := bcrypt.HashingPassword(request.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	user := models.User{
		Fullname: request.Fullname,
		Email:    request.Email,
		Password: password,
		Gender:   request.Gender,
		Phone:    request.Phone,
		Address:  request.Address,
		Subscribe: "true",
	}

	data, err := h.AuthRepository.Register(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerAuth) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	request := new(authdto.LoginRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	userData := models.User{
		Email:    request.Email,
		Password: request.Password,
	}

	// Cek Email
	user, err := h.AuthRepository.Login(userData.Email)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "Email yang anda masukan salah"}
		json.NewEncoder(w).Encode(response)
		return
	}

	//cek password
	// makePassword, _ := bcrypt.HashingPassword(request.Password)
	// fmt.Print(makePassword)
	// fmt.Print(user.Password)
	// fmt.Print("test", makePassword == user.Password)
	Valid := bcrypt.CheckPasswordHash(request.Password, user.Password)
	// fmt.Print(Valid)
	// fmt.Println(user.Password)
	// fmt.Println(request.Password)
	if !Valid {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "password yang anda masukan salah"}
		json.NewEncoder(w).Encode(response)
		return
	}

	//generate token
	claims := jwt.MapClaims{}
	claims["id"] = user.ID
	claims["expired"] = time.Now().Add(time.Hour * 2).Unix() //lama token 2 jam

	token, errToken := jwtToken.GenerateToken(&claims)
	if errToken != nil {
		log.Println(errToken)
		fmt.Println("unauthorize")
		return
	}

	loginResponse := authdto.LoginResponse{
		Email: user.Email,
		Token: token,
	}

	w.Header().Set("Content-Type", "application/json")
	response := dto.SuccessResult{Code: http.StatusOK, Data: loginResponse}
	json.NewEncoder(w).Encode(response)
	return
}
