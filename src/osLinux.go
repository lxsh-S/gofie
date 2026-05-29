package src

import (
	"os"
)

// in util.go we have the GetSystemKey func that will extract the value of "PRETTY_NAME" and will return it
func GetOSName() string {
	return GetSystemKey("/etc/os-release", "PRETTY_NAME")
}

// returning username
func GetUsername() string {
	username := os.Getenv("USER") // we know about $USER right? :)
	if username == "" {           // I dont think $USER will be blank but just so it doesnt crash
		return "user"
	}
	return username
}
