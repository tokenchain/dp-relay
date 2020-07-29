package x

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type WriterCert struct {
	path     string
	filename string
}

func NewCertWriter(path, filename string) WriterCert {
	return WriterCert{
		path:     path,
		filename: filename,
	}
}
func (w WriterCert) getPath() string {
	return w.path
}
func (w WriterCert) getNameOnly() string {
	return w.filename
}
func (w WriterCert) getFileName() string {
	return fmt.Sprintf("%s/%s", w.path, w.filename)
}
func (w WriterCert) ClearFile() WriterCert {
	ioutil.WriteFile(w.getFileName(), []byte(""), os.ModePerm)
	return w
}

func (w WriterCert) NewFile(data []byte) {
	ioutil.WriteFile(w.getFileName(), data, os.ModePerm)
}

func (w WriterCert) AppendLine(info string) {
	f, err := os.OpenFile(w.getFileName(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	line := fmt.Sprintf("%s\n", info)
	if _, err := f.Write([]byte(line)); err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
