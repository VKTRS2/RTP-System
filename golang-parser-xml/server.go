package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

//BankAccounts exports for sqlite3 driver
type BankAccounts struct {
	ID               int
	OrganizationName string
	BalanceCents     int
	Iban             string
	Bic              string
}

//Transactions exports for sqlite3 driver
type Transactions struct {
	ID               int
	CounterpartyName string
	CounterpartyIban string
	CounterpartyBic  string
	AmountCents      int
	AmountCurrency   string
	BankAccountID    int
	Description      string
}

func server(listen string) {
	http.HandleFunc("/", handleTransfer)

	err := http.ListenAndServe(listen, nil)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func handleTransfer(w http.ResponseWriter, req *http.Request) {

	decoder := json.NewDecoder(req.Body)
	var t Transfers
	err := decoder.Decode(&t)
	checkErr(err)
	defer req.Body.Close()

	db := initDB("qonto_accounts.sqlite")

	bankAcc := getBankAccountByIban(db, t.Iban)
	transfersAmount := fromFloatToCents(t.TotalAmount)

	if transfersAmount < bankAcc.BalanceCents {
		newBalance := bankAcc.BalanceCents - transfersAmount

		updateBalance(db, newBalance, bankAcc.ID)
		saveTransactions(db, t.Transfers, bankAcc.ID)

		getTransactions(db)
		getBankAccountByIban(db, t.Iban)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte("Not Enough Funds"))
	}
}

func initDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)
	checkErr(err)
	if db == nil {
		panic("db nil")
	}
	return db
}

func getBankAccountByIban(db *sql.DB, i string) BankAccounts {
	query := `
		SELECT * FROM bank_accounts
		WHERE Iban=?
		`

	var result BankAccounts

	rows, err := db.Query(query, i)
	checkErr(err)
	defer rows.Close()

	for rows.Next() {
		item := BankAccounts{}
		err2 := rows.Scan(
			&item.ID,
			&item.OrganizationName,
			&item.BalanceCents,
			&item.Iban,
			&item.Bic,
		)
		checkErr(err2)
		result = item
	}

	fmt.Println(result)
	return result
}

func getTransactions(db *sql.DB) []Transactions {
	query := "SELECT * FROM transactions"

	var result []Transactions

	rows, err := db.Query(query)
	checkErr(err)
	defer rows.Close()

	for rows.Next() {
		item := Transactions{}
		err2 := rows.Scan(
			&item.ID,
			&item.CounterpartyName,
			&item.CounterpartyIban,
			&item.CounterpartyBic,
			&item.AmountCents,
			&item.AmountCurrency,
			&item.BankAccountID,
			&item.Description,
		)
		checkErr(err2)
		result = append(result, item)
	}

	fmt.Println(result)

	return result
}

func updateBalance(db *sql.DB, b, i int) {
	query := `
		UPDATE bank_accounts
		SET balance_cents = $1
		WHERE id = $2
	`
	stmt, err := db.Prepare(query)
	checkErr(err)

	res, err := stmt.Exec(b, i)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println("Number of Bank Account affected: ", affect)
}

func saveTransactions(db *sql.DB, tr []Transfer, id int) {
	query := `
		INSERT INTO transactions(
			counterparty_name,
			counterparty_iban,
			counterparty_bic,
			amount_cents,
			amount_currency,
			bank_account_id,
			description
		) values(?, ?, ?, ?, ?, ?, ?)
	`

	stmt, err := db.Prepare(query)
	checkErr(err)

	defer stmt.Close()

	for _, item := range tr {
		_, err2 := stmt.Exec(
			item.BeneficiaryName,
			item.Iban,
			item.Bic,
			fromFloatToCents(item.Amount),
			"EUR",
			id,
			item.Description,
		)
		checkErr(err2)
	}

}
