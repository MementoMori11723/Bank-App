package web

import (
	"bank-app/database/bank"
	"bank-app/database/middleware"
	"bank-app/database/schema"
	"bytes"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"strings"
	"time"
)

type Data struct {
	FirstName string  `json:"first_name,omitempty"`
	LastName  string  `json:"last_name,omitempty"`
	Username  string  `json:"username,omitempty"`
	Password  string  `json:"password,omitempty"`
	Email     string  `json:"email,omitempty"`
	Amount    float64 `json:"amount,omitempty"`
}

func (d *Data) CheckString() (string, bool) {
	data := map[string]string{
		"first name": d.FirstName,
		"last name":  d.LastName,
		"username":   d.Username,
		"password":   d.Password,
		"email":      d.Email,
	}

	for k, v := range data {
		if !middleware.CheckString(
			strings.TrimSpace(v),
		) {
			return k + " is not valid\nContains special characters", false
		}
	}

	return "", true
}

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
	var data Data
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

	if str, ok := data.CheckString(); !ok {
		http.Error(w, str, http.StatusBadRequest)
		return
	}

	detail, err := json.Marshal(schema.GetAccountByUsernameParams{
		Username: data.Username,
		Password: data.Password,
	})
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
	var data Data

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

	if data.Email == "" {
		http.Error(w, "Email is required", http.StatusBadRequest)
		return
	}

	if str, ok := data.CheckString(); !ok {
		http.Error(w, str, http.StatusBadRequest)
		return
	}

	detail, err := json.Marshal(schema.CreateAccountParams{
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Username:  data.Username,
		Password:  data.Password,
		Email:     data.Email,
		Balance:   0,
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

	slog.Info("Web-UI - Response", "%v", res)

	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
