package main

import (
	"fmt"
	"github.com/bluele/gen-validator/examples/models"
)

func main() {
	user := models.User{
		Name: "b",
		Age:  20,
	}
	errs := user.Validate()
	if len(errs) != 0 {
		fmt.Println(errs[0])
	}
}
