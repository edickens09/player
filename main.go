package main

import (
	"bytes"
	"time"
	"os"

	"github.com/pion/opus"
	"github.com/pion/opus/oggreader"
	"github.com/gopxl/beep/v2"
	"github.com/gopxl/beep/v2/speaker"

)

func main () {
	f, err := os.Open("./sample4.opus")

	if err != nil {
		log.Fatal(err)
	}
}
