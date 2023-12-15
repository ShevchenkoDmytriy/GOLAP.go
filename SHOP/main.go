package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Users struct {
	id_user  uint
	email    string
	password string
	typeUser string
}
type Products struct {
	ProductId   uint
	ProductName string
	Price       float64
	Description string
}

var posts = []Users{}
var redirectURL string
var products = []Products{}

func MainPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("View/Mainpage.html", "View/header.html", "View/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/Shop.go")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	send, err := db.Query("SELECT * FROM `Products`")
	if err != nil {
		panic(err.Error())
	}
	products = []Products{}

	for send.Next() {
		var prod Products
		err = send.Scan(&prod.ProductId, &prod.ProductName, &prod.Price, &prod.Description)
		if err != nil {
			panic(err.Error())
		}

		products = append(products, prod)
	}
	t.ExecuteTemplate(w, "Mainpage", products)
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

		insert, err := db.Query(fmt.Sprintf("INSERT INTO `Users`(`email`,`password`,`type_user`) VALUES('%s','%s','%s')", email, password, userType))
		if err != nil {
			panic(err.Error)
		}

		defer insert.Close()

		http.Redirect(w, r, "/Loginpage", http.StatusSeeOther)
	}
}
func CheckUser(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/Shop.go")
	if err != nil {
		panic(err.Error)
	}

	defer db.Close()

	send, err := db.Query("SELECT *FROM `Users`")
	if err != nil {
		panic(err.Error)
	}
	posts = []Users{}

	for send.Next() {
		var post Users
		err = send.Scan(&post.id_user, &post.email, &post.password, &post.typeUser)
		if err != nil {
			panic(err.Error)
		}

		posts = append(posts, post)
		if email == post.email && password == post.password {
			redirectURL = fmt.Sprintf("/MainpageWithRegi/%d", post.id_user)
			http.Redirect(w, r, redirectURL, http.StatusSeeOther)
		} else {
			http.Redirect(w, r, "/Loginpage", http.StatusSeeOther)
		}
	}
}
func MainpageWithRegi(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("View/MainpageWithRegi.html", "View/header3.html", "View/footer3.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "MainpageWithRegi", nil)
}

// func AboutPost(w http.ResponseWriter, r *http.Request) {
// 	t, err := template.ParseFiles("template/show.html", "template/header.html", "template/footer.html")
// 	if err != nil {
// 		fmt.Fprintf(w, err.Error())
// 	}
// 	vars := mux.Vars(r)
// 	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/golang")
// 	if err != nil {
// 		panic(err.Error)
// 	}

// 	defer db.Close()

// 	send, err := db.Query(fmt.Sprintf("SELECT *FROM `articles` WHERE `Id`='%s'", vars["Id"]))
// 	if err != nil {
// 		panic(err.Error)
// 	}
// 	showPost = Article{}
// 	for send.Next() {
// 		var post Article
// 		err = send.Scan(&post.Id, &post.Title, &post.Anons, &post.Fulltext)
// 		if err != nil {
// 			panic(err.Error)
// 		}

//			showPost = post
//		}
//		t.ExecuteTemplate(w, "show", showPost)
//	}
func HandlePage() {
	http.Handle("/CSS/", http.StripPrefix("/CSS/", http.FileServer(http.Dir("./CSS/"))))
	http.Handle("/Img/", http.StripPrefix("/Img/", http.FileServer(http.Dir("./Img/"))))
	http.Handle("/Func/", http.StripPrefix("/Func/", http.FileServer(http.Dir("./Func/"))))
	rtr := mux.NewRouter()

	rtr.HandleFunc("/Registration", Registpage)
	rtr.HandleFunc("/Loginpage", Loginpage)
	rtr.HandleFunc("/", MainPage).Methods("GET")
	rtr.HandleFunc("/SaveUser", SaveUser).Methods("POST")
	rtr.HandleFunc("/CheckUser", CheckUser).Methods("POST")
	rtr.HandleFunc("/MainpageWithRegi/{Id:[0-9]+}", MainpageWithRegi).Methods("GET")
	// rtr.HandleFunc("/MainpageWithRegi/{Id:[0-9]+}", MainpageWithRegi).Methods("GET")
	// rtr.HandleFunc("/create", create).Methods("GET")
	// rtr.HandleFunc("/SaveArticle", SaveArticle).Methods("POST")
	// rtr.HandleFunc("/post/{Id:[0-9]+}", AboutPost).Methods("GET")
	http.Handle("/", rtr)
	http.ListenAndServe(":8080", nil)
}
func main() {
	HandlePage()
}
