package main

import (
	"flag"
	"fmt"
	"github.com/pquerna/otp/totp"
	"time"
)

func main() {
	var secret string
	flag.StringVar(&secret, "s", "", "MFA")
	flag.Parse()
	if len(secret) == 0 {
		fmt.Println("secret empty")
		return
	}

	k, err := totp.GenerateCode(secret, time.Now())
	if err != nil {
		fmt.Println("fail imput")
		return
	}

	fmt.Println(k)

}
