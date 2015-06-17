package temp

import (
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
)

func Create() (file *os.File, err error) {
	s := []string{"tmp-", generateRandomString()}
	f, err := os.Create(strings.Join(s, ""))
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		<-c
		err := os.Remove(f.Name())
		if err != nil {
			os.Exit(0)
		}
		os.Exit(1)
	}()
	return f, err
}

func generateRandomString() (s string) {
	return strconv.FormatInt(rand.Int63(), 10)
}
