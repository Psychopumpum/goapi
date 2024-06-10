package GoApi

import (
  "io"
  "io/ioutil"
  "net/http"
  "net/url"
)

const (
    GET    = "GET"
    POST   = "POST"
)

func (self *GoApi) urlEncode(host, path string, data map[string]string) string {
    baseurl, _  := url.Parse(host)
    baseurl.Path = path
    params := url.Values{}
    for k, v := range data {
        params.Add(k, v)
    }
    baseurl.RawQuery = params.Encode()
    return baseurl.String()
}

type Server struct {
    url     string
    headers map[string]string
    body    io.Reader
    client  *http.Client
    req     *http.Request
}

type Options struct {
    Url            string
    Headers        map[string]string
    Body           io.Reader
    FollowRedirect bool
}

type ResponseReader struct {
    Response *http.Response
}

func NewServer(opt  Options) *Server {
    var client *http.Client
    if opt.FollowRedirect {
        client = &http.Client{
            Transport: &http.Transport{},
            CheckRedirect: func(req *http.Request, via []*http.Request) error {
                return http.ErrUseLastResponse
            },
        }
    } else {
        client = &http.Client{
            Transport: &http.Transport{},
        }
    }
    server := &Server{
        url: opt.Url,
        headers: opt.Headers,
        body: opt.Body,
        client: client,
    }
    return server
}

func (self *Server) setupHeaders() {
    for k, v := range self.headers {
        self.req.Header.Add(k, v)
    }
}

func (self *Server) Get() (*Server) {
    self.req, _ = http.NewRequest(GET, self.url, self.body)
    self.setupHeaders()
    return self
}

func (self *Server) Post() (*Server) {
    self.req, _ = http.NewRequest(POST, self.url, self.body)
    self.setupHeaders()
    return self
}

func (self *Server) Exec() (r *ResponseReader, e error) {
    o, e := self.client.Do(self.req)
    r = &ResponseReader{o}
    return
}

func (self *ResponseReader) Content() (body []byte, err error) {
    defer self.Response.Body.Close()
    body, err = ioutil.ReadAll(self.Response.Body)
    return
}

func (self *ResponseReader) Text() (body string, err error) {
    t, err := self.Content()
    body = string(t)
    return
}