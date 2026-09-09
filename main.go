package main 
import(
  "fmt"
  "net/http"
)
func main(){
  http.HandleFunc("/",func(w http.ReponseWriter,r*http.Request){
    fmt.Fprint(w,"Hello from Google App Engine -Go!")
})
 http.ListenAndServer(":8080",nil)
}
