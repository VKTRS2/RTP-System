package funcs

import "time"

func Now(format string) string {
	return time.Now().Format(format)
}
