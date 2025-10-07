package handler

import (
    "encoding/json"
    "net/http"
)

func GetUser(env *Env, w http.ResponseWriter, r *http.Request) error {
	email := r.PathValue("email")

	if email == "" {
		return StatusError{400, nil}
	}

	user, err := env.DB.GetUserByEmail(email)

	if err != nil {
		return StatusError{500, err}
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(user); err != nil {return StatusError{500, err}}

	return nil
}