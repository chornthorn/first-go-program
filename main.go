package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {

	// declaration of a variable
	var a int // var [name] [type] (default value is 0)
	fmt.Println(a)

	// Assignment and Reassignment
	a = 10 // assignment
	fmt.Println(a)
	a = 20 // reassignment
	fmt.Println(a)

	// Declaration and Assignment
	var b int = 30 // var [name] [type] = [value]
	fmt.Println(b)

	// Declaration + Assignment
	var i2 = 30 // var [name] = [value]
	i3 := 30    // [name] := [value]
	fmt.Println(i2, i3)

	// Declaration and Assignment (shorthand)
	c := 40 // [name] := [value]
	fmt.Println(c)

	// Multiple Declaration and Assignment
	var (
		d = 50
		e = 60
	)
	fmt.Println(d, e)

	// Multiple Declaration and Assignment (shorthand)
	f, g := 70, 80
	fmt.Println(f, g)
	var h, j int = 90, 100
	fmt.Println(h, j)

	// Constants are placeholders for values that are not supposed to change
	const k int = 110
	fmt.Println("Constant k:", k)
	const (
		l = 120
		m = 130
		s
	)
	fmt.Println("Constant l:", l)
	fmt.Println("Constant m:", m)
	fmt.Println("Constant s:", s)

	const (
		n = iota + 5
		o
		p
	)
	//n = 5 // error - cannot assign to n
	fmt.Println("Constant n:", n, "o:", o, "p:", p)

	// Primitive Data Types In Go
	/*
		1. bool: true, false
		2. string: string is a collection of bytes
		3. int: int8, int16, int32, int64
	*/
	var (
		isActive bool   = true
		name     string = "John"
		age      int    = 25
	)
	fmt.Println("Name:", name, "Age:", age, "Is Active:", isActive)

	// Type Convertable Data Types
	var (
		pi      float32 = 3.14
		pi2     float64 = float64(pi)
		integer int     = int(pi)
	)
	fmt.Println("pi:", pi, "pi2:", pi2, "integer:", integer)

	// Type Declaration
	type myType int
	var myVar myType
	fmt.Printf("myVar type: %T value: %v\n", myVar, myVar)

	/*
	*	Categories of operators in Go
	*	1. Logical Operators (&&, ||, !)
	*	2. Relational Operators (==, !=, >, <, >=, <=)
	*	3. Arithmetic Operators (+, -, *, /, %)
	*	4. Bitwise Operators (&, |, ^, <<, >>)
	*	5. Assignment Operators (=, +=, -=, *=, /=, %=, &=, |=, ^=, <<=, >>=)
	*	6. Increment/Decrement Operators (++, --)
	*	7. Ternary Operator (?:)
	*	8. Pointer Operator (*)
	*	9. Address Operator (&)
	*	10. Sizeof Operator (sizeof)
	*	11. Comma Operator (,)
	 */

	// 1. Logical Operators
	var (
		condition1 bool = true
		condition2 bool = false
	)
	fmt.Println("condition1:", condition1, "condition2:", condition2)
	fmt.Println("condition1 && condition2:", condition1 && condition2)
	fmt.Println("condition1 || condition2:", condition1 || condition2)
	fmt.Println("!condition1:", !condition1)

	// 2. Relational Operators
	fmt.Println("condition1 == condition2:", condition1 == condition2)
	fmt.Println("condition1 != condition2:", condition1 != condition2)
	fmt.Println("One > Two:", 1 > 2)
	fmt.Println("One < Two:", 1 < 2)
	fmt.Println("One >= Two:", 1 >= 2)
	fmt.Println("One <= Two:", 1 <= 2)

	// 3. Arithmetic Operators
	fmt.Println("One + Two:", 1+2)
	fmt.Println("One - Two:", 1-2)
	fmt.Println("One * Two:", 1*2)
	fmt.Println("One / Two:", 1/2)
	fmt.Println("One % Two:", 1%2)

	// 4. Bitwise Operators
	fmt.Println("One & Two:", 1&2)   //& - AND operator is used to check if both bits are 1 or not (1&1 = 1, 1&0 = 0, 0&1 = 0, 0&0 = 0)
	fmt.Println("One | Two:", 1|2)   //| - OR operator is used to check if any of the bits is 1 or not (1|1 = 1, 1|0 = 1, 0|1 = 1, 0|0 = 0)
	fmt.Println("One ^ Two:", 1^2)   //^ - XOR operator is used to check if both bits are different or not (1^1 = 0, 1^0 = 1, 0^1 = 1, 0^0 = 0)
	fmt.Println("One << Two:", 1<<2) //<< - Left Shift operator is used to shift the bits of the first operand to the left by the number of bits specified by the second operand (1<<2 = 100)
	fmt.Println("One >> Two:", 1>>2) //>> - Right Shift operator is used to shift the bits of the first operand to the right by the number of bits specified by the second operand (1>>2 = 0)

	// 5. Assignment Operators
	var (
		one int = 1
		two int = 2
	)
	fmt.Println("One:", one, "Two:", two)
	one += two // one = one + two
	fmt.Println("One += Two:", one)
	one -= two // one = one - two
	fmt.Println("One -= Two:", one)
	one *= two // one = one * two
	fmt.Println("One *= Two:", one)
	one /= two // one = one / two
	fmt.Println("One /= Two:", one)
	one %= two // one = one % two
	fmt.Println("One %= Two:", one)
	one &= two // one = one & two
	fmt.Println("One &= Two:", one)
	one |= two // one = one | two
	fmt.Println("One |= Two:", one)
	one ^= two // one = one ^ two
	fmt.Println("One ^= Two:", one)
	one <<= two // one = one << two
	fmt.Println("One <<= Two:", one)
	one >>= two // one = one >> two
	fmt.Println("One >>= Two:", one)

	fmt.Println("------------------")

	// 6. Increment/Decrement Operators
	var (
		three int = 3
		four  int = 4
	)
	fmt.Println("Three:", three, "Four:", four)
	three++ // three = three + 1
	fmt.Println("Three++:", three)
	three-- // three = three - 1
	fmt.Println("Three--:", three)
	four++ // four = four + 1
	fmt.Println("Four++:", four)
	four-- // four = four - 1
	fmt.Println("Four--:", four)

	fmt.Println("------------------")

	// 7. Ternary Operator
	var (
		condition3 bool = true
	)
	if condition3 {
		fmt.Println("Condition 3 is true")
	} else {
		fmt.Println("Condition 3 is false")
	}

	// 8. Pointer Operator
	var (
		pointer *int = &three
	)
	fmt.Println("Pointer:", pointer)
	fmt.Println("Pointer value:", *pointer)

	fmt.Println("------------------")

	// 9. Address Operator
	fmt.Println("Address of three:", &three)

	fmt.Println("------------------")

	// 10. Sizeof Operator
	fmt.Println("Size of int:", unsafe.Sizeof(three))

	fmt.Println("------------------")

	// 11. Comma Operator
	var (
		five int = 5
		six  int = 6
	)
	fmt.Println("Five:", five, "Six:", six)
	five, six = six, five // swap values of five and six using comma operator (comma operator evaluates each expression from left to right and returns the value of the last expression)
	fmt.Println("Five:", five, "Six:", six)

	fmt.Println("------------------")

	// 12. Type Conversion
	var (
		seven int = 7
		eight int = 8
	)
	fmt.Println("Seven:", seven, "Eight:", eight)
	fmt.Println("Seven + Eight:", seven+eight)
	fmt.Println("Seven + Eight:", float64(seven)+float64(eight))

	fmt.Println("------------------")

	// 13. Type Assertion
	var (
		condition4 interface{} = true // interface{} is a type that can hold any type of value
	)
	fmt.Println("Condition 4:", condition4)
	fmt.Println("Condition 4 type:", reflect.TypeOf(condition4))
	fmt.Println("Condition 4 value:", reflect.ValueOf(condition4))
	fmt.Println("Condition 4 value:", reflect.ValueOf(condition4).Bool())
	condition4 = 9
	fmt.Println("Condition 4:", condition4)

	fmt.Println("------------------")

	// if else
	var (
		condition5 bool = true
	)
	if condition5 {
		fmt.Println("Condition 5 is true")
	} else {
		fmt.Println("Condition 5 is false")
	}

	// else if
	var (
		condition6 int = 6
	)
	if condition6 == 1 {
		fmt.Println("Condition 6 is 1")
	} else if condition6 == 2 {
		fmt.Println("Condition 6 is 2")
	} else {
		fmt.Println("Condition 6 is not 1 or 2")
	}

	fmt.Println("------------------")

	// switch
	var (
		condition7 int = 7
	)
	switch condition7 {
	case 1:
		fmt.Println("Condition 7 is 1")
	case 2:
		fmt.Println("Condition 7 is 2")
	default:
		fmt.Println("Condition 7 is not 1 or 2")
	}

	// switch with fallthrough
	var (
		condition8 int = 2
	)
	switch condition8 {
	case 1:
		fmt.Println("Condition 8 is 1")
	case 2:
		fmt.Println("Condition 8 is 2")
		fallthrough
	case 3:
		fmt.Println("Condition 8 is 3")
	default:
		fmt.Println("Condition 8 is not 1, 2 or 3")
	}

	// switch with no condition
	var (
		condition9 int = 2
	)
	switch {
	case condition9 == 1:
		fmt.Println("Condition 9 is 1")
	case condition9 == 2:
		fmt.Println("Condition 9 is 2")
	default:
		fmt.Println("Condition 9 is not 1 or 2")
	}

	// switch with type
	var (
		condition10 interface{} = "10"
	)
	switch condition10.(type) {
	case int:
		fmt.Println("Condition 10 is an int")
	case string:
		fmt.Println("Condition 10 is a string")
	default:
		fmt.Println("Condition 10 is not an int or a string")
	}

	fmt.Println("------------------")

	// for loop
	for i := 0; i < 5; i++ {
		fmt.Println("i:", i)
	}

	// for loop with break
	for i := 0; i < 5; i++ {
		if i == 3 {
			break
		}
		fmt.Println("i:", i)
	}

	fmt.Println("------------------")

	// for loop with continue
	for i := 0; i < 5; i++ {
		if i == 3 {
			continue // skip the rest of the code in the current iteration and continue with the next iteration
		}
		fmt.Println("i:", i)
	}

	fmt.Println("------------------")

	// for loop with goto
	for i := 0; i < 5; i++ {
		if i == 3 {
			goto label // jump to the label
		}
		fmt.Println("i:", i)
	}
label:
	fmt.Println("here we are at the label")

	fmt.Println("------------------")

	// for loop with range
	var (
		slice []int = []int{1, 2, 3, 4, 5} // slice is a dynamically sized, flexible view into the elements of an array
	)
	for i, v := range slice {
		fmt.Println("i:", i, "v:", v)
	}

	// for loop with range and _
	var (
		slice2 []int = []int{1, 2, 3, 4, 5} // slice is a dynamically sized, flexible view into the elements of an array
	)
	for _, v := range slice2 {
		fmt.Println("v:", v)
	}

	// for loop with range and only index
	var (
		slice3 []int = []int{1, 2, 3, 4, 5}
	)
	for i := range slice3 {
		fmt.Println("i:", i)
	}

	fmt.Println("------------------")

	// for loop with range and map
	var (
		myMap map[string]interface{} = map[string]interface{}{
			"key1": "value1",
			"key2": 2,
			"key3": true,
		}
	)
	for k, v := range myMap {
		fmt.Println("k:", k, "v:", v)
	}

	fmt.Println("------------------")

	// Functions
	fmt.Println("Sum:", sum(1, 2))
	fmt.Println("Sum:", sum(1, 2, 3))
	fmt.Println("Sum:", sum(1, 2, 3, 4))
	newFunc := sum // assign the function to a variable
	fmt.Println("Sum:", newFunc(1, 2, 3, 4, 5))

	fmt.Println("------------------")

	// Functions as parameters
	var (
		func1 func(int, int) int = func(a int, b int) int {
			return a + b
		}
		func2 func(int, int) int = func(a int, b int) int {
			return a - b
		}
		func3 func(string, bool) string = func(a string, b bool) string {
			if b {
				return a
			}
			return ""
		}
	)
	fmt.Println("Func1:", func1(1, 2))
	fmt.Println("Func2:", func2(1, 2))
	fmt.Println("Func3:", func3("Hello", true))

	fmt.Println("------------------")
	stateCallback(func(value string) {
		fmt.Println("Callback:", value)
	}, "Good morning! :) ")

}

// Functions
func sum(numbers ...int) int {
	result := 0
	for _, v := range numbers {
		result += v
	}
	return result
}

func stateCallback(onChanged func(string2 string), value string) {
	onChanged(value)
}
