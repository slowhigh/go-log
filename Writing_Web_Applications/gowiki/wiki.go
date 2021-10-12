package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

/*
func loadPage(title string) *Page {
	filename := title + ".txt"
	body, _ := ioutil.ReadFile(filename) // In loadPage, error isn't being handled yet; the "blank identifier" represented by the underscore (_) symbol is used to throw away the error return value (in essence, assigning the value to nothing).
	return &Page{Title: title, Body: body}
}
*/

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):] // len("/view/") => 6, [A:B] => index의 A부터 B까지 부분 문자열, ex) r.URL.Path = /view/TestPage, r.URL.Path[len("/view/"):] => TestPage
	log.Println("title : " + title)
	p, _ := loadPage(title) // Again, note the use of _ to ignore the error return value from loadPage. This is done here for simplicity and generally considered bad practice.
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	fmt.Fprintf(w, "<h1>Editing %s</h1>"+
	"<form action=\"/save/%s\" method=\"POST\">"+
	"<textarea name=\"body\">%s</textarea><br>"+
	"<input type=\"submit\" value=\"Save\">"+
	"</form>",
	p.Title, p.Title, p.Body)
}

func main() {
	// p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	// p1.save()
	// p2, _ := loadPage("TestPage")
	// fmt.Println(string(p2.Body))

	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	//http.HandleFunc("/save/", saveHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// http://localhost:8080/view/test => show a page titled "test" containing the words "Hello world".