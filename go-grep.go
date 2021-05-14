package main


import (
	"flag"
	"fmt"
	"log"
)


var h = flag.Bool("h", false, "go-grep -h")
var help = flag.Bool("help", false, "go-grep -help")
var e = flag.Bool("e", false, "go-grep [-e] [pattern] [file]")
var pattern = flag.String("", "", "go-grep [string] [file]")

var help_message = `
NAME
go-grep

Usage
go-grep [-e] [pattern] [file]

Example
1.
  $ go-grep 'string' file
2.
  $ go-grep -e 'runoo+b' file

`

func main() {
	
	flag.PrintDefaults()

	// Scans the arg list and sets up flags
	flag.Parse() 
	
	if *h || *help{
		fmt.Println(help_message)
	} else {
		if flag.NArg() != 2 {
			log.Fatal("The parameter is not supported, please refer to: go-grep -h")
		}
		
		var strPattern = flag.Arg(0)
		var dataFromFile = ReadFile(flag.Arg(1))

		RunGrep([]byte(strPattern), dataFromFile, *e)
	} 
}



