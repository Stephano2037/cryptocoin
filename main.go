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

*/

//go 함수는 대문자로 시작한다
package main

import (
	"cryptocoin/person"
	"fmt"
)

func main() {
	stephano := person.Person{}

	//copy
	stephano.SetDetails("stephano", 12)

	//it didn't show the previous assignment data
	//fmt.Println(stephano)

	fmt.Println(stephano.GetName())
}
