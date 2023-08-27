package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

/*
Simple go function to kick off
*/
func main() {
	fmt.Printf("Hello World!!")
	fmt.Println()
	// go variables
	var x int
	x = 20

	//or
	y := 20

	fmt.Println("sum", x+y) //proper way for printing stuff

	// type alias

	type idNumber int

	var id idNumber = 100

	fmt.Println("My Id number is:", id)

	// short hand notation for declare and init

	name := "Piyush"
	fmt.Print("My Name is " + name) //same as java to std out.

	// Pointers :: stay the fuck away from this stuff

	var addr *int // this will hold address of some int.

	var place int = 100

	//this works only if place is int32. other type it wont work.
	// default int is int32
	addr = &place
	fmt.Println("Address value:", *addr) // * to deref.

	// New , its alternate way of creating stuff.

	var myHome = new(string) // same as var myHome *string

	*myHome = "Solaris"
	fmt.Println("Home value", *myHome) // wtf and why bother.

	test := foo()

	fmt.Println("old home:", *test)

	// type conversion

	var typeA int32 = 100
	var typeB int64 = 200

	typeA = int32(typeB) // works with cast only, else compiler error
	typeB = int64(typeA) //both ways u need type conversion, bigger to smaller and vis

	fmt.Println(typeA, typeB)

	// strings are readonly here.

	chota := "Rakesh"
	pura := chota + "Choudhary"
	fmt.Println("Pura naam :", pura)

	//strings functions.
	fmt.Println("Comparision:", strings.Compare(chota, pura)) // 0 if chota == pura, -1 if chota < pura

	fmt.Println("Contains : ", strings.Contains("ABC", "BC"))
	fmt.Println("Index of : ", strings.Index("ABC", "BC"))     // strings are zero indexed
	fmt.Println("Contains : ", strings.HasSuffix("ABC", "BC")) // there is hasPrefix also.

	// string manipulation.

	fmt.Println(strings.Replace("ABCDEFBC", "BC", "A", 2)) // out: AADEFA, 2 occurances of BC replaced with A
	fmt.Println(strings.ToLower("ABdcerD"))
	fmt.Println(strings.TrimSpace("rakesh choudhary   "))

	// strconv functions
	val, err := strconv.Atoi("2")
	str := strconv.Itoa(2)

	if err == nil {
		fmt.Println(val, str)
	} else { // formatting is a bitch here, else must appear in same line as closing bracket of if
		fmt.Println("failed to convert string to number")
	}

	fl, err := strconv.ParseFloat("222.23343434", 2)
	fmt.Println(fl)

	// constants

	const t = "Final" // cant be changed, also var cant be mixed here, remember var is for variables only

	type WeekDays float32 //iota must be numeric int float etc, string is not allowed
	const (
		MONDAY    WeekDays = iota
		TUESDAY   WeekDays = iota
		WEDNESDAY WeekDays = iota // type is not needed on all
	)

	fmt.Println(WEDNESDAY)

	// for loop
	for i := 0; i < 10; i++ {
		fmt.Println("It will run till i is less than 10", i)
	}
	i := 0
	for i < 10 {
		fmt.Println("It will run till i is less than 10", i)
		i++
	}

	//for without i is infinite loop

	for {
		// i is set to zero on each iteration remember. so its infinite
		i := 0

		fmt.Println("Running for", i)
		i++
		fmt.Println("Running for", i)

		if i == 1 {
			fmt.Println("Going to Exit from here!!!")
			break // break will exit from here
		}

	}

	cond := 10

	switch cond {
	case 1:
		fmt.Println(1)
	case 10:
		fmt.Println("Yey")
	default:
		fmt.Println("Default")
	}

	//switch without tag

	switch {
	case true:
		fmt.Println("Always run this")
	case false:
		fmt.Print("Will not run this")
	default:
		fmt.Println("This also wont run")
	}

	//take input

	var input string
	fmt.Println("Enter a text::")
	fmt.Scanln(&input)
	//scan or scanln split the input at spaces..
	fmt.Println(input)

	var fName, mName, lName string
	fmt.Println("Enter your name")
	fmt.Scanln(&fName, &mName, &lName)

	fmt.Println("Welcome", fName, mName, lName)

	// arrays.
	var arr [5]int // proper declation. zero initiatlized values
	arr[0] = 1
	arr[2] = 4
	fmt.Println(arr)

	//predefined . like java::
	var arr2 [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr2)

	// more elegant way is this::
	var arr3 = [...]int{1, 2, 3}
	fmt.Println(arr3)
	//iterating on array::
	for i, v := range arr {
		fmt.Println(i, v) //i is index
	}

	//range loop generates two values i and v what if u passed one
	for val := range arr {
		fmt.Println(val) //prints index, because sequentially val is mapped to index
	}

	//can we change value while iterating??
	for i, v := range arr3 {
		i = 2
		v = 5
		fmt.Println(i, v) // i and v are local vars not the i and v from for loop
	}

	fmt.Println(arr3) //arr remained intact after above for loop
	//but what if above was run like arr3[i]= 5, then it will update the slice
	for i, v := range arr3 {
		arr3[i] += v      //its adding v to v, so doubling the value now.
		fmt.Println(i, v) // i and v are local vars not the i and v from for loop
	}

	fmt.Println(arr3) //arr3 is updated now.

	//slices.
	fmt.Println("=========Slices========")
	s := arr3[0:2]                 // its excluding 2nd index
	fmt.Println(len(s), cap(s), s) //cap is capacity of slice
	s1 := arr3[0:]
	fmt.Println(cap(s1)) //capacity is 3 . size of array basically.
	s1[0] = 10
	fmt.Println(arr3, s1, s) //updates all slices of array and array itself.

	//slice can be created without size.
	s2 := []int{1, 2, 3, 4, 5}
	fmt.Println(cap(s2))
	s3 := make([]int, 10) //10 is size.
	fmt.Println(cap(s3))

	s4 := make([]int, 5, 10)
	fmt.Println(cap(s4), len(s4)) //capacity is 10 of actual array, slice cap is 5
	s4 = append(s4, 100)          //added at last, making cap of slice to 6
	fmt.Println(s4)

	//maps
	fmt.Println("==================Maps=======")

	var idMap map[string]int //map with strings as key and int as val

	//idMap["joe"] = 1 //this will fail since its empty map

	fmt.Println(idMap)           //empty map, cant add to it
	idMap = make(map[string]int) // afther this we can add in map
	idMap["joe"] = 1
	fmt.Println(idMap)

	//literal maps
	idMap2 := map[string]int{
		"joe": 1, //this is shitty. need to add , to indicate closing of entry wtf
	}

	fmt.Println(idMap2)

	for key, val := range idMap2 {
		fmt.Println(key, val)
	}

	//assignment
	key, p := idMap2["len"]
	fmt.Println(key, p) //p is false if len key is missing in map

	//objects.

	fmt.Println("=====json marshal and unmarshal=======")

	bool1, _ := json.Marshal(true)
	fmt.Println(string(bool1)) //byte array to string, same as java.

	var bool2 bool
	err2 := json.Unmarshal(bool1, &bool2)

	fmt.Println(bool2, err2) //unmarshal back

	type boy struct {
		Name string
		Age  int
	}

	boy1 := boy{Name: "Piyush", Age: 12}
	fmt.Println(boy1)
	m_struct, _ := json.Marshal(boy1)
	fmt.Println(string(m_struct))

	var piyush boy
	m_err := json.Unmarshal(m_struct, &piyush)
	fmt.Println(piyush, m_err)

	//files

	fmt.Println("===Files===")
	//this way of reading files is poor and deprecated
	fileByteArr, _ := ioutil.ReadFile("Test.txt")
	fmt.Println(string(fileByteArr))

	//keeps formatting intact but its blocking like readFile and deprecated
	ioutil.WriteFile("Out.txt", fileByteArr, 0777)

	//os package is recommended way

	fileRef, _ := os.Open("Out.txt")
	fileBytes := make([]byte, 10)
	bytesRead, _ := fileRef.Read(fileBytes)
	fmt.Println("read", string(fileBytes), "bytes read:", bytesRead)

	//writes with os

	fo, _ := os.Create("Out2.txt")
	fo.WriteString("Slytherin house was not all bad!!")

	bytesToWrite := []byte{1, 2, 3}
	fo.Write(bytesToWrite)

}

//function pointers also exist here. damn

func foo() *string {
	myHome := new(string)
	*myHome = "Mahaveer Apt"
	return myHome
}
