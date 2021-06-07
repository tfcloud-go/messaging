# messaging

A go client for messaging service (nuntius).

## Usage

```go
func main() {
	options := messaging.Options{
		URL: "http://127.0.0.1:8080",
	}
	client, _ := messaging.NewClient(options)

	err := client.SendSMS([]string{"135****0000"}, "send from messaging client")
	if err != nil {
		panic(err)
	}

	to := []strings{"william****@****.com"}
	message := "send from messaging client"
	subject := "test messaging client"
	err = client.SendEmail(to, message, subject)
	if err != nil {
		panic(err)
	}
}
```

## Copyright

The TFCloud Go Team. [Apache License 2.0](./LICENSE)
