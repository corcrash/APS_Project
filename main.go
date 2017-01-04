package main

import (
	"APS_Project/database"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const STATIC_URL string = "/static/"
const STATIC_ROOT string = "static/"
const TEMPLATES_ROOT string = "templates/"

type Context struct {
	Title  string
	Static string
}

func Home(w http.ResponseWriter, req *http.Request) {
	context := Context{Title: "Index"}
	render(w, "index.html, partials-navbar.html, partials-verticalmenu.html", context)
}

func About(w http.ResponseWriter, req *http.Request) {
	context := Context{Title: "Login"}
	render(w, "login.html", context)
}

func render(w http.ResponseWriter, tmpl string, context Context) {
	context.Static = STATIC_URL
	tmpl_list := strings.Split(tmpl, ", ")
	tmpl_list = append([]string{"base.html"}, tmpl_list...)

	var t_list []string
	for _, tpl := range tmpl_list {
		t_list = append(t_list, TEMPLATES_ROOT+tpl)
	}

	fmt.Println(t_list)

	t, err := template.ParseFiles(t_list...)
	if err != nil {
		log.Print("template parsing error: ", err)
	}
	err = t.Execute(w, context)
	if err != nil {
		log.Print("template executing error: ", err)
	}
}

func StaticHandler(w http.ResponseWriter, req *http.Request) {
	static_file := req.URL.Path[len(STATIC_URL):]
	if len(static_file) != 0 {
		f, err := http.Dir(STATIC_ROOT).Open(static_file)
		if err == nil {
			content := io.ReadSeeker(f)
			http.ServeContent(w, req, static_file, time.Now(), content)
		}
	}
	http.NotFound(w, req)
}

func main() {

	// Open database connection
	db, err := gorm.Open("mysql", "aps:password@/apsdb?charset=utf8&parseTime=True")
	if err != nil {
		log.Fatal("Could not connect to DB: ", err)
	} else {
		defer db.Close()
	}

	database.GlobalDB = db

	log.Print("Connected to DB!")

	db.LogMode(true)

	http.HandleFunc("/", Home)
	http.HandleFunc("/login/", About)
	http.HandleFunc(STATIC_URL, StaticHandler)
	err = http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("Error listening on port 8000: ", err)
	}

	log.Print("Listening on port 8000")
}
