package jpush

import (
	"bytes"
	"encoding/base64"
	"io/ioutil"
	"net/http"
)

const (
	HOST_NAME_SSL = "https://device.jpush.cn/v3/devices/191e35f7e0f88d4f701"
	BASE64_TABLE  = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
)
const (
	CHARSET                    = "UTF-8"
	CONTENT_TYPE_JSON          = "application/json"
	DEFAULT_CONNECTION_TIMEOUT = 20 //seconds
	DEFAULT_SOCKET_TIMEOUT     = 30 // seconds
)

var base64Coder = base64.NewEncoding(BASE64_TABLE)

type AlisaName struct {
	Alias string `json:"alias"`
}

type AlisaClient struct {
	MasterSecret string
	AppKey       string
	AuthCode     string
	BaseUrl      string
}

func NewAlisaClient(secret, appKey string) *AlisaClient {
	//base64
	auth := "Basic " + base64Coder.EncodeToString([]byte(appKey+":"+secret))
	alisaer := &AlisaClient{secret, appKey, auth, HOST_NAME_SSL}
	return alisaer
}
func (a *AlisaClient) SetAlisa(data []byte) (string, error) {
	ret, err := SendPostBytes(a.BaseUrl, data, a.AuthCode)
	if err != nil {
		return ret, err
	}

	return ret, err

}
func SendPostBytes(url string, data []byte, authCode string) (string, error) {

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	req.Header.Add("Charset", CHARSET)
	req.Header.Add("Authorization", authCode)
	req.Header.Add("Content-Type", CONTENT_TYPE_JSON)
	resp, err := client.Do(req)

	if err != nil {
		if resp != nil {
			resp.Body.Close()
		}
		return "", err
	}
	if resp == nil {
		return "", nil
	}

	defer resp.Body.Close()
	r, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(r), nil
}
