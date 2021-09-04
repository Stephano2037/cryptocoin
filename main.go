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

	//result, name := plus(2, 2, "stephano")
	result := plus(2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(result)

	name := "Stephano ! !!!!! Is my name"
	for _, letter := range name {
		//print byte, should be convert formatting
		//fmt.Println(index, letter)

		//fmt.Println("%o ",letter) //octor notation
		//fmt.Println("%b ",letter) //binary notation
		//fmt.Println("%x ",letter) //hex notation
		fmt.Println(string(letter))
	}

}
