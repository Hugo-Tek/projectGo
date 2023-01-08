package errors

import "fmt"

func Check(err error) int {
	if err != nil {
		fmt.Println(err)
		return 1
	}
	return 0
}
