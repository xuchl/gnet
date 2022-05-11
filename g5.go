package main
import (
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "strings"
    "database/sql"
    "os"
)
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
type G5 struct{
    id int
    status int
    down_url string
}
func main(){
    InitDB();

    g5 := new(G5)
    rows, err := DB.Query("select id,status,down_url from g5 where title like '%%男同%%' or tag like '%%男同%%'")
    defer func() {
        if rows != nil {
            rows.Close() //可以关闭掉未scan连接一直占用
        }
    }()
    if err != nil {
        fmt.Printf("Query failed,err:%v", err)
        return
    }
    for rows.Next() {
        err = rows.Scan(&g5.id, &g5.status, &g5.down_url) //不scan会导致连接不释放
        if err != nil {
            fmt.Printf("Scan failed,err:%v", err)
            return
        }
        fmt.Print(*g5)
        fmt.Println()
        _,err := DB.Exec("UPDATE g5 set status=? where down_url=?","-3",g5.down_url)
        if err != nil{
            fmt.Printf("Insert failed,err:%v",err)
            return
        }else{
            fmt.Printf("%s状态已修改",g5.down_url)
            fmt.Println()
        }
        logFile := fmt.Sprintf("/Volumes/mynas/china/%s",g5.down_url)
        ferr := os.Remove(logFile)
        if ferr != nil{
            // fmt.Printf("%s文件删除失败",g5.down_url)
        }else{
            fmt.Printf(">>>>>%s文件已删除",g5.down_url)
        }
    }
}
