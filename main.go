/*
21.09.04 시작
https://nomadcoders.co/nomadcoin/lectures/2939

 GOPATH=D:\go_project
 GOPROXY=https://proxy.golang.org,direct
 GOROOT=C:\Program Files\Go
 GOBIN=

노마드 코인 go 언어로 제작해보기
- 단, 실제 코인이 아닌, 어떤 식으로 "동작"하는지 해당 개념을 구현하는 프로젝트임
- making "Useless Coin "


1일차
- go mod 는 package.json 과 같음

*/

package main

import "fmt"

const stephano string = "stephano"

//function can return multiple value
// func plus(a int, b int, name string) (int, string) {
// 	return a + b, name
// }

func plus(a ...int) int {
	var total int            // total :=0
	for _, item := range a { //index, item
		total += item
	}

	return total
}

func main() {
	//x := 405949484
	//return string format
	//xAsBinary := fmt.Sprintf("%b\n", x)

	// fmt.Printf("%o\n", x)
	// fmt.Printf("%x\n", x)
	// fmt.Printf("%U\n", x)
	//fmt.Printf("\n%v,%T\n", xAsBinary, xAsBinary)

	//go에서는 array 사이즈가 고정되어있음 (fixed size). 무한 아님
	//need slice (infinite size, = vector)

	//array
	foods := [3]string{"potato", "pizza", "pasta"}
	// for _, value := range foods {
	// 	fmt.Println(value)
	// }
	for i := 0; i < len(foods); i++ {
		fmt.Println(foods[i])
	}
	//slice
	foods_slice := []string{"potato", "pizza", "pasta"}
	fmt.Printf("%v\n", foods_slice)

	foods_slice_2 := append(foods_slice, "tomato")
	fmt.Printf("%v\n", foods_slice_2)
	fmt.Println(len(foods_slice_2))
}
