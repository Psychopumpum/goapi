package goapi

import (
  "bytes"
  "encoding/json"
  "log"
)

type PsychopumpumApi struct {
    BaseURL string
    Apikey  string
}

func NewPsychopumpumApi(apikey string) *PsychopumpumApi {
    client := &PsychopumpumApi{
        BaseURL: "https://api.psychopumpum.fun/",
        Apikey : apikey,
    }
    return client
}

func (self *PsychopumpumApi) Get(path string, params map[string]string) map[string]interface{} {
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

// ChatBot
func (self *PsychopumpumApi) OpenAi(text string) map[string]interface{} {
    return self.Get("openai/", map[string]string{
        "q": text,
    })
}

// Instagram
func (self *PsychopumpumApi) InstagramProfile(username string) map[string]interface{} {
    return self.Get("instagram/post/", map[string]string{
        "username": username,
    })
}

func (self *PsychopumpumApi) InstagramPost(url string) map[string]interface{} {
    return self.Get("instagram/post/", map[string]string{
        "url": url,
    })
}

func (self *PsychopumpumApi) InstagramStory(url string) map[string]interface{} {
    return self.Get("instagram/story/", map[string]string{
        "url": url,
    })
}