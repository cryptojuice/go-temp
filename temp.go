package temp

import (
	"math/rand"
	"os"
	// "os/signal"
	"strconv"
	"strings"
	// "syscall"
	// "path"
	"log"
	"time"
)

// var c = make(chan os.Signal, 1)

type File struct {
	*os.File
	dir string
}

var filePaths = make([]string, 0)

func (file *File) SetDir(d string) {
	file.dir = d
}

func (file *File) Create() (f *os.File, err error) {
	if file.dir == "" {
		file.SetDir("go-temp/")
	}

	if err = os.MkdirAll(file.dir, 0777); err != nil {
		panic(err)
	}

	p := []string{file.dir, "tmp-", generateRandomString()}
	t := strings.Join(p, "")

	f, err = os.Create(t)
	filePaths = append(filePaths, f.Name())

	if err != nil {
		panic(err)
	}

	return f, err

}

func (file *File) CleanUp() {
	if len(filePaths) < 0 {
		for _, value := range filePaths {
			log.Print("test111")
			if err := os.RemoveAll(value); err != nil {
				panic(err)
			}
		}
	}
}

func generateRandomString() (s string) {
	t := time.Now().Unix()
	rand.Seed(t)
	return strconv.FormatInt(rand.Int63(), 10)
}

// func handler() {
// 	signal.Notify(c, os.Interrupt)
// 	signal.Notify(c, syscall.SIGTERM)
// 	go func() {
// 		<-c
// 		os.Exit(1)
// 	}()
// }
