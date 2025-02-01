package main

import (
	"fmt"
	"github.com/rojack96/mcc"
)

func main() {
	testone := mcc.NewMccReader()
	res, err := testone.MccInfo("aa")
	fmt.Println(res, err)
}
