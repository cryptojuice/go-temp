package temp

import "os"

type File struct {
}

func Create() (f *os.File, err error) {
	return os.Create("temp123")
}
