package main
import(
	"fmt"
)
func main(){
	var n int = 1000
	var sum int
	for i := 1;i <= n;i++{
        sum += i
	}
	fmt.Println(sum)
}