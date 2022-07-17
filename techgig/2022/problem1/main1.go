/*
World Army vs Aliens (100 Marks)
The world is going to be attacked by the aliens. The space intelligence department has raised an alarm and the world armies are coming together to fight them. The planning and strategizing is being done to make the maximum impact on the alien ships. The deadly missiles are to be put into action. The missiles are targeted to destroy the alien ships in space and disable them to land on the Earth.

The army is planning to launch the attack at A time (hh mm) to get an advantage. For each attack, they know the time the missile will require to hit the coming alien ship. They all are busy in preparation and want to know the time at which the blast will occur and the alien ship will be destroyed in pieces. Can you find the time of the blast ?

Note: The World Army follows a 24-hour time format and you need to find the time of blast accordingly. The time will be in the hh mm format.

Example:

Input Format
The first line of input consists of the launch time in hh mm format separated by space.

The second line of input consists of the travel time for the missile in hh mm format separated by space.

Constraints
00<= hh <=23

00<= mm <=59

Output Format
Print the time at which the blast will occur and the spaceship will be destroyed.

Sample TestCase 1
Input
19 50
02 20

Output
22 10
 */

package main

import (
	"fmt"
)

func main() {
	var h int
	var m int

	fmt.Scan(&h, &m)

	var hh int
	var mm int

	fmt.Scan(&hh, &mm)
/*
	t := time.Date(0, 0, 0, h, m, 0, 0, time.UTC)
	t2 := t.Add(time.Duration(m) * time.Minute(mm))
	t2 = t.Add(time.Duration(h) * time.Hour(hh))

	fmt.Print(t2)
*/
	var rm int
	if m + mm >= 60 {
		h++
		rm = (m + mm) - 60
	} else {
		rm = m + mm
	}
	 rh := h + hh

	 fmt.Print(rh, rm)
}
