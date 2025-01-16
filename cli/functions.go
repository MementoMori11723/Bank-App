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
	} else {
		fmt.Println(res.Message)
	}
}

func get_id(username, password string) (schema.Account, error) {
	data, err := json.Marshal(schema.GetAccountByUsernameParams{
		Username: username,
		Password: password,
	})
	if err != nil {
		return schema.Account{}, err
	}

	res, err := get_response("getId", data)
	if err != nil {
		return schema.Account{}, err
	}

	var account schema.Account
	accountData, err := json.Marshal(res.Data)
	if err != nil {
		return schema.Account{}, err
	}
	err = json.Unmarshal(accountData, &account)
	if err != nil {
		return schema.Account{}, err
	}

	return account, nil
}

var subMenu = map[string]func() []byte{
	"create": func() []byte {
		var firstName, lastName, userName, password, email string
		var balance float64

		fmt.Print("First Name : ")
		fmt.Scanln(&firstName)
		fmt.Print("Last Name : ")
		fmt.Scanln(&lastName)
		fmt.Print("Username : ")
		fmt.Scanln(&userName)
		fmt.Print("Password : ")
		fmt.Scanln(&password)
		fmt.Print("Email : ")
		fmt.Scanln(&email)
		fmt.Print("Amount : ")
		fmt.Scanln(&balance)

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
			fmt.Println("Error", err)
		}

		return data
	},

	"deposit": func() []byte {
		var userName, password string
		var amount float64

		fmt.Print("Username : ")
		fmt.Scanln(&userName)
		fmt.Print("Password : ")
		fmt.Scanln(&password)
		fmt.Print("Amount : ")
		fmt.Scanln(&amount)

		res, err := get_id(userName, password)
		if err != nil {
			fmt.Println("Error: ", err)
		}

		data, err := json.Marshal(schema.DepositParams{
			Balance: amount,
			ID:      res.ID,
		})
		if err != nil {
			fmt.Println("Error", err)
		}

		return data
	},

	"withdraw": func() []byte {
		var userName, password string
		var amount float64

		fmt.Print("Username : ")
		fmt.Scanln(&userName)
		fmt.Print("Password : ")
		fmt.Scanln(&password)
		fmt.Print("Amount : ")
		fmt.Scanln(&amount)

		res, err := get_id(userName, password)
		if err != nil {
			fmt.Println("Error: ", err)
		}

		data, err := json.Marshal(schema.WithdrawParams{
			Balance: amount,
			ID:      res.ID,
		})
		if err != nil {
			fmt.Println("Error", err)
		}

		return data
	},

	"balance": func() []byte {
		var userName, password string
		type dataType struct {
			ID string
		}

		fmt.Print("Username : ")
		fmt.Scanln(&userName)
		fmt.Print("Password : ")
		fmt.Scanln(&password)
		res, err := get_id(userName, password)
		if err != nil {
			fmt.Println("Error: ", err)
		}

		data, err := json.Marshal(dataType{
			ID: res.ID,
		})
		if err != nil {
			fmt.Println("Error", err)
		}

		return data
	},

	"transactions": func() []byte { return []byte{} },
	"transfer":     func() []byte { return []byte{} },
}

func sub_menu(menu string) []byte {
	run, ok := subMenu[menu]
	if !ok {
		panic("Error Occured at sub_menu!")
	}
	data := run()
	return data
}
