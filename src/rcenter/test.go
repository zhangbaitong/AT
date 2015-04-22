package main

import (
	"common"
	"fmt"
	"rcenter/resource"
)

func testFunc() {

	fmt.Println(common.GetUID())
	resource.Delete("2")
	resource.Insert("zhang", 9527, 9528)
	resource.Update("f84c0e10-e811-11e4-8e75-3c075419d855", "baitong", 8888, 9999, 120)
	ress, _ := resource.Query(resource.SELECT_ALL)
	fmt.Println(ress)
}

func main() {
	testFunc()
}
