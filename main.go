package main
import (
	"fmt"
)
func main() {
	ch := make(chan string)
	go sendData(ch)
	getData(ch)
}
func sendData(ch chan string) {
	for i := 1; i <= 10; i++{
		ch <- "\n张三"
		ch <- "\n李四"
		ch <- "\n王五"
		ch <- "\n赵六"
	}
	close(ch)
}
func getData(ch chan string) {
	for{
		input, open := <-ch
		if !open {
			break
		}
		fmt.Printf("%s ", input)
	}
}