package flags

import (
	"errors"
	"strconv"
	"strings"
)

type NetAddress struct {
	Host string
	Port int
}

type Value interface {
	String() string
	Set(string) error
}

func (a NetAddress) String() string {

	if a.Host == "" {
		a.Host = "localhost"
	}
	if a.Port < 1 {
		a.Port = 8080
	}
	return a.Host + ":" + strconv.Itoa(a.Port)
}

func (a *NetAddress) Set(s string) error {
	hp := strings.Split(s, ":")
	if len(hp) != 2 {
		return errors.New("Need address in a form host:port")
	}
	port, err := strconv.Atoi(hp[1])
	if err != nil {
		return err
	}
	a.Host = hp[0]
	a.Port = port
	return nil
}
