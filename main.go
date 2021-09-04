/*
21.09.04 시작
https://nomadcoders.co/nomadcoin/lectures/2939



노마드 코인 go 언어로 제작해보기
- 단, 실제 코인이 아닌, 어떤 식으로 "동작"하는지 해당 개념을 구현하는 프로젝트임
- making "Useless Coin "


*/

package main

import (
	"crypto/sha256"
	"fmt"
)

/*
block chain structure


hash -> one way function (not duplix)
"hello" -> h_fn(x) -> "wjiofjewoijfowejof"
"hello1" -> h_fn(x) -> "wszcvzxcvbbnnnnnn!!"
can't get original value by reverse
*/

//block 1번의 hash 는 block 1번의 data + 이전 block의 hash 로 계산
//block 1번이 변경되면 block 3번도 변경된다 그래야 의미있다
//common block chain's way of hash is using [SHA256] Algorithm
//Go language has already have this algorithm
//hash 가 원하는 것은 불변하는 값
/*
	b1
		b1hash = (data + "")
			or
		b1hash. = (data + "x")
	b2
		b2hash.. = (data + b1hash)
	b3
		b3hash... = (data + b2hash)
*/
type block struct {
	data     string
	hash     string
	prevHash string
}

func main() {
	//블록체인의 시작을 보통 genesisBlock 이라 명명하는듯
	genesisBlock := block{"Genesis Block", "", ""}
	//get slice array byte
	hash := sha256.Sum256([]byte(genesisBlock.data + genesisBlock.prevHash))
	//fmt.Printf("%x,%T\n", hash, hash)
	//fmt.Println(len(hash))

	hexHash := fmt.Sprintf("%x", hash)
	genesisBlock.hash = hexHash
	fmt.Printf("%v,%T\n", genesisBlock.hash, genesisBlock.hash)

}
