package Utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func DingContext(userids []string, phoneNum []string, contentbody string, atALL bool, keyWord string, WebHook string) {
	type At struct {
		AtMobiles []string `json:"atMobiles"`
		AtUserIds []string `json:"atUserIds"`
		IsAtAll   bool     `json:"isAtAll"`
	}
	type Text struct {
		Content string `json:"content"`
	}
	type Context struct {
		At      *At    `json:"at"`
		Text    *Text  `json:"text"`
		MsgType string `json:"msgtype"`
	}
	keywords := keyWord
	s1 := userids
	s2 := phoneNum
	s3 := keywords + contentbody
	s4 := "text"
	dingContext := Context{
		MsgType: s4,
	}

	DingAt := new(At)
	DingAt.AtUserIds = s1
	DingAt.AtMobiles = s2
	DingAt.IsAtAll = atALL
	dingContext.At = DingAt

	DingText := new(Text)
	DingText.Content = s3
	dingContext.Text = DingText

	jsonDing, err := json.Marshal(dingContext)
	if err != nil {
		fmt.Println(err.Error())
	}
	webHook := WebHook
	resp, err := http.Post(webHook,
		"application/json",
		bytes.NewReader(jsonDing))
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(string(body))
	}

}
