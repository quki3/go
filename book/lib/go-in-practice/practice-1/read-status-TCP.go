package main 
import (
	"bufio"
	"fmt"
	"net"
)

func main () {
	conn, _ := net.Dial("tcp", "goland.org:80") // return connections
	/* sends string over the connection */
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	status, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Println(status)
}
