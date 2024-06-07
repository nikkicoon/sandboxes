package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func Hostname() (string, error) {
	// this sucks, but there's no standard way and
	// not calling an external binary involves a lot
	// of code
	// TODO: do this without the binary, which doesn't always come with the f function
	path, err := exec.LookPath("hostname")
	if err != nil {
		return "", fmt.Errorf("hostname binary not found: %v\n", err)
	}
	fmt.Printf("hostname binary found: %s\n", path)
	cmd := exec.Command(path, "-f")
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		return "", fmt.Errorf("getFQDN error: %v\n", err)
	}
	fqdn := out.String()
	fqdn = fqdn[:len(fqdn)-1] // remove EOL
	return fqdn, nil
}

func main() {
	res, err := Hostname()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}
