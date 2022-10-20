package utils

import "log"

func Assert(boo bool) {
	if !boo {
		log.Fatalf("NOPE")
	}
}
