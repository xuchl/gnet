package main
import "fmt"
import "github.com/go-sql-driver/mysql"

type persion struct{
	name string
	toal int
	color string
}

func (t persion) setName(name string) persion{
	t.name = name
	return t
}

func (t persion) setHeight(toal int) persion{
	t.toal = toal
	return t
}

func (t persion) setColor(color string) persion{
	t.color = color
	return t
}
func main(){
	var t = new(persion).setName("fuck").setHeight(170).setColor("red")
	fmt.Printf("\npersions color is%s\n",t.color)
	fmt.Printf("persions toal is%d\n",t.toal)
}