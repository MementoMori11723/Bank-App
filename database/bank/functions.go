package bank

import (
	"bank-app/database/schema"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type Responce struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Create(r *http.Request) (Responce, error) {
	var data schema.CreateAccountParams

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		return Responce{}, err
	}

	db, err := connect()
	if err != nil {
		return Responce{}, err
	}

	data.ID = uuid.New().String()

  h := sha256.New()
  h.Write([]byte(data.Password))
  res := h.Sum(nil)

  data.Password = hex.EncodeToString(res)

	user := schema.New(db)
	err = user.CreateAccount(context.Background(), data)
	if err != nil {
		return Responce{}, err
	}

	return Responce{
		Message: "Account Created!",
		Data: schema.Account{
			ID: data.ID,
		},
	}, nil
}

func Deposit(r *http.Request) (Responce, error) {
	var data schema.DepositParams

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		return Responce{}, err
	}

	db, err := connect()
	if err != nil {
		return Responce{}, err
	}

	user := schema.New(db)
	err = user.Deposit(context.Background(), data)
	if err != nil {
		return Responce{}, err
	}

	return Responce{
		Message: "Added Money!",
		Data:    nil,
	}, nil
}

func Balance(r *http.Request) (Responce, error) {
	var data struct {
		ID string `json:"id"`
	}

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		return Responce{}, err
	}

	db, err := connect()
	if err != nil {
		return Responce{}, err
	}

	user := schema.New(db)
	balance, err := user.GetBalance(context.Background(), data.ID)
	if err != nil {
		return Responce{}, err
	}

	return Responce{
		Message: "Here is your balance!",
		Data: schema.Account{
			Balance: balance,
		},
	}, nil
}

func Withdraw(r *http.Request) (Responce, error) {
	var data schema.WithdrawParams

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		return Responce{}, err
	}

	db, err := connect()
	if err != nil {
		return Responce{}, err
	}

	user := schema.New(db)
	err = user.Withdraw(context.Background(), data)
	if err != nil {
		return Responce{}, err
	}

	return Responce{
		Message: "Took Money!",
		Data:    nil,
	}, nil
}

func Delete(r *http.Request) (Responce, error) {
	var data struct {
		ID string `json:"id"`
	}

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		return Responce{}, err
	}

	db, err := connect()
	if err != nil {
		return Responce{}, err
	}

	user := schema.New(db)
	err = user.DeleteAccount(context.Background(), data.ID)
	if err != nil {
		return Responce{}, err
	}

	return Responce{
		Message: "Account Deleted!",
		Data: schema.Account{
			ID: data.ID,
		},
	}, nil
}

func Transactions(r *http.Request) (Responce, error) {
	var data schema.GetTransactionsParams

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		return Responce{}, err
	}

	db, err := connect()
	if err != nil {
		return Responce{}, err
	}

	user := schema.New(db)
	history, err := user.GetTransactions(context.Background(), data)
	if err != nil {
		return Responce{}, err
	}

	return Responce{
		Message: "Transactions History!",
		Data:    history,
	}, nil
}

func Transfer(r *http.Request) (Responce, error) {
  var data schema.InsertTransactionParams

  err := json.NewDecoder(r.Body).Decode(&data)
  if err != nil {
    return Responce{}, err
  }

  db, err := connect()
  if err != nil {
    return Responce{}, err
  }

  user := schema.New(db)
  err = user.InsertTransaction(context.Background(), data)
  if err != nil {
    return Responce{}, err
  }

	return Responce{
		Message: "Successfully transfered!",
		Data:    schema.Account{},
	}, nil
}
