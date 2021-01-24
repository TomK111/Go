package main

import (
	"fmt"
	"os"
)

const (
	usage     = "Usage: [username] [password]"
	errUser   = "Access denied for %q.\n"
	errPwd    = "Invalid password for %q.\n"
	accessOK  = "Access granted to %q.\n"
	username1 = "jack"
	password1 = "1888"
	username2 = "tommy"
	password2 = "1997"
)

func main() {
	args := os.Args

	if len(args) != 3 {
		fmt.Println(usage)
		return
	} //If the length of arguments provided is not 3, print "Usage: [username] [password["

	u, p := args[1], args[2] // set variables u, p to argument 1 and 2.

	if u != username1 && u != username2 {
		fmt.Printf(errUser, u) //if u does not equal user, print "Access denied for [u]"
	} else if p != password1 && u != password2 {
		fmt.Printf(errPwd, u) //if p does not equal password, print "Invalid password for [u]"
	} else {
		fmt.Printf(accessOK, u) //if u and p are correct, print "Access granted for [u]"
	}
}
