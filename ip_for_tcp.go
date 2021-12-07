package main

import (
    //. "fmt"
      "net"
	"flag"
	"log"
	"time"
)

func main() {
    a := flag.String("p", "ip4:tcp", "protocol")
    b := flag.String("l", "127.0.0.1", "listen ip")
    flag.Parse()
    laddr, err := net.ResolveIPAddr(*a, *b)
    // tcpServer, err := net.ResolveTCPAddr("tcp4", ":8888")
    if err != nil {
        log.Fatal(err)

    }
   c, err := net.ListenIP(*a, laddr)
        if err != nil {
            log.Fatal(err)

        }
    for {

     

        buf := make([]byte, 2048)
        n, addr, err := c.ReadFromIP(buf)
        if err != nil {
            log.Println(err)
            time.Sleep(time.Second)
            continue
        }
        ID := int(buf[5]) + int(buf[4])*256
        log.Println(addr, " ", ID, buf[0:n])
    }
}
