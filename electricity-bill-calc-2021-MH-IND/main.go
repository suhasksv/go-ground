package main

import (
	"fmt"
)

func main() {
	var units uint32
	fmt.Print("Enter number of units consumed: ")
	fmt.Scan(&units)

	var powerSize = calcSlab(units)
	var fixedSize float32 = 350
	var carryingSize = float32(units) * 1.38
	var electricityCharges = (fixedSize + powerSize + carryingSize) * 0.16

	fmt.Printf("\n\nFixed Size = %.2f\nPower Size = %.2f\nCarrying Size = %.2f\nElectricity charges (16) = %.2f\n", fixedSize, powerSize, carryingSize, electricityCharges)

	var totalPrice = fixedSize + powerSize + carryingSize + electricityCharges
	fmt.Printf("Your need to pay ₹%.2f\n", totalPrice)
	units = 0
}

func calcSlab(units uint32) float32 {
	switch {
	case units <= 100:
		return _calcSlab_0To100u(units)
	case units >= 101 && units <= 300:
		return _calcSlab_101To300u(units)
	case units >= 301 && units <= 500:
		return _calcSlab_301To500(units)
	case units > 500:
		return _calcSlab_501To1000(units)
	}
	return 0
}

func _calcSlab_0To100u(units uint32) float32 {
	var powerSize float32
	powerSize = float32(units) * 3.44
	return powerSize
}

func _calcSlab_101To300u(units uint32) float32 {
	var powerSize = _calcSlab_0To100u(100)
	units -= 100
	powerSize += float32(units) * 7.34
	return powerSize
}

func _calcSlab_301To500(units uint32) float32 {
	var powerSize = _calcSlab_101To300u(300)
	units -= 300
	powerSize += float32(units) * 10.36
	return powerSize
}

func _calcSlab_501To1000(units uint32) float32 {
	var powerSize = _calcSlab_301To500(500)
	units -= 500
	powerSize += float32(units) * 11.82
	return powerSize
}
