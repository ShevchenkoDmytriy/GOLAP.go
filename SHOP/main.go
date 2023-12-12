package main

import (
	"fmt"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func MainPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("View/Mainpage.html", "View/header.html", "View/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "Mainpage", nil)
}
func HandlePage() {
	http.Handle("/CSS/", http.StripPrefix("/CSS/", http.FileServer(http.Dir("./CSS/"))))
	http.Handle("/Img/", http.StripPrefix("/Img/", http.FileServer(http.Dir("./Img/"))))
	rtr := mux.NewRouter()
	rtr.HandleFunc("/", MainPage)
	// rtr.HandleFunc("/create", create).Methods("GET")
	// rtr.HandleFunc("/SaveArticle", SaveArticle).Methods("POST")
	// rtr.HandleFunc("/post/{Id:[0-9]+}", AboutPost).Methods("GET")
	http.Handle("/", rtr)
	http.ListenAndServe(":8080", nil)
}
func main() {
	HandlePage()
}
