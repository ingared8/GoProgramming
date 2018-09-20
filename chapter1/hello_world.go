package main

import "fmt"
import "math"

const s string = "constant"

func hello_world() {
    fmt.Println("hello world")
}

func print_values() {

    fmt.Println("go" + "lang")

    //Integers and floats
    fmt.Println("1+1 =", 1+1)
    fmt.Println("7.0/3.0 =", 7.0/3.0)

    //Booleans, with boolean operators as you’d expect.

    fmt.Println(true && false)
    fmt.Println(true || false)
    fmt.Println(!true)

}

func print_variables() {

     var a string = "initial"
     fmt.Println(a)

    var b, c int = 1, 2
    fmt.Println(b, c)

    var d = true
    fmt.Println(d)

    var e int
    fmt.Println(e)

    var x uint64 = 212324234
    var y int64 = 234784535342

    fmt.Println("Unsigned Integers is", x , "Signed Integer is ",  y)
}

func print_constants() {

     fmt.Println(s)

    const n = 500000000

    const d = 3e20 / n
    fmt.Println(d)

    fmt.Println(int64(d))

    fmt.Println(math.Sin(n))
}

func main() {
    //hello_world()
    //print_values()
    //print_variables()

    print_constants()
}