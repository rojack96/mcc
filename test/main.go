package main

import (
	"fmt"
	"github.com/rojack96/mcc"
)

func main() {
	testone := mcc.NewMccReader()
	res, _ := testone.MncMap("222", mcc.MncCode)
	fmt.Println(res)
}
