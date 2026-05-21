package main

import "github.com/rahulkumarpahwa/webhealthchecker/internals/checker"

func main() {

	checker := checker.Check{
		Domain: "google.com",
		Port:   "80",
	}
	checker.Checker()

}
