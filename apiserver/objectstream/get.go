package objectstream

import (
	"fmt"
	"io"
	"net/http"
)

type GetStream interface {
	Read(p []byte) (n int, err error)
}

type getStream struct {
	reader io.Reader
}

func NewGetStream(server, object string) (GetStream, error) {
	if server == "" || object == "" {
		return nil, fmt.Errorf("invaild server %s object %s", server, object)
	}

	url := fmt.Sprintf("http://%s/objects/%s", server, object)

	return newGetStream(url)
}

func newGetStream(url string) (GetStream, error) {
	r, e := http.Get(url)
	if e != nil {
		return nil, e
	}

	if e == nil && r.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("data server return http code %d", r.StatusCode)
	}

	return &getStream{reader: r.Body}, nil
}

func (g *getStream) Read(p []byte) (n int, err error) {
	return g.reader.Read(p)
}
