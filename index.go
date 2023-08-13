//1. Printing formats

// package main
// import "fmt"
// func main()  {
// 	x := 42
// 	y := "James Bond"
// 	z := true
// 	fmt.Println(x,y,z)
// }


// 2. Declare variables

// package main
// import "fmt"
// var (
// 	x int
// 	y int
// 	z int
// )
// func main()  {
// 	fmt.Println(x,y,z)
// }


// 3. Declare variables 2

// package main
// import "fmt"
// var (
//  x = 42
//  y = "James Bond"
//  z = true
// )
// func main() {
//  s := fmt.Sprintf("x: %v, y: %v, z: %v", x, y, z)
//  fmt.Println(s)
// }


// 4. Creating new type

// package main
// import "fmt"
// type myint int
// func main() {
// 	var x myint
// 	fmt.Println(x)
// 	x = 42
// 	fmt.Println(x)
// }


// 5. Printing variable and operator

// package main
// import "fmt"
// type myint int
// var y myint
// func main() {
//  var x myint
//  fmt.Println(x)        
//  fmt.Printf("%T\n", x)
//  x = myint(42)
//  fmt.Println(x) 
//  y = myint(x)
//  fmt.Println(y) 
//  fmt.Printf("%T\n", y)
// }


// 6. Quiz
// 23/36


// 7. Divisible by 7

// package main
// import "fmt"
// func main() {
//     var numbers []int
//     for i := 2000; i <= 3200; i++ {
//         if i%7 == 0 && i%5 != 0 {
//             numbers = append(numbers, i)
//         }
//     }
//     for i, num := range numbers {
//         if i != 0 {
//             fmt.Print(", ")
//         }
//         fmt.Print(num)
//     }
// }

// 8. Compute factorial

package main
import "fmt"
func factorial(n uint) uint{
    if n == 0{
        return 1
    }
    return n * factorial(n - 1)
}
func main() {
    fmt.Println(factorial(8))
}

