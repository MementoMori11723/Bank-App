package cli

import (
	"bank-app/database/bank"
	"bank-app/database/schema"
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func get_response(url string, reqBody []byte) (bank.Responce, error) {
	client := http.Client{
		Timeout: time.Second * 30,
	}

	req, err := http.NewRequest(
		"GET", "http://localhost:11000/"+url,
		bytes.NewBuffer(reqBody),
	)
	if err != nil {
		return bank.Responce{}, err
	}

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
	if res.Data != nil {
		for i, history := range res.Data {
			fmt.Printf(
				"Entry %d: ID: %s Sender: %s Receiver: %s Amount: %.2f Timestamp: %s\n",
				i+1, history.ID,
				history.Sender,
				history.Receiver,
				history.Amount,
				map[bool]string{
					true:  history.Timestamp.String,
					false: "NULL",
				}[history.Timestamp.Valid],
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

var subMenu = map[string]func() []byte{
	"create":       create,
	"deposit":      deposit,
	"withdraw":     withdraw,
	"balance":      balance,
	"transactions": history,
	"transfer":     transfer,
}

func inputFunc[T any](keys []string, m map[string]*T) {
	for _, key := range keys {
		fmt.Print(key, " : ")
		fmt.Scanln(m[key])
	}
}

func errorFunc(err error) {
	fmt.Println("Error: ", err)
	os.Exit(1)
}

func create() []byte {
	var firstName, lastName, userName, password, email string
	var balance float64

	keys := []string{"First Name", "Last Name", "Username", "Password", "Email"}

	inputFunc(keys, map[string]*string{
		"First Name": &firstName,
		"Last Name":  &lastName,
		"Username":   &userName,
		"Password":   &password,
		"Email":      &email,
	})

	inputFunc([]string{"Amount"}, map[string]*float64{
		"Amount": &balance,
	})

	data, err := json.Marshal(schema.CreateAccountParams{
		FirstName: firstName,
		LastName:  lastName,
		Username:  userName,
		Password:  password,
		Email: sql.NullString{
			String: email,
			Valid:  true,
		},
		Balance: balance,
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

	fmt.Println(res)

	data, err := json.Marshal(schema.WithdrawParams{
		Balance: amount,
		ID:      res,
	})
	if err != nil {
		errorFunc(err)
	}

	return data
}

func balance() []byte {
	var userName, password string
	type dataType struct {
		ID string
	}

	keys := []string{"Username", "Password"}

	inputFunc(keys, map[string]*string{
		"Username": &userName,
		"Password": &password,
	})

	res, err := get_id(userName, password)
	if err != nil {
		errorFunc(err)
	}

	data, err := json.Marshal(dataType{
		ID: res,
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

func sub_menu(menu string) []byte {
	run, ok := subMenu[menu]
	if !ok {
		panic("Error Occured at sub_menu!")
	}
	data := run()
	return data
}
