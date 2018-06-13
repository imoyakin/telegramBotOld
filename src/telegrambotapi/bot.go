package telegrambotapi

import (
	"fmt"
	"reflect"
	//"fmt"
	"crypto/tls"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

//APIadress = https://api.telegram.org/
//Only need set your token and proxy
//if you don't need proxy,don't set it
type Bot struct {
	APIadress string
	Token     string
	Proxy     string
}

func (c *Bot) getDate(method string) (json.RawMessage, error) {
	proxy, err := url.Parse(c.Proxy)
	if err != nil {
		fmt.Println("proxy is null")
		panic(err)
	}
	var tr *http.Transport
	if c.Proxy == "" {
		tr = &http.Transport{
			Proxy: http.ProxyURL(proxy),
			//disabled HTTP/2
			TLSNextProto: make(map[string]func(authority string, c *tls.Conn) http.RoundTripper),
		}
	}
	tr = &http.Transport{
		//Proxy: http.ProxyURL(proxy),
		//disabled HTTP/2
		TLSNextProto: make(map[string]func(authority string, c *tls.Conn) http.RoundTripper),
	}

	client := &http.Client{Transport: tr}
	resp, err := client.Get(c.APIadress + c.Token + method)
	if err != nil {
		panic(err)
	}
	var APIresp APIResponse
	_, err = decodeAPIResponse(resp.Body, &APIresp)
	if err != nil {
		panic(err)
	}
	return APIresp.Result, err
}

func (c *Bot) sendDate(method string, responseBody url.Values) (json.RawMessage, error) {
	proxy, err := url.Parse(c.Proxy)
	if err != nil {
		panic(err)
	}
	tr := &http.Transport{
		Proxy:        http.ProxyURL(proxy),
		TLSNextProto: make(map[string]func(authority string, c *tls.Conn) http.RoundTripper),
	}
	client := &http.Client{Transport: tr}
	resp, err := client.PostForm(c.APIadress+c.Token+method, responseBody)
	if err != nil {
		panic(err)
	}
	var APIresp APIResponse
	_, err = decodeAPIResponse(resp.Body, &APIresp)
	if err != nil {
		panic(err)
	}
	fmt.Println(APIresp)
	return APIresp.Result, err
}

func (c *Bot) GetMe() (User, error) {
	resp, err := c.getDate("getMe")
	if err != nil {
		return User{}, err
	}
	var user User
	json.Unmarshal(resp, &user)
	return user, err
}

func (c *Bot) GetUpdates() ([]Update, error) {
	resp, err := c.getDate("getUpdates")
	if err != nil {
		return []Update{}, err
	}
	var update []Update
	err = json.Unmarshal(resp, &update)
	return update, err
}

func (c *Bot) SendMessage(ChatID string, text string) (Message, error) {
	var sendMsg SendMessage
	sendMsg.ChatID = ChatID
	sendMsg.Text = text

	respbody := structToMap(&sendMsg)
	fmt.Println(respbody)
	resp, err := c.sendDate("sendMessage", respbody)
	if err != nil {
		panic(err)
	}
	var replyMsg Message
	json.Unmarshal(resp, &replyMsg)
	return replyMsg, err
}

func decodeAPIResponse(responseBody io.Reader, resp *APIResponse) (_ []byte, err error) {
	data, err := ioutil.ReadAll(responseBody)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, resp)
	if err != nil {
		return
	}
	return data, nil
}

func (c *Bot) SetWebHook(url string,  cert []byte) {
	var set SetWebHook
	set.Url = url
	//set.Certificate = 
}

func (c *Bot) DeleteWebhook() (json.RawMessage, error) {
	var s url.Values
	result, err := c.sendDate("deleteWebhook", s)
	return result, err
}

// func (c*Bot) ListenMode() {
// 	var msg Message

// }

// func (c *Bot)GetWebhookInfo()  {

// }

func structToMap(i interface{}) (values url.Values) {
	values = url.Values{}
	iVal := reflect.ValueOf(i).Elem()
	typ := iVal.Type()
	for i := 0; i < iVal.NumField(); i++ {
		f := iVal.Field(i)
		tag := typ.Field(i).Tag.Get("json")
		var v string
		switch f.Interface().(type) {
		case int, int8, int16, int32, int64:
			v = strconv.FormatInt(f.Int(), 10)
		case uint, uint8, uint16, uint32, uint64:
			v = strconv.FormatUint(f.Uint(), 10)
		case float32:
			v = strconv.FormatFloat(f.Float(), 'f', 4, 32)
		case float64:
			v = strconv.FormatFloat(f.Float(), 'f', 4, 64)
		case []byte:
			v = string(f.Bytes())
		case string:
			v = f.String()
		}
		values.Set(tag, v)
	}
	return values
}
