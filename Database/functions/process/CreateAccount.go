package process 

import "fmt"

type Account struct {}

func CreateAccount() string {
	fmt.Println("Create an account")
	return "Account created" //just for testing purposes.
}
