package main

import "github.com/nikitsingh/forky/backend/internal"

func main() {
	if err := internal.RunSetup(); err != nil {
		panic(err)
	}
}
