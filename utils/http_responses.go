package utils

import (
	"encoding/json"
	"net/http"
)

func Write(w http.ResponseWriter, responseStatus int, T interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(responseStatus)
	jsonObj, err := json.Marshal(T)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write(jsonObj)
}

func Success(w http.ResponseWriter, T interface{}) {
	Write(w, http.StatusOK, T)
}

func Created(w http.ResponseWriter, T interface{}) {
	Write(w, http.StatusCreated, T)
}

func BadRequest(w http.ResponseWriter, T interface{}) {
	Write(w, http.StatusBadRequest, T)
}

func Unauthorized(w http.ResponseWriter, T interface{}) {
	Write(w, http.StatusUnauthorized, T)
}

func Forbidden(w http.ResponseWriter, T interface{}) {
	Write(w, http.StatusForbidden, T)
}

func NotFound(w http.ResponseWriter, T interface{}) {
	Write(w, http.StatusNotFound, T)
}
