package process

type helpMenu struct {
	Route       string `json:"route"`
	Description string `json:"description"`
	Method      string `json:"method"`
}

func help() []helpMenu {
	return []helpMenu{
		{Route: "/create-account", Description: "Create an account", Method: "POST"},
		{Route: "/deposit-money", Description: "Deposit money", Method: "POST"},
		{Route: "/withdraw-money", Description: "Withdraw money", Method: "POST"},
		{Route: "/check-balance", Description: "Check balance", Method: "GET"},
		{Route: "/view-transactions-history", Description: "View transactions history", Method: "GET"},
		{Route: "/settings", Description: "Settings", Method: "PUT"},
	}
}
