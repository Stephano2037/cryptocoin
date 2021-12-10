package rest

import (
	"cryptocoin/blockchain"
	"cryptocoin/utility"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var port string

//const port string = ":4000"

// /
// GET
// See documentation
// /blocks
// POST
// Create a book

type url string

//마샬 ,언마샬 할 때, Field 의 결과물을 수정할 수 있는 인터페이스가 존재함
//우리가 인터페이스를 josn으로 변환할 때, 중간에 끼어서 Field 가 json 에서 어찌 보일지 원하는대로 바꿀수있음
//MarshalText 함수
//-- json string 으로서 어찌 보일지 결정하는 method
//--struct 를 json 으로 변환할 때 자동으로 호출
func (u url) MarshalText() ([]byte, error) {
	//fmt.Println("we are here!")
	url := fmt.Sprintf("http://localhost%s%s", port, u)
	//return []byte("hello!"), nil
	return []byte(url), nil
}

//go 가 struct 형식을 json 으로 변경

type urlDescription struct {
	URL         url    `json:"url"` //실제 보여지는 json key 항목이 대문자가 아닌 소문자로 보이게됨
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"`
	////json 영역
	// 무시하고 싶으면 `json:"-"`
}

type addBlockBody struct {
	Message string
}

//format package 가 뭔가를 format 해야할때 String method 를 찾는 거고, ethod 가 구현되있으면 호출하는것
//패키지가 인터페이스에 정의된 메서드를 호출함
//fmt package 에서사용하기 위함!
// func (u URLDescription) String() string {
// 	return "Hello I'm the URL DESCRIPTION"
// }

func documentation(rw http.ResponseWriter, r *http.Request) {
	//유저에게 json 을 먼저 보내는것 시작
	//json을 보내는건 GO 에서 뭔가를 받아서, 유효한 json으로 변환하는것

	data := []urlDescription{
		{
			//API 에서 URL에 하이퍼링크 추가하는 법은?
			URL: url("/"),
			//struct field tags -> 기본적으로 struct 가 json 으로 변환하는 방식을 바꾸는 법
			//보통 json 은 소문자 형식으로 작성됨
			// 아래 항목을 소문자로 바꾸면 go의 약속으로 인하여 public 이 아니게되어 filed는 웹에서 안보임
			//omitempty 는 Field 가 비어있으면 숨기는 역할
			Method: "GET",
			//유저에게 보여주고 싶은것
			//근데 이 data는 Go 의 세계에 있는 slice 임, struct의 slice 이므로, json으로 변경 필요
			//Marshal 은 메모리형식으로 저장된 객체를, 저장/송신 할 수 있도록 변환해 주는 것
			//Go 에서 인터페이스 받고 JSON 변환 -> 마샬
			//JSON을 받고 Go 언어로 변경 시키는 것 -> 언마샬
			Description: "See Documnetation",
		},
		{
			URL:         url("/blockes"),
			Method:      "UPDATE",
			Description: "Add a block",
			Payload:     "data:string",
		},
	}

	//fmt.Println(data)
	//브라우저에서 json으로 인식시켜야함 -> 브라우저 + f12 눌러서 network 항목 header 항목 , refresh  후 확인
	rw.Header().Add("Content-Type", "application/json")

	//json.marshal 보다 더 쉬운것 ? -> 이건 writer 를 받음
	json.NewEncoder(rw).Encode(data) //아래 jsonmarshal 부터 fmt.print 까지 동작하는 걸 한줄로 진행함
	// b, err := json.Marshal(data)
	//utility.HandleErr(err)
	// // fmt.Println(b)
	// // fmt.Printf("%s", b)
	// //console 이 아닌 write에 작성 후, string 된 data 를 보낼 예정
	// //깔끔하게 보이는 이유는 jsonview 라는것이 동작하므로
	// fmt.Fprintf(rw, "%s\n", b)
}

func blocks(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		//얘를 추가해야 진짜 json 형식인것
		fmt.Println("we receive GET")
		rw.Header().Add("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(blockchain.GetBlockchain().AllBlocks())
	case "POST":
		fmt.Println("we receive POST")
		//	유저가 아래와같은 데이터를 보낼거임
		//{data: my block data}
		//rest client extension 사용
		var addblockBody addBlockBody
		//decoder는 reader 를 받습니다.
		utility.HandleErr(json.NewDecoder(r.Body).Decode(&addblockBody))
		//fmt.Println(addblockBody)
		blockchain.GetBlockchain().AddBlock(addblockBody.Message)
		rw.WriteHeader(http.StatusCreated)
	}
}

func Start(Port int) {
	// fmt.Println(URLDescription{
	// 	URL:         "/",
	// 	Method:      "GET",
	// 	Description: "See Documnetation",
	// })
	// ServerMux 는 url(/blocks) 과  url 함수(blocks) 를 연결하는 역할함
	handler := http.NewServeMux()
	port = fmt.Sprintf(":%d", Port)
	handler.HandleFunc("/", documentation)
	handler.HandleFunc("/blocks", blocks)
	fmt.Printf("Listening on http://localhost%s\n", port)
	//nil 은 나중에 바뀔예정
	log.Fatal(http.ListenAndServe(port, handler))
}
