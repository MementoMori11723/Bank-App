package bank

import (
	"bank-app/database/schema"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type Responce struct {
	Message string           `json:"message"`
	Data    []schema.History `json:"data,omitempty"`
	UserId  string           `json:"user_id,omitempty"`
}

func Create(r *http.Request) (Responce, error) {
	var data schema.CreateAccountParams

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		return Responce{}, err
	}

	db, err := connect()
	defer db.Close()
	if err != nil {
		return Responce{}, err
	}

	data.ID = uuid.New().String()

	data.Password = encryptPassword(data.Password)

	user := schema.New(db)
	err = user.CreateAccount(context.Background(), data)
	if err != nil {
		return Responce{}, err
	}

	return Responce{
		Message: "Account Created!",
	}, nil
}

func Deposit(r *http.Request) (Responce, error) {
	var data schema.DepositParams

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		return Responce{}, err
	}

	db, err := connect()
	defer db.Close()
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
	}, nil
}

func Withdraw(r *http.Request) (Responce, error) {
	var data schema.WithdrawParams

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		return Responce{}, err
	}

	db, err := connect()
	defer db.Close()
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
	}, nil
}

func Delete(r *http.Request) (Responce, error) {
	var data struct {
		Username string `json:"username"`
		ID       string `json:"id"`
	}

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		return Responce{}, err
	}

	db, err := connect()
	defer db.Close()
	if err != nil {
		return Responce{}, err
	}

	user := schema.New(db)
	err = user.DeleteAccount(context.Background(), data.ID)
	if err != nil {
		return Responce{}, err
	}

	err = user.DeleteHistory(context.Background(), schema.DeleteHistoryParams{
		Sender:   data.Username,
		Receiver: data.Username,
	})
	if err != nil {
		return Responce{}, err
	}

	return Responce{
		Message: "Account Deleted!",
	}, nil
}

func Transactions(r *http.Request) (Responce, error) {
	var data schema.GetTransactionsParams

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		return Responce{}, err
	}

	db, err := connect()
	defer db.Close()
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
	defer db.Close()
	if err != nil {
		return Responce{}, err
	}

	data.ID = uuid.New().String()
	data.Timestamp = time.Now().UTC().Format("2006-01-02 15:04:05")

	user := schema.New(db)
	err = user.InsertTransaction(context.Background(), data)
	if err != nil {
		return Responce{}, err
	}

	return Responce{
		Message: "Successfully transfered!",
	}, nil
}

func GetIdByUserName(r *http.Request) (Responce, error) {
	var data schema.GetAccountByUsernameParams
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		return Responce{}, err
	}

	db, err := connect()
	defer db.Close()
	if err != nil {
		return Responce{}, err
	}

	data.Password = encryptPassword(data.Password)

	user := schema.New(db)
	res, err := user.GetAccountByUsername(context.Background(), data)
	if err != nil {
		return Responce{}, err
	}

	return Responce{
		Message: "Found User!",
		UserId:  res.ID,
	}, nil
}

func Details(r *http.Request) (Responce, error) {
	var data schema.GetAccountByUsernameParams
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		return Responce{}, err
	}

	db, err := connect()
	defer db.Close()
	if err != nil {
		return Responce{}, err
	}

	data.Password = encryptPassword(data.Password)

	user := schema.New(db)
	res, err := user.GetAccountByUsername(context.Background(), data)
	if err != nil {
		return Responce{}, err
	}

	return Responce{
		Message: fmt.Sprintf(
			"User Details:\n\tID: %s\n\tFirstName: %s\n\tLastName: %s\n\tUsername: %s\n\tEmail: %v\n\tBalance: %.2f",
			res.ID,
			res.FirstName,
			res.LastName,
			res.Username,
			res.Email.String,
			res.Balance,
		),
	}, nil
}

func CheckUser(r *http.Request) (Responce, error) {
  data := r.PathValue("username")
  if data == "" {
    return Responce{}, fmt.Errorf("Username is not set!")
  }

	db, err := connect()
	defer db.Close()
	if err != nil {
		return Responce{}, err
	}

	user := schema.New(db)
	res, err := user.CheckUserExists(context.Background(), data)
	if err != nil {
		return Responce{}, err
	}

  return Responce{
    Message: "User Found!",
    UserId: res,
  }, nil
}

func encryptPassword(password string) string {
  h := sha256.New()
  h.Write([]byte(password))
  return hex.EncodeToString(h.Sum(nil))
}
