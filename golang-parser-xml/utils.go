package main

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func fromFloatToCents(a float64) int {
	return int(a * 100)
}
