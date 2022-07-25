package main

import (
	_ "embed"
	"fmt"
	"github.com/rs/zerolog/log"
)

//go:embed sha.txt
var sha string

//go:embed VERSION
var version string

// appLogo contains the ASCII splash screen
//go:embed app-logo.txt
var appLogo []byte

func main() {

	fmt.Println(string(appLogo))
	fmt.Printf("Version: %s\n", version)
	fmt.Printf("Sha: %s\n", sha)
	_, err := ReadConfig()
	if nil != err {
		log.Fatal().Err(err).Send()
	}

}
