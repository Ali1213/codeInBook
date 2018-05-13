package main
import (
	"fmt"
	"net/http"
	"log"
	"sync"
)

var mu sync.Mutex
var count int

func main(){
	http.HandleFunc("/",hander)
	http.HandleFunc("/count",counter)
	log.Fatal(http.ListenAndServe(":9090",nil))
}


func hander(w http.ResponseWriter, r *http.Request){
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w,"URL.Path=%q\n", r.URL.Path)
}


func counter(w http.ResponseWriter, r *http.Request)  {
	mu.Lock()
	fmt.Fprintf(w, "count: %d\n", count)
	mu.Unlock()
}