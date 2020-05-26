package main

import (
	"fmt"
	"github.com/n3wscott/gogo/a"
	"github.com/n3wscott/gogo/b/v2"
)

func main() {
	aa := &a.Pusher{}
	bb := &b.Puller{}
	fmt.Printf("%s %s", aa, bb)
}