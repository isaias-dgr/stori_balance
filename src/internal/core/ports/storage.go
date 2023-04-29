package ports

import (
	"bytes"
	"io"
)

type Storage interface {
	GetListFiles(address string) ([]string, error)
	GetFile(address string) (io.ReadCloser, error)
	PutFile(address string, doc *bytes.Buffer) (string, error)
}
