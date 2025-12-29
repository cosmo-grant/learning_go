package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var ErrInvalidID = errors.New("invalid ID")

type EmptyFieldError struct {
	fieldName string
}

func (e EmptyFieldError) Error() string {
	return e.fieldName
}

func main() {
	d := json.NewDecoder(strings.NewReader(data))
	count := 0
	for d.More() {
		count++
		var emp Employee
		err := d.Decode(&emp)
		if err != nil {
			fmt.Printf("record %d: %v\n", count, err)
			continue
		}
		err = ValidateEmployee(emp)
		message := fmt.Sprintf("record %d: %+v", count, emp)
		/*
		* I tried to convince the compiler that err implements Unwrap().
		* E.g. `e, ok := error.(interface{ Unwrap() []error })`, then `if ok`.
		* But no dice.
		* So resorting to a type switch.
		* Not happy with this approach but going with it for now.
		 */

		if err != nil {
			switch err := err.(type) {
			case interface{ Unwrap() []error }:
				// this case will always obtain
				allErrors := err.Unwrap()
				var messages []string
				for _, e := range allErrors {
					messages = append(messages, processError(e, emp))
				}
				message = message + " allErrors: " + strings.Join(messages, ", ")
			}
		}
		fmt.Println(message)
	}
}

func processError(err error, emp Employee) string {
	var emptyFieldErr EmptyFieldError
	if errors.Is(err, ErrInvalidID) {
		return fmt.Sprintf("invalid ID: %s", emp.ID)
	} else if errors.As(err, &emptyFieldErr) {
		return fmt.Sprintf("empty field %s", emptyFieldErr.fieldName)
	} else {
		return fmt.Sprintf("%v", err)
	}
}

const data = `
{
	"id": "ABCD-123",
	"first_name": "Bob",
	"last_name": "Bobson",
	"title": "Senior Manager"
}
{
	"id": "XYZ-123",
	"first_name": "Mary",
	"last_name": "Maryson",
	"title": "Vice President"
}
{
	"id": "BOTX-263",
	"first_name": "",
	"last_name": "Garciason",
	"title": "Manager"
}
{
	"id": "HLXO-829",
	"first_name": "Pierre",
	"last_name": "",
	"title": "Intern"
}
{
	"id": "MOXW-821",
	"first_name": "Franklin",
	"last_name": "Watanabe",
	"title": ""
}
{
	"id": "",
	"first_name": "Shelly",
	"last_name": "Shellson",
	"title": "CEO"
}
{
	"id": "YDOD-324",
	"first_name": "",
	"last_name": "",
	"title": ""
}
`

type Employee struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Title     string `json:"title"`
}

var validID = regexp.MustCompile(`\w{4}-\d{3}`)

func ValidateEmployee(e Employee) error {
	validationErrors := []error{}
	if len(e.ID) == 0 {
		validationErrors = append(validationErrors, errors.New("missing ID"))
	} else if !validID.MatchString(e.ID) {
		validationErrors = append(validationErrors, ErrInvalidID)
	}
	if len(e.FirstName) == 0 {
		validationErrors = append(validationErrors, EmptyFieldError{"FirstName"})
	}
	if len(e.LastName) == 0 {
		validationErrors = append(validationErrors, EmptyFieldError{"LastName"})
	}
	if len(e.Title) == 0 {
		validationErrors = append(validationErrors, EmptyFieldError{"Title"})
	}

	/*
			When len(s) > 0, errors.Join(s...) implements Unwrap().
		  When len(s) == 0 (or contains only nils), it returns nil.
		  But even behind a len(s) > 0 guard, the compiler complains if you give the return value as implementing Unwrap().
		  So we type switch in the caller instead.
	*/
	return errors.Join(validationErrors...)
}
