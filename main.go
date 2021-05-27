package main

// import (
// 	"log"
// 	"time"
// )

// func main() {
// 	for {
// 		log.Println("hello there")
// 		<-time.After(time.Second * 5)
// 	}
// }

import (
	"fmt"
	"log"
	"time"

	"github.com/dimiro1/banner"
	"github.com/mattn/go-colorable"
)

func main() {
	templ := `{{ .Title "Hello World" "" 0 }}
   GoVersion: {{ .GoVersion }}
   GOOS: {{ .GOOS }}
   GOARCH: {{ .GOARCH }}
   NumCPU: {{ .NumCPU }}
   GOPATH: {{ .GOPATH }}
   GOROOT: {{ .GOROOT }}
   Now: {{ .Now "Monday, 2 Jan 2006" }}`
	banner.InitString(colorable.NewColorableStdout(), true, true, templ)

	fmt.Println()
	fmt.Println()

	for {
		log.Println("hello there: ", time.Now().String())
		<-time.After(time.Second * 5)
	}
}
