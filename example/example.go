package main

import (
	"github.com/sharkattack51/pushhook"
	"log"
)

func main() {
	service := "PushBullet"
	token := "YOUR_PUSHBULLET_API_TOKEN"

	push := pushhook.NewPushHook(service, token)
	push.Subscribe(received)
}

func received(msg string) {
	log.Println(msg)
}
