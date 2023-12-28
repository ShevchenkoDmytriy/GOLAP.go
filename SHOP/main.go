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
type PageVariables struct {
	Check bool
}

var posts = []Users{}
var redirectURL string
var products = []Products{}
var about = Products{}

func MainPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("View/Mainpage.html")
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
	t, err := template.ParseFiles("View/Registrpage.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "Registrpage", nil)
}
func Loginpage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("View/Loginpage.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "Loginpage", nil)
}
func SaveUser(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")
	userType := r.FormValue("userType")
	check := false
	t, _ := template.ParseFiles("View/Registrpage.html")
	if email == "" || password == "" || (userType != "user" && userType != "seller") {
		fmt.Fprintf(w, "Try again")
	} else {
		db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/Shop.go")
		if err != nil {
			panic(err.Error)
		}
		var exists bool
		err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM Users WHERE email = ?)", email).Scan(&exists)
		if err != nil {
			panic(err.Error)
		}
		if exists {
			check = true
			t.ExecuteTemplate(w, "Registrpage", PageVariables{Check: check})
			return
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	send, err := db.Query("SELECT * FROM `Users`")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	posts = []Users{}

	for send.Next() {
		var post Users
		err = send.Scan(&post.id_user, &post.email, &post.password, &post.typeUser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		posts = append(posts, post)
		if email == post.email && password == post.password {
			redirectURL = fmt.Sprintf("/User/%d", post.id_user)
			http.Redirect(w, r, redirectURL, http.StatusSeeOther)
			return
		}
	}
	http.Redirect(w, r, "/Loginpage", http.StatusSeeOther)
}

func MainpageWithRegi(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("View/MainpageWithRegi.html")
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
	t.ExecuteTemplate(w, "MainpageWithRegi", products)
}
func About(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("View/About.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(r)
	productName := vars["product_name"]

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/Shop.go")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	row := db.QueryRow("SELECT * FROM `Products` WHERE `product_name`=?", productName)

	var about Products
	err = row.Scan(&about.ProductId, &about.ProductName, &about.Price, &about.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			http.NotFound(w, r)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t.ExecuteTemplate(w, "About", about)
}
func SearchProducts(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("productsName")
	t, err := template.ParseFiles("View/Search.html")
	if name == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/Shop.go")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM Products WHERE product_name LIKE ?", "%"+name+"%")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var products []Products

	for rows.Next() {
		var product Products
		err = rows.Scan(&product.ProductId, &product.ProductName, &product.Price, &product.Description)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		products = append(products, product)

	}
	t.ExecuteTemplate(w, "Search", products)
}
func SearchPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("View/Search.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "Search", nil)
}
func Basketpage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("View/Basket.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "Basket", nil)
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

	rtr.HandleFunc("/Registration", Registpage).Methods("GET")
	rtr.HandleFunc("/Loginpage", Loginpage)
	rtr.HandleFunc("/", MainPage)
	rtr.HandleFunc("/SaveUser", SaveUser).Methods("POST", "GET")
	rtr.HandleFunc("/CheckUser", CheckUser).Methods("POST")
	rtr.HandleFunc("/User/{Id:[0-9]+}", MainpageWithRegi).Methods("GET")
	rtr.HandleFunc("/Products/{product_name}", About).Methods("GET")
	rtr.HandleFunc("/SearchPage", SearchProducts).Methods("POST", "GET")
	rtr.HandleFunc("/Search", SearchPage).Methods("GET")
	// rtr.HandleFunc("/Products/{{.Id}}", About).Methods("GET")
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
