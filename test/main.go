package test

import (
	"fmt"
	"github.com/rojack96/mcc"
)

func main() {
	testone := mcc.NewMccReader()
	res, _ := testone.FindByCode("204")
	fmt.Println(res)
}
