package explorer

import (
	"cryptocoin/blockchain"
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type homeData struct {
	//lower case , is "private" , template 에도 영향을 줌 다른곳에서 사용할땐 대문자로!
	PageTitle string
	Blocks    []*blockchain.Block
}

const (
	//port        string = ":4000"
	templateDir string = "explorer/templates/"
)

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

var templates *template.Template

//StartExplorer 라고 짓지않는 이유: main 변수에서 중복 explorer 표현이 되는것을 피하기위함
func Start(port int) {
	handler := http.NewServeMux()
	//1. home page 2. add page
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml"))
	handler.HandleFunc("/", home)
	handler.HandleFunc("/add", add)
	fmt.Printf("Listening on http://localhost%d\n", port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), handler))
}
