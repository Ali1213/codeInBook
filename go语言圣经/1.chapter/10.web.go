package main
import (
	"fmt"
	"log"
	"net/http"
)
func main(){
	http.HandleFunc("/", hander)
	log.Fatal(http.ListenAndServe(":9090",nil))
}

func hander(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"URL.Path = %q\n", r.URL.Path)
}