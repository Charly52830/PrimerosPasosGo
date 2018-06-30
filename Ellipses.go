package main
 
import "fmt"
 
func demoEllipsis(arg ...int) {
    fmt.Println("Length of Arguments: ",len(arg))
    fmt.Printf("Arguments: %+v\n",arg)
}
 
func main() {
    fmt.Println("\nExample I")
    demoEllipsis([]int{1,2,3}...)
     
    fmt.Println("\nExample II")
    demoEllipsis(1,2,3,4,5)
     
    fmt.Println("\nExample III")
    demoEllipsis()
     
    fmt.Println("\nDynamic Array")
    intArray := [...] int {1,2,3,4,5,6,7}
    fmt.Println("Length: ",len(intArray))
    fmt.Println("Capacity: ",cap(intArray))
    fmt.Println(intArray)
     
    fmt.Println("\nInterface")
    interfaceEx := []interface{}{"Australia", "Canada", "Japan"}    
    fmt.Println(interfaceEx...)
}
