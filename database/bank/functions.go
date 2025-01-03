package bank

import (
	"net/http"
)

type Responce struct {
  Message string `json:"message"`
  Data interface{} `json:"data"`
}

func Create(r *http.Request) (Responce, error){
  return Responce{
    Message: "Account Created!",
    Data: nil,
  }, nil
}

func Deposit(r *http.Request) (Responce, error){
  return Responce{
    Message: "Added Money!",
    Data: nil,
  }, nil
}


func Balance(r *http.Request) (Responce, error){
  return Responce{
    Message: "Here is your balance!",
    Data: nil,
  }, nil
}

func Withdraw(r *http.Request) (Responce, error){
  return Responce{
    Message: "Took Money!",
    Data: nil,
  }, nil
}

func Delete (r *http.Request) (Responce, error){
  return Responce{
    Message: "Account Deleted!",
    Data: nil,
  }, nil
}

func Transactions(r *http.Request) (Responce, error){
  return Responce{
    Message: "Transactions History!",
    Data: nil,
  }, nil
}

func Transfer(r *http.Request) (Responce, error){
  return Responce{
    Message: "Successfully transfered!",
    Data: nil,
  }, nil
}
