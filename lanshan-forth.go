package main
import(
	"fmt"
)
func fact(n int) int {
	var result int = 1
	for i := 1;i <= n;i++{
		result *= i
	}
	return result
}
func main(){
    var n int = 10
    var ans int 
	ans = fact(n)
	fmt.Println("n的阶乘为:", ans)
}