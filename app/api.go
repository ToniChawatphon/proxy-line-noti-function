package app

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"strconv"
)

type Api struct {
	client *http.Client
	Url    string
	Token  string
}

func (a *Api) connect() *http.Client {
	a.client = &http.Client{}
	return a.client
}

// SendNotification receive request from Trading View
// and send it to Line Notify
func (a *Api) SendNotification(r *http.Request) {
	TrackingChannel = make(chan int)
	MessageChannel = make(chan string)

	go a.channelTracking()
	go a.parseRequest(r)
	a.forwardRequest()

	close(TrackingChannel)
	close(MessageChannel)
}

// forwardRequest create message payload
// and send it to specific url
func (a *Api) forwardRequest() {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("message", <-MessageChannel)
	err := writer.Close()
	if err != nil {
		log.Panicln(err)
	}

	client := a.connect()
	req, err := http.NewRequest("POST", a.Url, payload)
	if err != nil {
		log.Panicln(err)
	}

	req.Header.Add("Authorization", "Bearer "+a.Token)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		log.Panicln(err)
	}

	TrackingChannel <- res.StatusCode
	defer res.Body.Close()
}

func (a *Api) parseRequest(req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Panicln(err)
	}
	MessageChannel <- string(body)
}

// channelTracking tracks response metrices
func (a *Api) channelTracking() {
	var status string
	var statusCode int
	statusCode = <-TrackingChannel

	switch statusCode {
	case 200:
		status = fmt.Sprintf("OK %s", strconv.Itoa((statusCode)))
	case 400:
		status = fmt.Sprintf("Bad Request %s", strconv.Itoa((statusCode)))
	case 401:
		status = fmt.Sprintf("Authorized Error %s", strconv.Itoa((statusCode)))
	case 429:
		status = fmt.Sprintf("Too Many %s", strconv.Itoa((statusCode)))
	case 500:
		status = fmt.Sprintf("Channel Error %s", strconv.Itoa((statusCode)))
	default:
		status = fmt.Sprintf("Other Error %s", strconv.Itoa((statusCode)))
	}
	log.Println(status)
}
