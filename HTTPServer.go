package main

import (
	// "fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

type HTTPServer struct {
	channel chan string
	port    string
}

func NewHTTPServer() *HTTPServer {
	ch := make(chan string)

	return &HTTPServer{
		channel: ch,
		port:    "8080",
	}
}

func (httpServer *HTTPServer) start() {
	// start http server
	address := "localhost:" + httpServer.port

	http.HandleFunc("/", viewHandler)

	log.Fatal(http.ListenAndServe(address, nil))
}

// http Stuff
func handler(w http.ResponseWriter, r *http.Request) {
	// html := "hallo"
	//fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	// fmt.Fprintf(w, html)
	path := r.URL.Path[1:]
	data, err := ioutil.ReadFile(string(path))

	if err == nil {
		w.Write(data)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("404 Pate not found"))
	}
}

// func handler2(w http.ResponseWriter, r *http.Request) {
// 	var err error
//
// 	tpl := template.New("tpl.gohtml")
// 	tpl = tpl.Funcs(template.FuncMap{
// 		"uppercase": func(str string) string {
// 			return strings.ToUpper(str)
// 		},
// 	})
// 	tpl, err = tpl.ParseFiles("tpl.gohtml")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	err = tpl.Execute(os.Stdout, Page{
// 		Title: "My Title 2",
// 		Body:  `hello world <script>alert("Yow!");</script>`,
// 	})
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// }

// ##############################

// Page struct, which will contain template
type Page struct {
	Title   string
	Body    template.HTML
	Number  int
	Players []string
}

// Loads a page for use
func loadPage(title string, r *http.Request) (*Page, error) {
	body, err := ioutil.ReadFile(title)
	if err != nil {
		return nil, err
	}

	p := []string{"player 1", "player 2", "player 3"}
	return &Page{Title: title, Body: template.HTML(body), Number: 23, Players: p}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	httpServer.channel <- r.URL.Path
	// Parses URL to obtain title of file to add to .body
	title := r.URL.Path[len("/"):]

	// Load templatized page, given title
	page, _ := loadPage(title, r)

	// Generate template t
	t, _ := template.ParseFiles("index.html")

	// Write the template attributes of page (from load page) to t
	t.Execute(w, page)
}