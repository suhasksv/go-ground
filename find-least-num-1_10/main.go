// package main

// import "fmt"

// func main() {

// 	var cnt int
// 	var result int

// 	for i := 100; i > 0; i++ {

// 		for j := 2; j <= 10; j ++ {
// 			if i % j == 0 {
// 				result = i
// 				cnt++
// 			} else {
// 				cnt = 0
// 				break
// 			}
// 		}

// 		if cnt == 9 {
// 			fmt.Println(result)
// 			break		
// 		} 
// 	}
// }
package main

import "fmt"

func main() {

	var cnt int

	for i := 1; i > 0; i++ {

		var j int

		for j = 2; j <= 10; j++ {
			if i % j != 0 {
				break
			}
		}

		if j > 10 {
			cnt++
			if cnt>=31 && cnt%3==0 && cnt%5==0 && cnt%6==0 {
				fmt.Printf("i: %d cnt: %d\n", i, cnt)
				break
			}		
		}
	}
}
