/*
Prime Game 
Rax, a school student, was bored at home in the pandemic. He wanted to play but there was no one to play with. He was doing some mathematics questions including prime numbers and thought of creating a game using the same. After a few days of work, he was ready with his game. He wants to play the game with you.

GAME:
Rax will randomly provide you a range [ L , R ] (both inclusive) and you have to tell him the maximum difference between the prime numbers in the given range. There are three answers possible for the given range.
There are two distinct prime numbers in the given range so the maximum difference can be found.
There is only one distinct prime number in the given range. The maximum difference in this case would be 0.
There are no prime numbers in the given range. The output for this case would be -1.
To win the game, the participant should answer the prime difference correctly for the given range.

Example:
Range: [ 1, 10 ]
The maximum difference between the prime numbers in the given range is 5.
Difference = 7 - 2 = 5

Range: [ 5, 5 ]
There is only one distinct prime number so the maximum difference would be 0.

Range: [ 8 , 10 ]
There is no prime number in the given range so the output for the given range would be -1.

Can you win the game?

Input Format
The first line of input consists of the number of test cases, T
Next T lines each consists of two space-separated integers, L and R

Constraints
1<= T <=10
2<= L<= R<=10^6

Output Format
For each test case, print the maximum difference in the given range in a separate line. 

Sample TestCase 1
Input
5
5 5
2 7
8 10
10 20
4 5

Output
0
5
-1
8
0

*/

package main

import (
	"fmt"
	"math"
)

func main() {
	n := 0
	min := 0
	max := 0

	_, _ = fmt.Scan(&n)
	for j := 0; j < n; j++ {
		_, _ = fmt.Scan(&min, &max)

		k := -1
		l := -1

		for i := min; i <= max; i++ {
			if isPrime(i) {
				if k < 0 {
					k = i
				} else {
					l = i
				}
			}
		}

		if k < 0 {
			fmt.Println(-1)
		} else if l < 0 {
			fmt.Println(0)
		} else {
			fmt.Println(l - k)
		}
	}
}

func isPrime(a int) bool {
	if a != 2 && a%2 == 0 {
		return false
	}

	for i := 2; i <= int(math.Sqrt(float64(a))); i++ {
		if a%i == 0 {
			return false
		}
	}
	return true
}

/*
func minimum() int {
	min := arr[0]
	for i, j := range arr {
		if j < min {
			min = arr[i]
		}
		continue
	}
	return min
}

func maximum() int {
	max := arr[0]
	for i, j := range arr {
		if j > max {
			max = arr[i]
		}
		continue
	}
	return max
}
*/
