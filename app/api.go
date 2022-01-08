package app

import (
	"bytes"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
)

type Api struct {
	Client *http.Client
	Url    string
	Token  string
}

func (a *Api) Connect() *http.Client {
	a.Client = &http.Client{}
	return a.Client
}

func (a *Api) SendNotification(message string) *http.Response {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("message", message)
	err := writer.Close()
	if err != nil {
		log.Println(err)
	}

	client := a.Connect()
	req, err := http.NewRequest("POST", a.Url, payload)
	if err != nil {
		log.Println(err)
	}

	req.Header.Add("Authorization", "Bearer "+a.Token)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()

	return res
}

func WriteRequest(req *http.Request) []byte {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println(err)
	}
	return body
}

func WriteReponse(res *http.Response) []byte {
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}
	return body
}
