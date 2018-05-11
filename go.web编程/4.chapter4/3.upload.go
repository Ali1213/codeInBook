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
	"os"
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

func upload(w http.ResponseWriter, r *http.Request){
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime,10))
		token := fmt.Sprintf("%x",h.Sum(nil))

		t,_ := template.ParseFiles("upload.html")
		t.Execute(w,token)
	}else{
		r.ParseMultipartForm(30 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return 
		}
		defer file.Close()
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime,10))
		token := fmt.Sprintf("%x",h.Sum(nil))
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./" + handler.Filename + token, os.O_WRONLY|os.O_CREATE,777)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		io.Copy(f,file)
	}
}


func main(){

	http.HandleFunc("/login",login)
	http.HandleFunc("/upload",upload)
	err := http.ListenAndServe(":9090",nil)
	if err != nil{
		log.Fatal("ListenAndServe:",err)
	}
}