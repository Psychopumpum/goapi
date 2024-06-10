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

func (self *GoApi) Get(path string, params map[string]string) map[string]interface{} {
    fullUrl := self.urlEncode(self.BaseURL, path, params)
    client := NewServer(Options{
        Url: fullUrl,
        Headers: map[string]string{
            "Authorization": self.Apikey,
        },
    })
    res, err := client.Get().Exec()
    if err != nil {
        log.Fatal(err)
    }

    var result map[string]interface{}
    body, _ := res.Content()
    json.NewDecoder(bytes.NewReader(body)).Decode(&result)
    return result
}

func (self *GoApi) InstagramProfile(username string) map[string]interface{} {
    return self.Get("instagram/post/", map[string]string{
        "username": username,
    })
}

func (self *GoApi) InstagramPost(url string) map[string]interface{} {
    return self.Get("instagram/post/", map[string]string{
        "url": url,
    })
}

func (self *GoApi) InstagramStory(url string) map[string]interface{} {
    return self.Get("instagram/story/", map[string]string{
        "url": url,
    })
}