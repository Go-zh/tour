// +build no-build OMIT

package main

import "github.com/Go-zh/tour/pic"

type Image struct{}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
