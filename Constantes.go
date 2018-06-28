package main

import "fmt"

func main() {

	/*
	* Las siguientes sentencias hacen todas exactamente lo mismo:
	* str:="Hello"
	* var str="Hello"
	* var str string="Hello"
	* sin embargo trabajan de distinta manera.
	*/
	
	//Esto es una constante de string sin tipo.
	const hello = "Hello, 世界"

	//Esto también.
	"Charly"

	//Esto es una constante de string de tipo string	
	const typedHello string = "Hello, 世界"

	//Esto es correcto
	var s string
	s = typedHello
	fmt.Println(s)

	//Esto no
	type MyString string
	var m MyString
	m = typedHello // Type error
	fmt.Println(m)

	//Correcto
	m=MyString(typedHello)

	type MyFloat64 float64
    const Zero = 0.0
    const TypedZero float64 = 0.0
    var mf MyFloat64
    mf = 0.0       // OK
    mf = Zero      // OK
    mf = TypedZero // Bad
    fmt.Println(mf)

	
	Pi    = 3.14159265358979323846264338327950288419716939937510582097494459

}

