# Adapter

Adapter connect to other platform to receive and send message.

Adapter's main function works in a goroutine. It communicate with gonebot through the cache channel.  

```go
type GoneAdapter struct{
  Name string
  Description string
  Version string
  SupportedPlatform string // "qq", "wechat", etc.

  Init func() // Run when starting
  Conn func(incomingChan chan message,resultChan chan message)
  Final func() // Run when closing
}
```
