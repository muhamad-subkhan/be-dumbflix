package handlers

import (
	// filmdto "dumbflix/dto/film"

	dto "dumbflix/dto/result"
	"dumbflix/models"
	"dumbflix/repositories"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type handlerFilm struct {
	FilmRepository repositories.FilmRepository
}

func HandlerFilm(FilmRepository repositories.FilmRepository) *handlerFilm {
	return &handlerFilm{FilmRepository}
}

func (h *handlerFilm) CreateFilm(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "aplication/json")
	
	var request models.Film
	if err := json.NewDecoder(r.Body).Decode(&request)
	err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult {Code: http.StatusBadRequest, Message: err.Error()}
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

	film := models.Film{
		Title: request.Title,
		TumbnailFilm: request.TumbnailFilm,
		Year: request.Year,
		Description: request.Description,
	}

	data, err := h.FilmRepository.CreateFilm(film)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: FilmResponse(data)}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerFilm) GetFilm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	
	Film, err := h.FilmRepository.GetFilm(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: Film}
	json.NewEncoder(w).Encode(response)
}



func FilmResponse(f models.Film) models.FilmResponse {
	return models.FilmResponse{
		Title: 			f.Title,
		TumbnailFilm: 	f.TumbnailFilm,
		Year: 			f.Year,
		Description: 	f.Description,
	}
}