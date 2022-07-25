package funcs

import "time"

func RtpId() string {
	return time.Now().Format("20060102150304")
}
