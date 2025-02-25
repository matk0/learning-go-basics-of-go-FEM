package data

import "fmt"

var Countries [10]string
var Slice []int
var Codes map[int]string

func init() {
    Countries[0] = "India"
    Countries[1] = "USA"
    Countries[2] = "Japan"
    Countries[3] = "Germany"
    Countries[4] = "France"
    Countries[5] = "Italy"
    Countries[6] = "Brazil"
    Countries[7] = "Canada"
    Countries[8] = "Australia"
    Countries[9] = "China"
    
    qty := len(Countries)
    fmt.Println("Length of Countries:", qty)
}
