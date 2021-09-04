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

import (
	"fmt"
	"time"
)

const stephano string = "stephano"

type person struct {
	name  string
	birth int
}

//특정 구조체를 명시해야함, 함수를 구조체 안에다가 선언하는게 아님

//func (struct) function name (argument) return types  {}
func (p person) sayHello() {
	fmt.Printf("hello! my name is %s and I'm %d\n", p.name, p.myAge())
}

func (p person) myAge() int {

	var age int = int(time.Now().Year()) - p.birth + 1
	return age
}

func main() {
	// a := 2
	// b := &a
	// c := "stephano"

	// d := &c
	// fmt.Println(*b)
	// fmt.Println(&c, *d)

	// fmt.Println(&a, b)
	// fmt.Println(&c, *d)

	// c = strings.Replace(c, "s", "y", 1)
	// fmt.Println(c)

	stephano := person{name: "stephano", birth: 1993}
	fmt.Println(stephano)
	//stephano := person{"stephano", 12}
	stephano.sayHello()

}
