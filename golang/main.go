package main

import (
	"fmt"

	"github.com/anthony-dong/vscode-config/golang/utils"
)

func main() {
	sun := utils.Add(1, 2)
	fmt.Println("sum: ", sun)
}
