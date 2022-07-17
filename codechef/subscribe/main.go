package main

import "fmt"

func main() {
	var nxt int
	fmt.Scan(&nxt)
	for i := 0; i < nxt; i++ {
		var j int
		fmt.Scan(&j)

		if j <= 30 {
			print("NO")
		} else {
			print("YES")
		}
	}
}
/*
t = int(input())
while t:
n = int(input())
if n <= 30:
print('NO')
else:
print('YES')
t -= 1
*/
