package main

import (
	"fmt"
	"strconv"
	"strings"
)

type IPAddr [4]byte

func (a IPAddr) String() string {
	s := make([]string, 4)
	for i := range a {
		s[i] = strconv.Itoa(int(a[i]))
	}
	return strings.Join(s, ".")
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
