package cli

import (
	"bank-app/database"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func get_response(url string) (string, error) {
	res, err := http.Get("http://localhost:11000/" + url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	var bodyRes database.Response
	err = json.Unmarshal(body, &bodyRes)
	if err != nil {
		return "", err
	}
	return bodyRes.Message, nil
}

func fetch_responce(url string) {
	res, err := get_response(url)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println(res)
	}
}
