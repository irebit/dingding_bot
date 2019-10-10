package dingtalk_group_robot

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Robot struct {
	RequestUrl string
	Message    interface{}
}

type Resp struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

var instance *Robot

func New() *Robot {
	if instance == nil {
		instance = &Robot{}
	}
	return instance
}

//
func (r *Robot) SetAccessToken(accessToken string) *Robot {
	r.RequestUrl = fmt.Sprintf("https://oapi.dingtalk.com/robot/send?access_token=%s", accessToken)
	return r
}

// log.Println(string(body))

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
