package main

import "fmt"

func maps_example() {

    m := make(map[string]int)

    m["k1"] = 7
    m["k2"] = 13


    fmt.Println("map:", m)

    v1, x := m["k1"]
    fmt.Println("v1: ", v1, x)

    fmt.Println("len:", len(m))

    delete(m, "k2")
    fmt.Println("map:", m)

    y, prs := m["k2"]
    fmt.Println("prs:", y, prs)

    //You can also declare and initialize a new map in the same line with this syntax.

    n := map[string]int{"foo": 1, "bar": 2}
    fmt.Println("map:", n)
}

func range_func(){

    nums := []int{2, 3, 4}
    sum := 0

    for _, num := range nums {
        sum += num
    }

    fmt.Println("sum:", sum)

    for i, num := range nums {
        if num == 3 {
            fmt.Println("index:", i)
        }
    }

    kvs := map[string]string{"a": "apple", "b": "banana"}
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)
    }

    for k := range kvs {
        fmt.Println("key:", k)
    }

    for i, c := range "go" {
        fmt.Println(i, c)
    }
}

func main() {
   //maps_example()
   range_func()
}