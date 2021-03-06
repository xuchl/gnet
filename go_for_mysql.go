package main
import (
    // "github.com/go-sql-driver/mysql" 
    "database/sql"
    "strings"
    "fmt")
const (
    userName = "root"
    password = "123456"
    ip = "127.0.0.1"
    port = "3307"
    dbName = "valuation"
)
//Db数据库连接池
var DB *sql.DB

//注意方法名大写，就是public
func InitDB()  {
    //构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
    path := strings.Join([]string{userName, ":", password, "@tcp(",ip, ":", port, ")/", dbName, "?charset=utf8"}, "")

    //打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
    DB, _ = sql.Open("mysql", path)
    //设置数据库最大连接数
    DB.SetConnMaxLifetime(100)
    //设置上数据库最大闲置连接数
    DB.SetMaxIdleConns(10)
    //验证连接
    if err := DB.Ping(); err != nil{
        fmt.Println("opon database fail")
        return
    }
    fmt.Println("connnect success")
}


func main() {
    // type Books struct {
    //    name string
    //    id int
    // }
    // u := Books{}
    var userName string
    userName=""
    id:=1
    err := DB.QueryRow("SELECT * FROM nk_user WHERE id = ?", id).Scan(&userName)
    if err != nil{
        fmt.Println("查询出错了")
    }
    fmt.Print(userName)
}