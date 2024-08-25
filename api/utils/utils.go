package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Decode[T any](r io.Reader) (T, error) {
	var v T
	if err := json.NewDecoder(r).Decode(&v); err != nil {
		return v, fmt.Errorf("error decoding the json: %v", err.Error())
	}
	return v, nil
}

func Encode[T any](w io.Writer, v T) error {
	if err := json.NewEncoder(w).Encode(v); err != nil {
		return fmt.Errorf("error encoding to json: %v", err.Error())
	}
	return nil
}

func PostJSON[T any](w http.ResponseWriter, v T, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	Encode(w, v)
}
