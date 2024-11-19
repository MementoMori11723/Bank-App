package bank

func (acc *Account) get(query string, args ...interface{}) error {
	db, err := connect()
	defer db.Close()
	if err != nil {
		return err
	}

	err = db.
		QueryRow(
			query, args...,
		).Scan()
	return err
}
