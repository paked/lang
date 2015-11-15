// Command lang interprets and runs
package main

import (
	"fmt"
	"os"

	"github.com/paked/lang"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("lang: requires the file name to be parsed")
		fmt.Println("\texample: lang sample.lang")
		return
	}

	file := os.Args[1]

	f, err := os.Open(file)
	if err != nil {
		fmt.Printf("could not open file '%v': %v\n", file, err)
	}

	defer f.Close()

	l := lang.NewLexer(f)
	p := lang.NewParser(l)

	fmt.Print("parsing...")
	prog := p.Parse()
	fmt.Println(" DONE!")

	fmt.Println("running: ")
	prog.Run()
}
