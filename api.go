package GoApi

import (
  "bytes"
  "encoding/json"
  "log"
)

type GoApi struct {
    BaseURL string
    Apikey  string
}

func NewPsychopumpumApi(apikey string) *GoApi {
    client := &GoApi{
        BaseURL: "https://api.psychopumpum.fun/",
        Apikey : apikey,
    }
    return client
}

func (self *GoApi) request(path string, params map[string]string) *Server {
    fullUrl := self.urlEncode(self.BaseURL, path, params)
    server := NewServer(Options{
        Url: fullUrl,
        Headers: map[string]string{
            "Authorization": self.Apikey,
        },
    })
    return server
}

func (self *GoApi) InstagramPost(url string) map[string]interface{} {
    cl := self.request("instagram/post/", map[string]string{
        "url": url,
    })
    res, err := cl.Get().Exec()
    if err != nil {
        log.Fatal(err)
    }
    var result map[string]interface{}
    body, err := res.Content()
    if err != nil {
        log.Fatal(err)
    }
    err = json.NewDecoder(bytes.NewReader(body)).Decode(&result)
    if err != nil {
        log.Fatal(err)
    }
    return result
}