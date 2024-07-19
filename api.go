package goapi

import (
  "bytes"
  "encoding/json"
  "log"
  "strings"
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

// Facebook Downloader
func (self *PsychopumpumApi) FacebookDownload(url string) map[string]interface{} {
    return self.Get("facebook/download/", map[string]string{
        "url": url,
    })
}

// Instagram
func (self *PsychopumpumApi) InstagramProfile(username string) map[string]interface{} {
    return self.Get("instagram/profile/", map[string]string{
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

// TikTok
func (self *PsychopumpumApi) TikTokProfile(username string) map[string]interface{} {
    return self.Get("tiktok/profile/", map[string]string{
        "username": username,
    })
}

func (self *PsychopumpumApi) TikTokDownload(url string) map[string]interface{} {
    return self.Get("tiktok/download/", map[string]string{
        "url": url,
    })
}

// Pinterest 
func (self *PsychopumpumApi) PinterestDownload(url string) map[string]interface{} {
    return self.Get("pinterest/download/", map[string]string{
        "url": url,
    })
}

func (self *PsychopumpumApi) PinterestSimilar(url string) map[string]interface{} {
    return self.Get("pinterest/similar/", map[string]string{
        "url": url,
    })
}

func (self *PsychopumpumApi) PinterestSearch(q string) map[string]interface{} {
    return self.Get("pinterest/search/", map[string]string{
        "q": q,
    })
}

// YouTube
func (self *PsychopumpumApi) YouTubeDownload(url string) map[string]interface{} {
    return self.Get("youtube/download/", map[string]string{
        "url": url,
    })
}

func (self *PsychopumpumApi) YouTubeSearch(q, maxResult string) map[string]interface{} {
    return self.Get("youtube/search/", map[string]string{
        "q": q,
        "maxResult": maxResult, // default 5
    })
}

// RedTube
func (self *PsychopumpumApi) RedTubeSearch(q string) map[string]interface{} {
    return self.Get("redtube/search/", map[string]string{
        "q": q,
    })
}

func (self *PsychopumpumApi) RedTubeDownload(url string) map[string]interface{} {
    return self.Get("redtube/download/", map[string]string{
        "url": url,
    })
}

// PornHub
func (self *PsychopumpumApi) PornHubSearch(q, page string) map[string]interface{} {
    return self.Get("pornhub/search/", map[string]string{
        "q": q,
        "page": page, // default 1
    })
}

func (self *PsychopumpumApi) PornHubDownload(url string) map[string]interface{} {
    return self.Get("pornhub/download/", map[string]string{
        "url": url,
    })
}

// Twitter
func (self *PsychopumpumApi) TwitterProfile(username string) map[string]interface{} {
    return self.Get("twitter/profile/", map[string]string{
        "username": username,
    })
}

func (self *PsychopumpumApi) TwitterDownload(url string) map[string]interface{} {
    return self.Get("twitter/download/", map[string]string{
        "url": url,
    })
}

// LINE
func (self *PsychopumpumApi) LineVoom(url string) map[string]interface{} {
    return self.Get("line/voom/info/", map[string]string{
        "url": url,
    })
}

func (self *PsychopumpumApi) LoginWithQrCode(appType string) map[string]interface{} {
    return self.Get("line/login/qrcode", map[string]string{
        "appType": strings.ToUpper(appType), // IOSIPAD, ANDROIDSECONDARY, DESKTOPWIN, DESKTOPMAC, CHROMEOS
    })
}

func (self *PsychopumpumApi) LoginPinCode(session string) map[string]interface{} {
    return self.Get("line/login/qrcode", map[string]string{
        "session": session,
    })
}

func (self *PsychopumpumApi) LoginGetToken(session string) map[string]interface{} {
    return self.Get("line/login/token/", map[string]string{
        "session": session,
    })
}

func (self *PsychopumpumApi) LoginWithCredential(email, pwd, appType string) map[string]interface{} {
    return self.Get("line/login/credential/", map[string]string{
        "email": email,
        "password": pwd,
        "appType": strings.ToUpper(appType), // IOSIPAD, ANDROIDSECONDARY, DESKTOPWIN, DESKTOPMAC, CHROMEOS
    })
}

func (self *PsychopumpumApi) PrimaryToSecondary(authToken, appType string) map[string]interface{} {
    return self.Get("line/primary/secondary/", map[string]string{
        "authToken": authToken,
        "appType": strings.ToUpper(appType), // IOSIPAD, ANDROIDSECONDARY, DESKTOPWIN, DESKTOPMAC, CHROMEOS
    })
}

// Smule
func (self *PsychopumpumApi) SmuleProfile(username string) map[string]interface{} {
    return self.Get("smule/profile/", map[string]string{
        "username": username,
    })
}

func (self *PsychopumpumApi) SmuleDownload(url string) map[string]interface{} {
    return self.Get("smule/download/", map[string]string{
        "url": url,
    })
}

// BMKG
func (self *PsychopumpumApi) BMKG() map[string]interface{} {
    return self.Get("bmkg/info/", map[string]string{})
}

// Anime Search
func (self *PsychopumpumApi) AnimeSearch(query string) map[string]interface{} {
    return self.Get("anime/info/", map[string]string{
        "q": query,
    })
}

// Play Phrase
func (self *PsychopumpumApi) PhraseSearch(query string) map[string]interface{} {
    return self.Get("phrase/search/", map[string]string{
        "q": query,
    })
}

// Random Name
func (self *PsychopumpumApi) RandomName(country, count string) map[string]interface{} {
    return self.Get("random/name/", map[string]string{
        "country": country, //Indonesia
        "count": count, // 10
    })
}
