package main

import "github.com/wwwdotru/go-sandbox/app/go-cookbook/chapter01/tempfiles"

func main() {
	if err := tempfiles.WorkWithTemp(); err != nil {
		panic(err)
	}
}
