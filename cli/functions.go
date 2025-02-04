package cli

import (
	"bank-app/database/bank"
	"bank-app/database/middleware"
	"bank-app/database/schema"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

var baseURL string

func get_response(url string, reqBody []byte) (bank.Responce, error) {
	client := http.Client{
		Timeout: time.Second * 30,
	}

	if baseURL == "" {
		return bank.Responce{}, fmt.Errorf("Base URL not set")
	}

	req, err := http.NewRequest(
		http.MethodPost, baseURL+url,
		bytes.NewBuffer(reqBody),
	)
	if err != nil {
		return bank.Responce{}, err
	}

  middleware.BaseURL(baseURL)
  token := middleware.GenerateToken()

  req.Header.Set("Content-Type", "application/json")
  req.Header.Set("Accept", "application/json")
  req.Header.Set("Authorization", "Bearer " + token)

	res, err := client.Do(req)
	if err != nil {
		return bank.Responce{}, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return bank.Responce{}, err
	}

	var bodyRes bank.Responce

	err = json.Unmarshal(body, &bodyRes)
	if err != nil {
		return bank.Responce{}, err
	}

	return bodyRes, nil
}

func fetch_responce(url string) {
	data := sub_menu(url)
	res, err := get_response(url, data)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println(res.Message)
	if res.Data.Transactions != nil {
		for i, history := range res.Data.Transactions {
			fmt.Printf(
				"Entry %d: ID: %s Sender: %s Receiver: %s Amount: %.2f Timestamp: %s\n",
				i+1, history.ID,
				history.Sender,
				history.Receiver,
				history.Amount,
				history.Timestamp,
			)
		}
	}
}

func get_id(username, password string) (string, error) {
	data, err := json.Marshal(schema.GetAccountByUsernameParams{
		Username: username,
		Password: password,
	})
	if err != nil {
		return "", err
	}

	res, err := get_response("getId", data)
	if err != nil {
		return "", err
	}

	return res.UserId, nil
}

func check_user(username string) (string, error) {
	check, err := get_response("checkUser/"+username, nil)
	if err != nil {
		return "", err
	}

	return check.UserId, nil
}

func create() []byte {
	var firstName, lastName, userName, password, confirmPassword, email string
	var balance float64

	keys := []string{"First Name", "Last Name", "Email", "Username", "Password", "Confirm Password"}

	inputFunc(keys, map[string]*string{
		"First Name":       &firstName,
		"Last Name":        &lastName,
		"Username":         &userName,
		"Password":         &password,
		"Confirm Password": &confirmPassword,
		"Email":            &email,
	})

	if password != confirmPassword {
		errorFunc(fmt.Errorf("Passwords do not match!"))
	}

	inputFunc([]string{"Amount"}, map[string]*float64{
		"Amount": &balance,
	})

	data, err := json.Marshal(schema.CreateAccountParams{
		FirstName: firstName,
		LastName:  lastName,
		Username:  userName,
		Password:  password,
		Email:     email,
		Balance:   balance,
	})
	if err != nil {
		errorFunc(err)
	}

	return data
}

func deposit() []byte {
	var userName, password string
	var amount float64

	keys := []string{"Username", "Password"}

	inputFunc(keys, map[string]*string{
		"Username": &userName,
		"Password": &password,
	})

	inputFunc([]string{"Amount"}, map[string]*float64{
		"Amount": &amount,
	})

	res, err := get_id(userName, password)
	if err != nil {
		errorFunc(err)
	}

	data, err := json.Marshal(schema.DepositParams{
		Balance: amount,
		ID:      res,
	})
	if err != nil {
		errorFunc(err)
	}

	insert, err := json.Marshal(schema.InsertTransactionParams{
		Sender:   userName,
		Receiver: "credit",
		Amount:   amount,
	})
	if err != nil {
		errorFunc(err)
	}

	_, err = get_response("transfer", insert)
	if err != nil {
		errorFunc(err)
	}

	return data
}

func withdraw() []byte {
	var userName, password string
	var amount float64

	keys := []string{"Username", "Password"}

	inputFunc(keys, map[string]*string{
		"Username": &userName,
		"Password": &password,
	})

	inputFunc([]string{"Amount"}, map[string]*float64{
		"Amount": &amount,
	})

	res, err := get_id(userName, password)
	if err != nil {
		errorFunc(err)
	}

	data, err := json.Marshal(schema.WithdrawParams{
		Balance: amount,
		ID:      res,
	})
	if err != nil {
		errorFunc(err)
	}

	insert, err := json.Marshal(schema.InsertTransactionParams{
		Sender:   userName,
		Receiver: "debit",
		Amount:   amount,
	})
	if err != nil {
		errorFunc(err)
	}

	_, err = get_response("transfer", insert)
	if err != nil {
		errorFunc(err)
	}

	return data
}

func balance() []byte {
	var userName, password string
	type dataType schema.GetAccountByUsernameParams

	keys := []string{"Username", "Password"}

	inputFunc(keys, map[string]*string{
		"Username": &userName,
		"Password": &password,
	})

	data, err := json.Marshal(schema.GetAccountByUsernameParams{
		Username: userName,
		Password: password,
	})
	if err != nil {
		errorFunc(err)
	}

	return data
}

func history() []byte {
	var userName, password string
	type dataType struct {
		ID string
	}

	keys := []string{"Username", "Password"}

	inputFunc(keys, map[string]*string{
		"Username": &userName,
		"Password": &password,
	})
	_, err := get_id(userName, password)
	if err != nil {
		errorFunc(err)
	}

	data, err := json.Marshal(schema.GetTransactionsParams{
		Sender:   userName,
		Receiver: userName,
	})

	return data
}

func transfer() []byte {
	var userName, password, reciverUserName string
	var amount float64

	keys := []string{"Username", "Password", "Receiver Username"}

	inputFunc(keys, map[string]*string{
		"Username":          &userName,
		"Password":          &password,
		"Receiver Username": &reciverUserName,
	})

	inputFunc([]string{"Amount"}, map[string]*float64{
		"Amount": &amount,
	})

	res, err := get_id(userName, password)
	if err != nil {
		errorFunc(err)
	}

	check, err := check_user(reciverUserName)
	if err != nil {
		errorFunc(err)
	}

	if check == "" {
		errorFunc(fmt.Errorf("Receiver does not exist!"))
	}

	withdrawData, err := json.Marshal(schema.WithdrawParams{
		Balance: amount,
		ID:      res,
	})
	if err != nil {
		errorFunc(err)
	}
	_, err = get_response("withdraw", withdrawData)
	if err != nil {
		errorFunc(err)
	}

	depositData, err := json.Marshal(schema.DepositParams{
		Balance:  amount,
		Username: reciverUserName,
	})
	if err != nil {
		errorFunc(err)
	}
	_, err = get_response("deposit", depositData)
	if err != nil {
		errorFunc(err)
	}

	data, err := json.Marshal(schema.InsertTransactionParams{
		Sender:   userName,
		Receiver: reciverUserName,
		Amount:   amount,
	})
	if err != nil {
		errorFunc(err)
	}

	return data
}

func deleteFunc() []byte {
	var userName, password string

	keys := []string{"Username", "Password"}

	inputFunc(keys, map[string]*string{
		"Username": &userName,
		"Password": &password,
	})

	res, err := get_id(userName, password)
	if err != nil {
		errorFunc(err)
	}

	data, err := json.Marshal(struct {
		Username string `json:"username"`
		ID       string `json:"id"`
	}{
		Username: userName,
		ID:       res,
	})
	if err != nil {
		errorFunc(err)
	}

	return data
}
