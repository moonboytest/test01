package main

import (
	"fmt"
	"module10/db"
	"module10/search"
)

//Список подозреваемых
var Suspects = []string{"Roman", "Sergay", "Sasha", "Matvey", "Bobrb"}

func main() {
	fmt.Println(search.FindTheCriminal(db.Peoples, Suspects))
}
