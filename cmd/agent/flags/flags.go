package flags

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type NetAddress struct {
	Host string
	Port int
}

type ReportInterval struct {
	TimeInterval time.Duration
}

type PollInterval struct {
	TimeInterval time.Duration
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

func (r ReportInterval) String() string {
	return strconv.Itoa(int(r.TimeInterval))
}

func (r *ReportInterval) Set(s string) error {
	fmt.Println("s:", s)
	if s == "" || s == "0" {
		r.TimeInterval = 10 * time.Second // Default value
		return nil
	}
	interval, err := strconv.Atoi(s)
	if err != nil {
		return errors.New("You need to enter the кeport interval, an integer")
	}
	if interval < 0 {
		return errors.New("You need to enter the report interval, a positive integer")
	} else if interval == 0 {
		r.TimeInterval = 10 * time.Second // Default value
	} else {
		r.TimeInterval = time.Duration(interval) * time.Second
	}
	return nil
}

func (p PollInterval) String() string {
	return strconv.Itoa(int(p.TimeInterval))
}

func (p *PollInterval) Set(s string) error {
	if s == "" || s == "0" {
		p.TimeInterval = 10 * time.Second // Default value
		return nil
	}
	interval, err := strconv.Atoi(s)
	if err != nil {
		return errors.New("You need to enter the report interval, an integer")
	}

	if interval < 0 {
		return errors.New("You need to enter the report interval, a positive integer")
	} else if interval == 0 {
		p.TimeInterval = 2 * time.Second // Default value
	} else {
		p.TimeInterval = time.Duration(interval) * time.Second
	}

	return nil
}
