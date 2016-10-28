package pushhook

import (
	"github.com/sharkattack51/pushbullet-go"
	"strings"
)

type PushHook struct {
	Service string
	Token   string
}

type RecievedCallback func(msg string)

func NewPushHook(service string, token string) *PushHook {
	return &PushHook{service, token}
}

func (p *PushHook) Subscribe(cb RecievedCallback) error {
	errCh := make(chan error)
	if strings.ToLower(p.Service) == "pushbullet" {
		go p.pushbullet(cb, errCh)
		err := <-errCh
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *PushHook) pushbullet(cb RecievedCallback, errCh chan error) {
	notify := make(chan bool)
	go pushbullet.SubscribeStream(p.Token, notify)
	<-notify

	pb := pushbullet.New(p.Token)
	pushes, err := pb.GetListPushes()
	if err != nil {
		errCh <- err
	}

	msg := ""
	if len(pushes) > 0 {
		msg = pushes[0].Body
	}

	if cb != nil {
		cb(msg)
	}

	errCh <- nil
}
