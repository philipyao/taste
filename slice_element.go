package main

import(
    "container/list"
    "fmt"
)

var myList *list.List
var myMap map[int]*list.Element

func main() {
    myList = list.New()
    myMap = make(map[int]*list.Element)

    for i := 0; i < 10; i++ {
        element := myList.PushBack(i)
        myMap[i] = element
    }
    for e := myList.Front(); e != nil; e = e.Next() {
        fmt.Printf("element %p, %+v\n", e, e)
    }

    if e, ok := myMap[4]; ok {
        myList.Remove(e)
        delete(myMap, 4)
    }
    fmt.Printf("After remove\n")
    for k, e := range myMap {
        fmt.Printf("k %v, element %p %+v\n", k, e, e)
    }
}
