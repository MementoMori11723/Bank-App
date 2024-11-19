package config

import (
	"log"
	"os"
)

var DB_path = "bank.db"

func init() {
	if _, err := os.Stat(DB_path); os.IsNotExist(err) {
		if _, ferr := os.Create(DB_path); ferr != nil {
			log.Fatalf(
        "Error creating %s: %v", 
        DB_path, ferr,
      )
		}
	} else if err != nil {
		log.Fatalf(
      "Error checking %s: %v", 
      DB_path, err,
    )
	}
}
