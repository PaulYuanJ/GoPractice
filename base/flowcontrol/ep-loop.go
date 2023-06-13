package flowcontrol

import (
	"fmt"
	"github.com/jasonlvhit/gocron"
	"time"
)

func EP_for() {
	// use break to jump loop
	{
		count := 0
		for {
			count += 1
			fmt.Println("FirstFor: ", count)
			time.Sleep(1 * time.Second)
			if count >= 5 {
				break
			}
		}
	}

	// use condition to do loop
	{
		for count1:=0; count1 <= 5; count1++ {
			fmt.Println("SecondFor: ", count1)
		}
	}

	// use range to loop string array
	{
		var myArray = []string{"Third: 1", "Third: 2", "Third: 3"}
		for ix, str := range myArray {
			fmt.Println("Index: ", ix, "content: ", str)
		}
	}

}

func demoFunc2(myarg string) {
	fmt.Println("In demoFunc2", myarg)
}

func EP_cron() {

	// use framework to do cron loop
	s := gocron.NewScheduler()

	_ = s.Every(1).Second().Do(func() {fmt.Println("In demoFunc")})
	_ = s.Every(2).Seconds().Do(demoFunc2, "Hello World!")

	// will be blocking
	<- s.Start()
}