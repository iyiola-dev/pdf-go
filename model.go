package main

type Data struct {
	Name          string
	Address       string
	Currency      string
	AccountNumber string
	SwiftCode     string
	AccountType   string
	SortCode      string
	Transactions  []Transaction
}

type Transaction struct {
	Date        string
	Amount      float64
	Type        string
	Description string
	Reference   string
}

func ReturnData() Data {
	return Data{
		Name:          "John Doe",
		Address:       "1234 Main Street",
		Currency:      "USD",
		AccountNumber: "1234567890",
		SwiftCode:     "1234567890",
		AccountType:   "Checking",
		SortCode:      "1234567890",
		Transactions: []Transaction{
			{
				Date:        "2020-01-01",
				Amount:      100.00,
				Type:        "credit",
				Description: "Deposit",
				Reference:   "1234567890",
			},
			{
				Date:        "2020-01-02",
				Amount:      -50.00,
				Type:        "debit",
				Description: "Withdrawal",
				Reference:   "1234567890",
			},
		},
	}
}
