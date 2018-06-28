package main

import "fmt"

func main() {
	
	/**
	* rune: alias para int32
	* byte: alias para uint8
	* 
	*	Imprimir rune
	* %q	muestra el símbolo UTF8
	* %+q	muestra el código UTF8
	* %v	muestra el valor decimal
	* %x	muestra el valor hexadecimal
	* %U	muestra el unicode code point
	* %#U	muestra el unicode code point y el símbolo UTF8
	*/
	
	var r rune='a'
	fmt.Printf("%q\n",r)
	//fmt.Printf("%+q\n",r)
	fmt.Printf("%v\n",r)
	fmt.Printf("%x\n",r)
	fmt.Printf("%U\n",r)
	fmt.Printf("%#U\n",r)
	println()
	
	/** Sentencias iguales */
	var a rune= 97 			//decimal
	var b rune= 0x61		//hexadecimal
	var c rune= 'a'			//símbolo
	var d rune= '\u0061'	//unicode code point
	
	fmt.Printf("%v %v %v %v\n",a,b,c,d)
	println()
	
	var s string="\u0043\u0068\u0061\u0072\u006c\u0079"
	fmt.Printf("%v\n",s)
	
	str:=[]uint8{67,104,97,114,108,121}
	fmt.Printf("%s\n",str)
	println()
	
	s="\u2318"
	fmt.Printf("%x %v\n\n",s,len(s))
	
	/** for vs for range*/
	const nihongo = "日本語"
	for index, runeValue := range nihongo {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
	}
	println()
	
	for index:=0;index<len(nihongo);index++ {
		fmt.Printf("%#U ",nihongo[index])
	}
	println()
	
}
