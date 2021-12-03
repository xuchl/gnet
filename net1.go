package main;
import (
	"fmt"
	"net"
	"bufio"
)
func main() {
	conn, err := net.Dial("tcp", "baidu.com:80")
	if err != nil {
		// handle error
	}
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	status, err := bufio.NewReader(conn).ReadString('\n')
	fmt.Println(status)
}