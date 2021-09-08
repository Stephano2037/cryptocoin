/*
21.09.04 시작
https://nomadcoders.co/nomadcoin/lectures/2939



노마드 코인 go 언어로 제작해보기
- 단, 실제 코인이 아닌, 어떤 식으로 "동작"하는지 해당 개념을 구현하는 프로젝트임
- making "Useless Coin "


*/

package main

import (
	"fmt"
	"net/http"
)

const port string = ":4000"

//home should be store 2 arguments
// 1. writer (Responsewriter , 유저에 보내고 싶은 데이터)
//2. 포인터 (request는 복사 용도가 아님)
func home(rw http.ResponseWriter, r *http.Request) {
	//not print in console, but writer
	fmt.Fprint(rw, "hello from home!")
}

func main() {
	//1. home page 2. add page
	http.HandleFunc("/", home)
	fmt.Printf("Listening on http://localhost%s\n", port)

	//log.Fatal(http.ListenAndServe(port, nil))
}
