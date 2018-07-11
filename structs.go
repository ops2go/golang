
1
2
3
4
5
6
7
8
9
10
11
12
13
14
15
16
17
18
19
20
21
22
23
24
25
26
27
28
29
30
31
32
33
34
35
36
37
38
package main
 
import "fmt"
 
// STRUCTS
 
func main() {
 
// Define a rectangle
rect1 := Rectangle{leftX: 0, TopY: 50, height: 10, width: 10}
 
// Leave off attribute names if we know the order
// rect1 := Rectangle{0, 50, 10, 10}
 
// We access values with the dot operator
fmt.Println("Rectangle is", rect1.width, "wide")
 
// Call the method area for Rectangle
fmt.Println("Area of the rectangle =", rect1.area())
 
}
 
// We can define our own types using struct
type Rectangle struct{
	leftX float64
	TopY float64
	height float64
	width float64
}
 
// We can define methods for our Rectangle by adding the receiver
// rect *Rectangle between func and the function name so we can 
// call it with the dot operator
func (rect *Rectangle) area() float64{
 
	return rect.width * rect.height
 