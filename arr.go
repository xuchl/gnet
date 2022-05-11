package main
import ("fmt"
	"reflect")

func main(){
   var arr = [6]int{1,2,3,4,5}
   for _,v := range arr {
	fmt.Println(v)
	if(&v == nil){
            fmt.Println("空指针");
	}
   }
   var s []int
   if(s == nil){
	fmt.Println("未初始化的切片");
   }
   
   //var a [5]int nil 只能表示通道、函数、切片、接口、和映射
   //if(a == nil){
	//fmt.Println("未初始化的数组")
   //}
   var ptr *int; 
   if(ptr == nil){
	fmt.Println("空指针");
   }

   b:=[5]int{}
   fmt.Printf("%T",b)
   fmt.Println();
   varname := fmt.Sprintf("变量b的类型是%T",b)
   fmt.Println(varname);
   fmt.Println(reflect.TypeOf(b))
}
