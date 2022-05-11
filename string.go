package main
import "fmt"

func main(){
    var code  int = 123
    var enddate string = "2022"
    var url = "code=%d&enddate=%s"
    var format_str = fmt.Sprintf(url,code,enddate)
    fmt.Println(format_str)
}
