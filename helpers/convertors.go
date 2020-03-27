package helpers

import (
	"fmt"
	"os"
	"strconv"
)

//StringToInt : converts string to int.
func StringToInt(toConv string) int {
	result, err := strconv.Atoi(toConv)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	return result
}

//StringToBool : converts string to bool.
func StringToBool(toConv string) bool {
	result, err := strconv.ParseBool(toConv)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	return result
}
