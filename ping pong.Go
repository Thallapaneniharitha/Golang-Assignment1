1.//Ping Pong
package main
import (
	"fmt"
	"time"
)
func main() {
	pingChan := make(chan int, 1)
	pongChan := make(chan int, 1)
 
	go ping(pingChan, pongChan)
	go pong(pongChan, pingChan)
 
	pingChan <- 1
 
	select {}
}
func ping(pingChan <-chan int, pongChan chan<- int) {
	for {
		ball := <-pingChan
 
		fmt.Println("Ping", ball)
		time.Sleep(1 * time.Second)
 
		pongChan <- ball+1
	}
}
func pong(pongChan <-chan int, pingChan chan<- int) {
	for {
		ball := <-pongChan
 
		fmt.Println("Pong", ball)
		time.Sleep(1 * time.Second)
 
		pingChan <- ball+1
	}
}
$go run main.go
Ping 1
Pong 2
Ping 3
Pong 4
Ping 5
Pong 6
Ping 7
Pong 8
Ping 9
Pong 10
Ping 11
Pong 12
...
...
...