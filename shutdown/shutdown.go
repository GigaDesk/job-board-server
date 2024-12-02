package shutdown

import (
	"fmt"
	"os"
	"time"
)

var IsShutdown *bool

func InitiateShutdown(err error){
    ticker := time.NewTicker(time.Second)
	go func(){
	now := time.Now()
	futureTime := now.Add(10 * time.Minute) // Add 10 minutes
	fmt.Println("critical error: ", err)
	*IsShutdown = true
	fmt.Println("system shutdown initiated at: ", now)
       for{
		select{
		case t:= <-ticker.C:
			fmt.Printf("\rsystem shutdown in: %v", futureTime.Sub(t))
		}
	   }
	}()
	time.Sleep(10*time.Minute)
	fmt.Println("system shutdown completed at : ", time.Now())
	os.Exit(1)
}