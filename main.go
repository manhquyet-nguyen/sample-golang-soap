package main

import (
	"fmt"
	"sample-golang-soap-xml/pkg"
)

func main()  {
	server := pkg.NewSOAPServer("9999")
	if server == nil {
		panic("Cant load")
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Server listening on port 9999...")
}
