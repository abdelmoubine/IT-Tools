package engine

import (
	"io/ioutil"
	"os"
)

// WriteTempAndWipe writes a temp file and returns filename.
// Caller MUST call WipeFile to overwrite and remove later.
func WriteTempAndWipe(prefix string, content []byte) (string, error) {
	dir := "tmp"
	os.MkdirAll(dir, 0700)
	f, err := ioutil.TempFile(dir, prefix)
	if err != nil {
		return "", err
	}
	name := f.Name()
	_, err = f.Write(content)
	f.Sync()
	f.Close()
	return name, nil
}

func WipeFile(path string) {
	// best-effort overwrite and remove
	f, err := os.OpenFile(path, os.O_WRONLY, 0600)
	if err == nil {
		info, _ := f.Stat()
		size := info.Size()
		buf := make([]byte, size)
		f.Write(buf)
		f.Sync()
		f.Close()
	}
	os.Remove(path)
}