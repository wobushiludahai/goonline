package main
// 导入系统包
import "fmt"

func main() {
	str := new(string)
	*str = "通过new生成指针测试"
	fmt.Println(*str)
}