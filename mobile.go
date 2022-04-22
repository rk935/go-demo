package main

import (
	"fmt"
	"github.com/nyaruka/phonenumbers"
)

func main() {
	num, _ := phonenumbers.Parse("895613913545", "US")
	formattedNum := phonenumbers.Format(num, phonenumbers.NATIONAL)
	fmt.Println(formattedNum)
	fmt.Println(phonenumbers.IsPossibleNumber(num))
	fmt.Println(phonenumbers.IsValidNumber(num))
}
