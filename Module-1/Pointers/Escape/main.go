package main

// user represents a user in the system.
type user struct {
	name  string
	email string
}

// main is the entry point for the application.
func main() {
	u1 := createUserV1()
	u2 := createUserV2()

	println("u1", &u1, "u2", u2)
}

// createUserV1 creates a user value and passed
// a copy back to the caller.
//go:noinline
func createUserV1() user {
	u := user{
		name:  "Bill",
		email: "bill@ardanlabs.com",
	}

	println("V1", &u)

	return u
}

// createUserV2 creates a user value and shares
// the value with the caller.
//go:noinline
func createUserV2() *user {
	u := user{
		name:  "Bill",
		email: "bill@ardanlabs.com",
	}

	println("V2", &u)

	return &u
}