package main

import "fmt"

type Handler func(map[string]string) string

func BaseHandler(req map[string]string) string {
	return fmt.Sprintf("Processed request for %s", req["user"])
}

func WithLogging(next Handler) Handler {
	return func(req map[string]string) string {
		fmt.Println("Logging:", req)
		return next(req)
	}
}

func WithAuth(next Handler) Handler {
	return func(req map[string]string) string {
		if req["user"] == "" {
			panic("Unauthorized")
		}
		return next(req)
	}
}

func WithRateLimit(next Handler) Handler {
	return func(req map[string]string) string {
		fmt.Println("Rate limit check passed")
		return next(req)
	}
}

func main() {
	handlers := WithLogging(WithAuth(WithRateLimit(BaseHandler)))

	req := map[string]string{"user": "Rushikesh"}

	fmt.Println(handlers(req))
}
