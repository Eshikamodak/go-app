package main 
import(
  "fmt"
  "net/http"
)
func main(){
  http.HandleFunc("/",func(w http.ResponseWriter,r*http.Request){
    fmt.Fprint(w,"Hello from Google App Engine -Go!")
})
 http.ListenAndServe(":8080",nil)
}
