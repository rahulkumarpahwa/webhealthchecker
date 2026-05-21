package main

import (
	"github.com/rahulkumarpahwa/webhealthchecker/internals/checker"
)

func main() {

	c := checker.Check{}

	c.AddCheck("apple.com", "80")
	c.AddCheck("microsoft.com", "80")
	c.AddCheck("google.com", "80")
	c.AddCheck("rahulkumarpahwa.me", "80")
	checker.CloseCheck()
	c.Runner()
}
