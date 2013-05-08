// socket.go
package main  
  
import (  
  "fmt"  
  "net"  
  "io"  
)  
  
func main() {  
  var (  
    host = "www.apache.org"  
    port = "80"  
    remote = host + ":" + port  
    msg string = "GET / \n"  
    data = make([]uint8, 4096)  
    read = true  
    count = 0  
  )  
  // create the socket  
  con, err := net.Dial("tcp", remote)   
  // send our message.  an HTTP GET request in this case   
  io.WriteString(con, msg)  
  // read the response from the webserver   
  for read {  
    count, err = con.Read(data)  
    read = (err == nil)  
    fmt.Printf(string(data[0:count]))  
  }  
  
  con.Close()  
}  
