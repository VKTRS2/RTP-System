package pain001

import "testing"

func TestTransactionSum(t *testing.T) {
	order1 := Order{
		Transactions: []Transaction{
			{Amount: "10.10"},
			{Amount: "20.20"}},
	}
	order2 := Order{
		Transactions: []Transaction{
			{Amount: "10.10"}},
	}
	sum1, err := order1.transactionSum()
	if err != nil {
		t.Error(err)
	}
	sum2, err := order2.transactionSum()
	if err != nil {
		t.Error(err)
	}
	if sum1 != "30.30" {
		t.Errorf("sum of %s and %s was %s not 30.30", order1.Transactions[0].Amount, order1.Transactions[1].Amount, sum1)
	}
	if sum2 != "10.10" {
		t.Errorf("sum of %s (and nothing else) was %s not 10.10", order2.Transactions[0].Amount, sum2)
	}
}

