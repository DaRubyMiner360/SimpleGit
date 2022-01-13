wpackage main

import (
	"utils"
	_ "commands"

	"github.com/TwiN/go-color"

    "fmt"
	// "flag"
	"os"
    "log"
)

func main() {
	args := os.Args[1:]

	if len(args) > 0 {
		out, errout, err := utils.ExecuteCommand(args)

		if out != "" {
			if err != nil {
				log.Printf("Error: %v\n", err)
			}

			fmt.Println(out)
			if errout != "" {
				fmt.Println(color.Ize(color.Red, errout))
			}
		}
	} else {
		fmt.Println(utils.GetHelp())
	}
}
