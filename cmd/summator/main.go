package main

import (
	"fmt"
	summator "template/internal/services"
)

func main() {
	res := summator.DigitSummator(10, 10)
	fmt.Println(res)
}
