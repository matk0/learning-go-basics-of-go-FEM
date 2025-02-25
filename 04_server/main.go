package main

import (
	"fmt"
	"frontendmasters.com/go/server/data"
)

func main()  {
    matej := data.Instructor{ Id: 1, FirstName: "Matej", LastName: "Lukášik", Score: 100 }
    // kyle := data.NewInstructor("Kyle", "Bass")

    goCourse := data.Course{Id: 2, Name: "Go Fundamentals", Instructor: matej}

    fmt.Printf("%v", goCourse)
    // print(matej.Print())
    // print(kyle.Print())

    swiftWs := data.NewWorkshop("Swift with iOS", matej)

    var courses [2]data.Signable
    courses[0] = goCourse
    courses[1] = swiftWs

    for _, course := range courses {
        fmt.Println(course)
    }
}
