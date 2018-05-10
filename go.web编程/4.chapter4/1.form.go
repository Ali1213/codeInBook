package main

import (
	"fmt"
	"net/http"
	"strings"
	"log"
	"html/template"
)

func sayhelloName(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("Path", r.URL.Path)
	fmt.Println("Scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k,v := range(r.Form) {
		fmt.Println("key:",k)
		fmt.Println("Value:",strings.Join(v,""))
	}
	fmt.Fprint(w, "hello astaxie!")
}

func login(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t,_ := template.ParseFiles("login.html")
		t.Execute(w, nil)
	} else {	
		for k,v := range(r.Form) {
			fmt.Println("key:",k)
			fmt.Println("Value:",strings.Join(v,""))
		}
		fmt.Println("username", r.Form["username"])
		fmt.Println("password", r.Form["password"])
	}
}


func main(){
	http.HandleFunc("/",sayhelloName)
	http.HandleFunc("/login",login)
	err := http.ListenAndServe(":9090",nil)
	if err != nil{
		log.Fatal("ListenAndServe:",err)
	}
}