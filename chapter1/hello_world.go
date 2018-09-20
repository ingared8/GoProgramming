package main

import "fmt"
import "math"
import "time"

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

func prnt(vars interface{}) {
    fmt.Println(vars)
}


func for_loop() {

    i := 1

    for i <= 3 {
        prnt(i)
        i += 1
    }

    for n := 6; n <= 10 ; n++ {

     if n%2 == 0{
        continue
     }

       prnt(n)
    }

}


func if_else() {

    if 7%2 == 0 {
            fmt.Println("7 is even")
        } else {
            fmt.Println("7 is odd")
        }

     //You can have an if statement without an else.

        if 8%4 == 0 {
            fmt.Println("8 is divisible by 4")
        }

    //A statement can precede conditionals; any variables declared in this statement are available in all branches.

        if num := 9; num < 0 {
            fmt.Println(num, "is negative")
        } else if num < 10 {
            fmt.Println(num, "has 1 digit")
        } else {
            fmt.Println(num,"has nultiple digits")
        }
}




func main() {

    //hello_world()

    //print_values()

    //print_variables()

     //print_constants()

     //for_loop()

     if_else()
}
