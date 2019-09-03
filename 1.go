package main
import (
	"fmt"
	"time"
)
func sum(nums []int, ch chan int, sleep time.Duration){
	sum := 0
	for _, elem := range nums{
		sum += elem
	}
	time.Sleep(sleep * time.Millisecond)
	ch <- sum
}
func main(){
	nums := []int{1,2,3,4,5,6}
	ch := make(chan int)
	go sum(nums[:3], ch,0)
	go sum(nums[3:], ch,1000)

	//time.Sleep(1000 * time.Millisecond)
	front, rear := <- ch, <-ch	
	fmt.Println(front, rear)
}