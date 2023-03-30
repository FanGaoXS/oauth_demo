package gitee

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
	ClientId     = environment.Env.GiteeClientID
	ClientSecret = environment.Env.GiteeSecretID
)

type UserInfo struct {
	HtmlUrl   string    `json:"html_url"`
	AvatarUrl string    `json:"avatar_url"`
	Name      string    `json:"name"`
	Login     string    `json:"login"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

// Authorize
// 向服务提供商gitee发起获得授权码的申请
func Authorize(c *gin.Context) {
	u, _ := url.Parse("https://gitee.com/oauth/authorize")
	values := u.Query()
	values.Set("client_id", ClientId)
	values.Set("redirect_uri", "http://localhost:8090/gitee/oauth/callback")
	values.Set("response_type", "code")
	u.RawQuery = values.Encode()

	// redirect to -> https://gitee.com/oauth/authorize?client_id={CLIENT_ID}&redirect_uri=http://localhost:8090/gitee/oauth/callback&response_type=code
	c.Redirect(http.StatusMovedPermanently, u.String())
	return
}

// Callback
// 用户同意授权后
func Callback(c *gin.Context) {
	code := c.Query("code")
	u, _ := url.Parse("https://gitee.com/oauth/token")
	values := u.Query()
	values.Set("client_id", ClientId)
	values.Set("client_secret", ClientSecret)
	values.Set("code", code)
	values.Set("grant_type", "authorization_code")
	values.Set("redirect_uri", "http://localhost:8090/gitee/oauth/callback")
	u.RawQuery = values.Encode()

	client := http.DefaultClient
	req, _ := http.NewRequest(http.MethodPost, u.String(), nil)
	req.Header.Set("accept", "application/json")
	res, _ := client.Do(req)
	// HTTP POST -> https://gitee.com/oauth/token?client_id={CLIENT_ID}&client_secret={CLIENT_SECRET}&code={CODE}&grant_type=authorization_code&redirect_uri=http://localhost:8090/gitee/oauth/callback
	defer res.Body.Close()
	bytes, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(bytes))

	obj := make(map[string]interface{})
	json.Unmarshal(bytes, &obj)
	c.JSON(http.StatusOK, obj)
	return
}

func RefreshToken(c *gin.Context) {
	refreshToken := c.Query("refresh_token")
	u, _ := url.Parse("https://gitee.com/oauth/token")
	values := u.Query()
	values.Set("grant_type", "refresh_token")
	values.Set("refresh_token", refreshToken)
	u.RawQuery = values.Encode()

	client := http.DefaultClient
	req, _ := http.NewRequest(http.MethodPost, u.String(), nil)
	res, _ := client.Do(req)
	// HTTP POST -> https://gitee.com/oauth/token?grant_type=refresh_token&refresh_token={REFRESH_TOKEN}
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

	u, _ := url.Parse("https://gitee.com/api/v5/user")
	values := u.Query()
	values.Set("access_token", token)
	u.RawQuery = values.Encode()

	client := http.DefaultClient
	req, _ := http.NewRequest(http.MethodGet, u.String(), nil)
	req.Header.Set("accept", "application/json")
	res, _ := client.Do(req)
	// HTTP GET -> https://gitee.com/api/v5/user?access_token={TOKEN}
	defer res.Body.Close()
	bytes, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(bytes))

	info := new(UserInfo)
	json.Unmarshal(bytes, &info)
	c.JSON(http.StatusOK, info)
	return
}
