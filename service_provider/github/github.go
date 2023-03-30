package github

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	"oauth-demo/environment"
)

var (
	ClientId     = environment.Env.GithubClientID
	ClientSecret = environment.Env.GithubSecretID
)

type UserInfo struct {
	HtmlUrl   string    `json:"html_url"`
	AvatarUrl string    `json:"avatar_url"`
	Name      string    `json:"name"`
	Login     string    `json:"login"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

func Authorize(c *gin.Context) {
	u, _ := url.Parse("https://github.com/login/oauth/authorize")
	values := u.Query()
	values.Set("client_id", ClientId)
	values.Set("redirect_uri", "http://localhost:8090/github/oauth/callback")
	u.RawQuery = values.Encode()

	c.Redirect(http.StatusMovedPermanently, u.String())
	return
}

func Callback(c *gin.Context) {
	code := c.Query("code")
	u, _ := url.Parse("https://github.com/login/oauth/access_token")
	values := u.Query()
	values.Set("client_id", ClientId)
	values.Set("client_secret", ClientSecret)
	values.Set("code", code)
	u.RawQuery = values.Encode()

	client := http.DefaultClient
	req, _ := http.NewRequest(http.MethodPost, u.String(), nil)
	req.Header.Set("accept", "application/json")
	res, _ := client.Do(req)
	defer res.Body.Close()
	bytes, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(bytes))

	obj := make(map[string]interface{})
	json.Unmarshal(bytes, &obj)
	c.JSON(http.StatusOK, obj)
	return
}

func Userinfo(c *gin.Context) {
	token := c.Query("token")
	u, _ := url.Parse("https://api.github.com/user")

	client := http.DefaultClient
	req, _ := http.NewRequest(http.MethodGet, u.String(), nil)
	req.Header.Set("accept", "application/json")
	req.Header.Set("Authorization", "token "+token)
	res, _ := client.Do(req)
	defer res.Body.Close()
	bytes, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(bytes))

	info := new(UserInfo)
	json.Unmarshal(bytes, &info)
	c.JSON(http.StatusOK, info)
	return
}
