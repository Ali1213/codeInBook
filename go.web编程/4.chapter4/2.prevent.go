package main

import (
	"log"
	"fmt"
	"strconv"
	"io"
	"crypto/md5"
	"time"
	"net/http"
	"html/template"
)

func login(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h,strconv.FormatInt(crutime,10))
		token := fmt.Sprintf("%x",h.Sum(nil))
		t,_ := template.ParseFiles("login.2.html")
		t.Execute(w,token)
	}else{
		r.ParseForm()
		token := r.Form.Get("token")
		if token != "" {

		}else{

		}

		fmt.Println("username length",len(r.Form["username"][0]))
		fmt.Println("username:",template.HTMLEscapeString(r.Form.Get("username")))
		fmt.Println("password:",template.HTMLEscapeString(r.Form.Get("password")))

		template.HTMLEscape(w, []byte(r.Form.Get("username")))
	}

}


func main(){

	http.HandleFunc("/login",login)
	err := http.ListenAndServe(":9090",nil)
	if err != nil{
		log.Fatal("ListenAndServe:",err)
	}
}