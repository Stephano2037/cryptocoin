/*
21.09.04 시작
https://nomadcoders.co/nomadcoin/lectures/2939



노마드 코인 go 언어로 제작해보기
- 단, 실제 코인이 아닌, 어떤 식으로 "동작"하는지 해당 개념을 구현하는 프로젝트임
- making "Useless Coin "


*/

package main

import (
	explorer "cryptocoin/explorer/templates"
	"cryptocoin/rest"
)

func main() {

	go explorer.Start(3000)
	//단순히 두개 동시에 실행하면 아래 rest.Start 는 블로킹 상태임
	// 위에 explorer 를 go 로 돌리면? -> panic
	// 문제는 url을 다루는 무언가가 rest 와 explore 에 있는 url 둘다 다루는 것 (/) <- 이게문제
	// -- panic: http: multiple registrations for /
	// url 과 url 함수를 매핑해주는 건 같음

	//handler 는 보통 nil 임 -> nil 이면 DefaultServeMux 가 사용된다는 의미
	//Mux -> Multiplexer url 로의 request 를 다루는것
	//-- url 을 보고있다가 handler를 호출하는 결정을 하는 것
	// 단순히 포트 중복되지않는다고하여 에러가 안나는게 아님
	// 커스텀 multiplexer 제작 필요
	rest.Start(5000)

}
