/*
21.09.04 시작
https://nomadcoders.co/nomadcoin/lectures/2939



노마드 코인 go 언어로 제작해보기
- 단, 실제 코인이 아닌, 어떤 식으로 "동작"하는지 해당 개념을 구현하는 프로젝트임
- making "Useless Coin "


*/

package main

import (
	"cryptocoin/blockchain"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type homeData struct {
	//lower case , is "private" , template 에도 영향을 줌 다른곳에서 사용할땐 대문자로!
	PageTitle string
	Blocks    []*blockchain.Block
}

const (
	port        string = ":4000"
	templateDir string = "templates/"
)

var templates *template.Template

//home should be store 2 arguments
// 1. writer (Responsewriter , 유저에 보내고 싶은 데이터)
//2. 포인터 (request는 복사 용도가 아님)
func home(rw http.ResponseWriter, r *http.Request) {
	//not print in console, but writer
	//templ, err := template.ParseFiles("templates/home.html")
	/*
		if err != nil {
			log.Fatal(err)
		}
	*/
	data := homeData{"Home", blockchain.GetBlockchain().AllBlocks()}
	templates.ExecuteTemplate(rw, "home", data)

}

func add(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		templates.ExecuteTemplate(rw, "add", nil)
	case "POST":
		r.ParseForm()
		data := r.Form.Get("blockData")
		blockchain.GetBlockchain().AddBlock(data)
		http.Redirect(rw, r, "/", http.StatusPermanentRedirect)
	}
	templates.ExecuteTemplate(rw, "add", nil)
}

func main() {
	//1. home page 2. add page
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml"))
	http.HandleFunc("/", home)
	http.HandleFunc("/add", add)
	fmt.Printf("Listening on http://localhost%s\n", port)

	log.Fatal(http.ListenAndServe(port, nil))
}
