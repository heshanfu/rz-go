// +build !binary_log

package diode_test

import (
	"fmt"
	"os"

	"github.com/bloom42/astro-go"
	"github.com/bloom42/astro-go/diode"
)

func ExampleNewWriter() {
	w := diode.NewWriter(os.Stdout, 1000, 0, func(missed int) {
		fmt.Printf("Dropped %d messages\n", missed)
	})
	log := zerolog.New(w)
	log.Print("test")

	w.Close()

	// Output: {"level":"debug","message":"test"}
}