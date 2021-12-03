package main;
import ("net"
	"fmt"
	"time"
	"io/ioutil"
        "strconv"
)
func main() {

        tcpServer, _ := net.ResolveTCPAddr("tcp4", ":8888")
        listener, _ := net.ListenTCP("tcp", tcpServer)

        for {
                //当有新的客户端请求来的时候，拿到与客户端的连接
                conn, err := listener.Accept()
                if err != nil {
                        fmt.Println(err)
                        continue
                }   
 
                //处理逻辑
                go handle(conn)
        }   
}

func handle(conn net.Conn) {
        defer conn.Close()                                                                                                                                                                                                                                                    

        //读取客户端传送的消息
        go func() {
                response, _ := ioutil.ReadAll(conn)
                fmt.Println(string(response))
        }() 

        //向客户端发送消息
        time.Sleep(1 * time.Second)
        // now := time.Now().String()
        b:=[]byte("")
        res := appendResp(b, "200 OK", "", "<div style='color:red;font-size:50px;'>fuck</div>")
        conn.Write([]byte(res))
}

func appendResp(b []byte, status, head, body string) []byte {
        b = append(b, "HTTP/1.1"...)
        b = append(b, ' ')
        b = append(b, status...)
        b = append(b, '\r', '\n')
        b = append(b, "Server: gnet\r\n"...)
        b = append(b, "Date: "...)
        b = time.Now().AppendFormat(b, "Mon, 02 Jan 2006 15:04:05 GMT")
        b = append(b, '\r', '\n')
        if len(body) > 0 {
                b = append(b, "Content-Length: "...)
                b = strconv.AppendInt(b, int64(len(body)), 10)
                b = append(b, '\r', '\n')
        }
        b = append(b, head...)
        b = append(b, '\r', '\n')
        if len(body) > 0 {
                b = append(b, body...)
        }
        return b
}
