package main

import (
	"database/sql"
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
func Registpage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("View/Registrpage.html", "View/header2.html", "View/footer2.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "Registrpage", nil)
}
func Loginpage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("View/Loginpage.html", "View/header2.html", "View/footer2.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "Loginpage", nil)
}
func SaveUser(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")
	userType := r.FormValue("userType")

	if email == "" || password == "" || (userType != "user" && userType != "solder") {
		fmt.Fprintf(w, "Try again")
	} else {
		db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/Shop.go")
		if err != nil {
			panic(err.Error)
		}

		defer db.Close()

		insert, err := db.Query(fmt.Sprintf("INSERT INTO `Users`(`Email`,`Password`,`TypeUser`) VALUES('%s','%s','%s')", email, password, userType))
		if err != nil {
			panic(err.Error)
		}

		defer insert.Close()

		http.Redirect(w, r, "/Loginpage", http.StatusSeeOther)
	}
}
func HandlePage() {
	http.Handle("/CSS/", http.StripPrefix("/CSS/", http.FileServer(http.Dir("./CSS/"))))
	http.Handle("/Img/", http.StripPrefix("/Img/", http.FileServer(http.Dir("./Img/"))))
	http.Handle("/Func/", http.StripPrefix("/Func/", http.FileServer(http.Dir("./Func/"))))
	rtr := mux.NewRouter()
	rtr.HandleFunc("/", MainPage)
	rtr.HandleFunc("/Registration", Registpage)
	rtr.HandleFunc("/Loginpage", Loginpage)
	rtr.HandleFunc("/SaveUser", SaveUser).Methods("POST")
	// rtr.HandleFunc("/create", create).Methods("GET")
	// rtr.HandleFunc("/SaveArticle", SaveArticle).Methods("POST")
	// rtr.HandleFunc("/post/{Id:[0-9]+}", AboutPost).Methods("GET")
	http.Handle("/", rtr)
	http.ListenAndServe(":8080", nil)
}
func main() {
	HandlePage()
}
