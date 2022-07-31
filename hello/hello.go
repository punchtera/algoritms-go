package hello

// This is the way of renaming default imports
import "rsc.io/quote/v3"

func Hello() string {
	return quote.HelloV3()
}

func Proverb() string {
	return quote.Concurrency()
}
