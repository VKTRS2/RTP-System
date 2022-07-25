package funcs

func Decimal(s string) string {
	if len(s) < 3 {
		s = "000" + s
	}
	return s[:len(s)-2] + "." + s[len(s)-2:]
}
