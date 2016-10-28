pushhook
====

`pushhook` is a Go package for push message event hook.

## Usage

```golang
func main() {
	service := "PushBullet"
	token := "YOUR_PUSHBULLET_API_TOKEN"

	push := pushhook.NewPushHook(service, token)
	push.Subscribe(received)
}

func received(msg string) {
	log.Println(msg)
}
```