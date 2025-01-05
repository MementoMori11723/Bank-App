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

func get_response(_ string, reqBody []byte) (string, error) {
	client := http.Client{
		Timeout: time.Second * 30,
	}

	req, err := http.NewRequest(
		"GET", "http://localhost:11000/create",
		bytes.NewBuffer(reqBody),
	)
	if err != nil {
		return "", err
	}

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	var bodyRes bank.Responce

	err = json.Unmarshal(body, &bodyRes)
	if err != nil {
		return "", err
	}

	return bodyRes.Message, nil
}

func fetch_responce(url string) {
	data := sub_menu(url)
	res, err := get_response(url, data)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println(res)
	}
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
		fmt.Print("balance : ")
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

	"deposit":      func() []byte { return []byte{} },
	"withdraw":     func() []byte { return []byte{} },
	"balance":      func() []byte { return []byte{} },
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
