//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.

package main

import (
	"fmt"
	"math"
)

//* Create a rectangle structure containing its coordinates
type Coordinate struct {
	x float64
	y float64
}

type Rectangle struct {
	topLeft     Coordinate
	bottomRight Coordinate
}

func (r *Rectangle) area() float64 {
	return r.width() * r.height()
}

func (r *Rectangle) perimeter() float64 {
	return r.width()*2 + r.height()*2
}

func (r *Rectangle) doubleInSize() {
	halfWidth := r.width() / 2
	halfHeight := r.height() / 2
	if r.topLeft.x <= r.bottomRight.x {
		r.topLeft.x -= halfWidth
		r.bottomRight.x += halfWidth
	} else {
		r.topLeft.x += halfWidth
		r.bottomRight.x -= halfWidth
	}

	if r.topLeft.y <= r.bottomRight.y {
		r.topLeft.y -= halfHeight
		r.bottomRight.y += halfHeight
	} else {
		r.topLeft.y += halfHeight
		r.bottomRight.y -= halfHeight
	}
}

func (r *Rectangle) width() float64 {
	return math.Abs(r.bottomRight.x - r.topLeft.x)
}

func (r *Rectangle) height() float64 {
	return math.Abs(r.bottomRight.y - r.topLeft.y)
}

func main() {
	rectangle := Rectangle{
		topLeft: Coordinate{
			-1,
			-1,
		},
		bottomRight: Coordinate{
			1,
			3,
		},
	}

	//* Using functions, calculate the area and perimeter of a rectangle,
	//  - Print the results to the terminal
	//  - The functions must use the rectangle structure as the function parameter
	fmt.Println("rectangle", rectangle)
	fmt.Println("area of rectangle", rectangle.area())
	fmt.Println("perimeter of rectangle", rectangle.perimeter())

	//* After performing the above requirements, double the size
	fmt.Println("doubling rectangle")
	rectangle.doubleInSize()

	//  of the existing rectangle and repeat the calculations
	//  - Print the new results to the terminal
	fmt.Println("rectangle", rectangle)
	fmt.Println("area of rectangle", rectangle.area())
	fmt.Println("perimeter of rectangle", rectangle.perimeter())

	//
	//--Notes:
	//* The area of a rectangle is length*width
	//* The perimeter of a rectangle is the sum of the lengths of all sides
}
