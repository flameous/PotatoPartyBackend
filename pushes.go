package potato
//
//import (
//	"net/http"
//	"io/ioutil"
//	"github.com/pkg/errors"
//	"bytes"
//	"encoding/json"
//	firebase "firebase.google.com/go"
//
//	"firebase.google.com/go/messaging"
//	"context"
//	"log"
//)
//
//func init() {
//	log.SetFlags(log.LstdFlags | log.Lshortfile)
//}
//
//type PushService struct {
//	ServiceToken string
//	c            *http.Client
//}
//
//type MessageConfig struct {
//	Message Message `json:"message"`
//}
//
//type Message struct {
//	Token        string       `json:"token"`
//	Notification Notification `json:"notification"`
//}
//
//type Notification struct {
//	Title string `json:"title"`
//	Body  string `json:"body"`
//}
//
//func (p *PushService) sendPush(deviceToken, title, data string) error {
//	app, err := firebase.NewApp(context.Background(), &firebase.Config{
//		ProjectID: "potatoparty-199712",
//
//	})
//	if err != nil {
//		log.Fatal(err)
//	}
//	m, err := app.Messaging(context.Background())
//	log.Fatal(err)
//
//	_, err = m.Send(context.Background(), &messaging.Message{
//		Token: deviceToken,
//	})
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	if p.c == nil {
//		p.c = &http.Client{}
//	}
//
//	raw, err := json.Marshal(MessageConfig{Message{deviceToken, Notification{title, data}}})
//	if err != nil {
//		return errors.Wrap(err, "failed to unmarshal data")
//	}
//	buff := bytes.NewBuffer(raw)
//	//if err != nil {
//	//	return errors.Wrap(err, "read title and data struct to buffer")
//	//}
//	req, err := http.NewRequest("POST", "https://fcm.googleapis.com/v1/projects/potatoparty-199712/messages:send", buff)
//	if err != nil {
//		return errors.Wrap(err, "invalid request")
//	}
//
//	//req.Header.Set("Content-Type", "application/json")
//	//req.Header.Set("Authorization", "Bearer AIzaSyCDOiDEmfaGpn5hPP43oaDRGnWFbfs3G0M")
//
//	resp, err := p.c.Do(req)
//	if err != nil {
//		return errors.Wrap(err, "do request")
//	}
//	defer resp.Body.Close()
//
//	b, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		return errors.Wrap(err, "read resp body")
//	}
//
//	if b == nil {
//		return nil
//	}
//	return errors.New(string(b))
//}
