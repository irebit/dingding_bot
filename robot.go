package dingtalk_group_robot

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type Robot struct {
	RequestUrl string
	Message    interface{}
}

type Resp struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

const (
	WEBHOOK_URL = "https://oapi.dingtalk.com/robot/send?access_token=%s"
)

var instance *Robot

// New 获取一个 Robot instance
func New() *Robot {
	if instance == nil {
		instance = &Robot{}
	}
	return instance
}

// SetAccessToken 初始设置AccessToken
func (r *Robot) SetAccessToken(accessToken string) *Robot {
	r.RequestUrl = fmt.Sprintf(WEBHOOK_URL, accessToken)
	return r
}

// AddSign 更新SDK 支持Sign模式
func (r *Robot) AddSign(secret string) *Robot {
	microTimestamp := time.Now().UnixNano() / 1e6
	h := hmac.New(sha256.New, []byte(secret))
	io.WriteString(h, fmt.Sprintf("%d\n%s", microTimestamp, secret))
	if r.RequestUrl != "" {
		r.RequestUrl = r.RequestUrl + "&timestamp=" + strconv.Itoa(int(microTimestamp)) + "&sign=" + url.QueryEscape(base64.StdEncoding.EncodeToString(h.Sum(nil)))
	}
	return r
}

// Send 发送notification
func (r *Robot) Send(message interface{}) (bool, error) {
	b, err := json.Marshal(message)
	if err != nil {
		return false, err
	}
	// log.Println(string(b))
	resp, err := http.Post(r.RequestUrl, "application/json", bytes.NewReader(b))
	if err != nil {
		return false, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	ret := &Resp{}

	if err := json.Unmarshal(body, &ret); err != nil {
		return false, err
	}

	if ret.ErrCode == 0 {
		return true, nil
	} else {
		return false, errors.New(ret.ErrMsg)
	}
}
