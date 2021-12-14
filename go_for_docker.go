package main

import "fmt"
import "time"

func main() {
    for true  {
        fmt.Printf("这是无限循环。\n");
        time.Sleep(time.Second)
    }
}
