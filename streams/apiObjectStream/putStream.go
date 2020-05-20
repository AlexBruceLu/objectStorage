package apiObjectStream

import (
	"fmt"
	"io"
	"net/http"
)

type PutStream struct {
	writer *io.PipeWriter
	c      chan error
}

func NewPutStream(server, object string) *PutStream {
	reader, writer := io.Pipe()
	c := make(chan error)
	go func() {
		req, _ := http.NewRequest("PUT", "http://"+server+"/objects/"+object, reader)
		cli := http.Client{}
		r, e := cli.Do(req)
		if e == nil && r.StatusCode != http.StatusOK {
			e = fmt.Errorf("dataserver return http code: %d", r.StatusCode)
		}
		c <- e
	}()
	return &PutStream{writer: writer, c: c}
}

func (p *PutStream) Write(b []byte) (n int, e error) {
	return p.writer.Write(b)
}

func (p *PutStream) Close() error { // 统一返回状态码
	_ = p.writer.Close()
	return <-p.c
}
