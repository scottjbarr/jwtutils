package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
)

func main() {
	// the JWT string
	input := os.Args[1]

	// create a parser
	p := new(jwt.Parser)

	// Claims to write in to.
	claims := jwt.MapClaims{}

	// Parse the token, without verifying authenticity.
	//
	// This is only intended for reading tokens.
	token, _, err := p.ParseUnverified(input, &claims)
	if err != nil {
		panic(err)
	}

	// JSON output
	b, err := json.Marshal(token.Claims)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", b)
}
