package main
import(
	"fmt"
)
func main(){
	var(
		n  int
		m  int = 0
		sum  int
	)
	for {
		fmt.Scanln(&n)
		if n == 0 {
			break
		}
		m++
		sum += n
	}

	var ans float64
	ans = aver(float64(sum),float64(m))

	if ans >= 60 {
		fmt.Printf("平均成绩为 %f,成绩合格", ans)
	} else {
		fmt.Printf("平均成绩为 %f,成绩不合格", ans)
	}
	
}

func aver (sum float64 , count float64) float64{
	return sum/count
}