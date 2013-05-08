// simple_webserver.go
package main  
  
import (  
     "net/http"  
     "io"  
)  

const form = `<html><body><form action="#" method="post" name="bar">
		      <input type="text" name="in"/>
			  <input type="submit" value="Submit"/>
			  </form></html></body>`

/* handle a simple get request */  
func SimpleServer(w http.ResponseWriter, request *http.Request) {  
     io.WriteString(w, "<h1>hello, world</h1>")  
}  
  
/* handle a form, both the GET which displays the form 
   and the POST which processes it.*/  
func FormServer(w http.ResponseWriter, request *http.Request) {
     w.Header().Set("Content-Type", "text/html")  
     switch request.Method {  
            case "GET":  
                 /* display the form to the user */  
                 io.WriteString(w, form );  
            case "POST":  
                 /* handle the form data, note that ParseForm must 
                    be called before we can extract form data*/  
                 //request.ParseForm();  
                 //io.WriteString(w, request.Form["in"][0])
				io.WriteString(w, request.FormValue("in")) 
     }  
}  
  
func main() {  
     http.HandleFunc("/test1", SimpleServer)  
     http.HandleFunc("/test2", FormServer)  
     if err := http.ListenAndServe(":8088", nil); err != nil {
		panic(err)
     }
}  
