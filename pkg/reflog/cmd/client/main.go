package main

import (
	"fmt"

	"reflog.com/reflog/pkg/app"
)

func main() {
	app.RegisterFunc("PrintArgs", PrintArgs)
}

func PrintArgs(args ...string) (string, error) {
	return fmt.Sprint(args), nil
}
