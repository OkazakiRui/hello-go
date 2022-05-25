package main

import "fmt"

func main()  {
	for i := 1; i <= 100; i++ {
		// if
		if(i % 2 == 0){
			fmt.Printf("%d-偶数\n", i)
		}else {
			fmt.Printf("%d-奇数\n", i)
		}
		// switch
		switch i % 2 {
			case 0:
				fmt.Printf("%d-偶数\n", i)
			case 1:
				fmt.Printf("%d-奇数\n", i)
		}
	}

}