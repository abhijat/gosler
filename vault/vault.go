package vault

import (
	"net/http"
	"fmt"
	"strings"
	"bytes"
	"io"
)

const TokenHeader = "X-Vault-Token"

func asBytes(r *http.Response) []byte {
	var b bytes.Buffer
	io.Copy(&b, r.Body)
	return b.Bytes()
}

type Client struct {
	Server string
	Token  string
}

func NewClient(url string, token string) *Client {
	if strings.HasSuffix(url, "/") {
		url = url[:len(url)-1]
	}

	return &Client{
		Server: url,
		Token:  token,
	}
}

func (c *Client) url(subpath string) string {
	if strings.HasPrefix(subpath, "/") {
		subpath = strings.Replace(subpath, "/", "", 1)
	}

	return fmt.Sprintf("%s/%s", c.Server, subpath)
}

func (c *Client) addToken(r *http.Request) *http.Request {
	r.Header.Set(TokenHeader, c.Token)
	return r
}

func (c *Client) do(r *http.Request) (*http.Response, error) {
	client := &http.Client{}
	return client.Do(c.addToken(r))
}

func (c *Client) HealthProbe() ([]byte, error) {
	request, err := http.NewRequest("GET", c.url("sys/health"), nil)
	if err != nil {
		return nil, err
	}

	response, err := c.do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	return asBytes(response), nil
}

func (c *Client) ReadSecret(path string) ([]byte, error) {
	path = fmt.Sprintf("secret/data/%s", path)
	request, err := http.NewRequest("GET", c.url(path), nil)
	if err != nil {
		return nil, err
	}

	response, err := c.do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%d", response.StatusCode)
	}

	return asBytes(response), nil
}
