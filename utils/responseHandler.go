package utils

import (
	"encoding/json"
	"net/http"
)

type response struct {
	Success bool
	Error   string
	Data    interface{}
}

func Success(w http.ResponseWriter, data interface{}) {
	var res response

	res.Success = true
	res.Data = data

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(res)
}

func Error(w http.ResponseWriter, code int, err string) {
	var res response

	res.Success = false
	res.Error = err
	res.Data = nil

	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(res)
}

func SysError(w http.ResponseWriter) {
	var res response

	res.Success = false
	res.Error = "Unexpected Database Error."
	res.Data = nil

	w.WriteHeader(http.StatusInternalServerError)
	_ = json.NewEncoder(w).Encode(res)
}
