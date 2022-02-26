package objectstream

import (
	"fmt"
	"io"
	"net/http"
)

type PutStream interface {
	Write(p []byte) (n int, err error)
	Close() error
}

type putStream struct {
	writer *io.PipeWriter
	c      chan error
}

func NewPutStream(server, object string) PutStream {
	reader, writer := io.Pipe()
	c := make(chan error)

	go func() {
		url := fmt.Sprintf("http://%s/objects/%s", server, object)

		req, _ := http.NewRequest(http.MethodPut, url, reader)
		cli := http.Client{}
		r, e := cli.Do(req)
		if e == nil && r.StatusCode != http.StatusOK {
			e = fmt.Errorf("data server return http code %d", r.StatusCode)
		}
		c <- e
	}()

	return &putStream{writer: writer, c: c}
}

func (ps *putStream) Write(p []byte) (n int, err error) {
	return ps.writer.Write(p)
}

func (ps *putStream) Close() error {
	ps.writer.Close()
	return <-ps.c // 为了让reader 读到io.EOF
}
