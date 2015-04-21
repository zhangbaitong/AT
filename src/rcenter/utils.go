package main

import (
	"fmt"
	"github.com/zhangbaitong/go-uuid/uuid"
)

func main() {
	uuid1 := uuid.NewUUID()
	fmt.Print(uuid1)
}
