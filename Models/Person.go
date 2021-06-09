package models

import (
	"strconv"
	"strings"
)

type Person struct {
	Name        string
	Age         int
	Credentials []string
}

func (p Person) String() string {
	return p.Name + " Age: " + strconv.Itoa(p.Age) + " has " + strings.Join(p.Credentials, ",")
}
