package main

import "fmt"

func main() {

	//Slices are immutable, Appending a slice will give you new slice.
	// Case 1: Slice of Slice

	//Create a slice.
	testSlice := []string{"one", "two", "three", "four", "five"}
	fmt.Println(testSlice)
	fmt.Println("Length of slice ", len(testSlice))
	fmt.Println("Capacity  of slice ", cap(testSlice))

	// output
	// [one two three four five]
	// Length of slice  5
	// Capacity  of slice  5

	//Now Create a slice of slice.
	testSlice = testSlice[0:2]
	fmt.Println(testSlice)
	fmt.Println("Length of slice ", len(testSlice))
	fmt.Println("Capacity  of slice ", cap(testSlice))

	//output
	// [one two]
	// Length of slice  2
	// Capacity  of slice  5

	// We can see that capacity of the slice remains the same. But only length has changed.
	// This means that original slice is still unmodified.
	// Slicing here refers to changing of length.

	//Lets increase the length of slice.

	testSlice = testSlice[0:4]
	fmt.Println(testSlice)
	fmt.Println("Length of slice ", len(testSlice))
	fmt.Println("Capacity  of slice ", cap(testSlice))

	//output
	// [one two three four]
	// Length of slice  4
	// Capacity  of slice  5

	//Now lets change the value of the slice and lets see weather it is reflected in original array.

	testSlice[1] = "two" //Assign a new value.

	// lets expand the slice back to original.
	testSlice = testSlice[:5]
	fmt.Println(testSlice)
	fmt.Println("Length of slice ", len(testSlice))
	fmt.Println("Capacity  of slice ", cap(testSlice))

	//output
	// [one two three four five]
	// Length of slice  5
	// Capacity  of slice  5

	// case 2: Creating slice with make
	// Format:
	// s := make(sliceName, int len, int capacity)

	makeSlice := make([]string, 5, 5)
	fmt.Println(makeSlice)
	fmt.Println("Length of slice ", len(makeSlice))
	fmt.Println("Capacity  of slice ", cap(makeSlice))

	//output
	// [    ]
	// Length of slice  5
	// Capacity  of slice  5

	makeSlice[0] = "zero"
	makeSlice[1] = "one"
	makeSlice[2] = "two"
	makeSlice[3] = "three"
	makeSlice[4] = "four"

	makeSlice = makeSlice[:3]

	for i := 0; i < len(makeSlice); i++ {
		fmt.Println(makeSlice[i])
	}

	//output
	// zero
	// one
	// two

	//Even through its capacity is 5 and its has five values but it printed only three values.
	// Basically slice maintains two variables
	// Capacity : maximum length that array can hold.
	// length : current restricted length for array can be operated.

	//lets try to print variable at 4.

	//fmt.Println(makeSlice[4])
	// Output
	// runtime error: index out of range [4] with length 3

	//Lets increase the capacity and try to print.

	makeSlice = makeSlice[:5]
	fmt.Println(makeSlice[4])
	// output
	// four

	//Append.: With the help of append, we can create an unlimited list.

	var appendableSlice []int
	fmt.Println(appendableSlice)
	fmt.Println(len(appendableSlice))
	fmt.Println(cap(appendableSlice))

	// output
	// []
	// 0
	// 0

	appendableSlice = append(appendableSlice, 10)
	fmt.Println(appendableSlice)
	fmt.Println(len(appendableSlice))
	fmt.Println(cap(appendableSlice))

	//output
	// [10]
	// 1
	// 1

	appendableSlice = append(appendableSlice, 20)
	fmt.Println(appendableSlice)
	fmt.Println(len(appendableSlice))
	fmt.Println(cap(appendableSlice))

	//output
	// [10 20]
	// 2
	// 2

	appendableSlice = append(appendableSlice, 30, 40, 50)
	fmt.Println(appendableSlice)
	fmt.Println(len(appendableSlice))
	fmt.Println(cap(appendableSlice))

	// output
	// [10 20 30 40 50]
	// 5
	// 6

}
