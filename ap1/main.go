package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world! This is Ap1 container.")
}

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	// r.ParseForm()

	fmt.Println(r.Form)
	fmt.Println("Path", r.URL.Path)
	fmt.Println("Scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello, astaxie!")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method: ", r.Method)
	// r.Method で クライアントからのメソッドが取れる
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	} else {
		// ログインデータがリクエストされ、ログインのロジック判定が実行される。
		// Formからの入力をParseするには、r.ParseForm() を明示的に呼び出す必要がある。
		r.ParseForm()
		fmt.Println("Username: ", r.Form["username"])
		fmt.Println("Password: ", r.Form["password"])
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/hello", sayHelloName)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":3000", nil)
	log.Println("ListenAndServe on port :3000")
	if err != nil {
		log.Fatal("http.ListenAndServe: ", err)
	}
}
