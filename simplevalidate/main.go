package main

import (
	"fmt"
	"reflect"
	"strings"
)

type Post struct {
	Title string `json:"title" validate:"required,max=100"`
}

func main() {
	// Create instance of struct
	post := Post{}

	// use reflect to access the struct tag
	t := reflect.TypeOf(post)
	field, _ := t.FieldByName("Title")

	// Extract the tag content for validate
	validateTag := field.Tag.Get("validate")

	// Optionally, parse the validate tag to get individual rules
	validateRules := strings.Split(validateTag, ",")

	fmt.Println("Validate tag:", validateTag)
	fmt.Println("Validate rules:", validateRules)

}
