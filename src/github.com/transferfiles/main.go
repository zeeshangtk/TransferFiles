package main

import (
	"fmt"

	r "github.com/transferfiles/router"

	"github.com/urfave/negroni"
)

func main() {

	n := negroni.Classic()
	n.UseHandler(r.GetRouter())
	n.Run(":3000")
	fmt.Println("Transfer files is up")

}
