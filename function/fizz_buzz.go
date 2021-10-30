package function

import "fmt"

func FizzBuss(total int)  {
	for i := 1; i <= total; i++ {
		if i % 3 == 0 && i % 5 ==0 {
			fmt.Println("fizz & buzz")
		} else if i % 3 == 0 {
			fmt.Println("FIZZ")
		} else if i % 5 == 0{
			fmt.Println("BUZZ")
		} else {
			fmt.Println(i)
		}
	}
}
