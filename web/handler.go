package web

import (
	"bank-app/database/bank"
	"bank-app/database/schema"
	"bytes"
	"database/sql"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"time"
)

func Handler(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Web-UI - Request", r.Method, r.URL.Path)
		next(w, r)
	}
}

func get_data(url string, data []byte) (bank.Responce, error) {
	client := http.Client{
		Timeout: time.Second * 30,
	}

	req, err := http.NewRequest(http.MethodPost, baseUrl+url, bytes.NewBuffer(data))
	if err != nil {
		return bank.Responce{}, err
	}

	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return bank.Responce{}, err
	}

	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return bank.Responce{}, err
	}

	var result bank.Responce
	err = json.Unmarshal(b, &result)
	if err != nil {
		return bank.Responce{}, err
	}

	return result, nil
}

func postLogin(w http.ResponseWriter, r *http.Request) {
	var data schema.GetAccountByUsernameParams
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if data.Username == "" || data.Password == "" {
		http.Error(w, "Username and password are required", http.StatusBadRequest)
		return
	}

	detail, err := json.Marshal(data)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := get_data("/getId", detail)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func postSignup(w http.ResponseWriter, r *http.Request) {
	var data struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Username  string `json:"username"`
		Password  string `json:"password"`
		Email     string `json:"email"`
	}

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if data.FirstName == "" || data.LastName == "" {
		http.Error(w, "First name and last name are required", http.StatusBadRequest)
		return
	}

	if data.Username == "" || data.Password == "" {
		http.Error(w, "Username and password are required", http.StatusBadRequest)
		return
	}

  detail, err := json.Marshal(schema.CreateAccountParams{
    ID: "",
    FirstName: data.FirstName,
    LastName: data.LastName,
    Username: data.Username,
    Password: data.Password,
    Email: sql.NullString{
      String: data.Email,
      Valid: data.Email != "",
    },
    Balance: 0,
  })
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := get_data("/create", detail)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

  slog.Info("Web-UI - Response","%v", res)

	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
