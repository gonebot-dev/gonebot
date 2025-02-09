# Plugin

A gonebot plugin contains info and functions.  

```go
type GonePlugin struct{
  Name string
  Description string
  Version string

  Handler []GoneHandler
  ActiveHandler []GoneActiveHandler
}
```

## Handler

`Handler` is a passive message handler. When the bot received message, the Handler will be triggered to process the message.

```go
type GoneHandler struct{
  Matcher func(msg Message) bool
  Handler func(incomingMsg Message,resultMsg *Message) bool
}
```

### Matcher

A function who checks the message. If condition matched, it returns true, and then gonebot calls the Handler.  

The Handler can add MessageSegment to resultMsg, modify result Msg or do whatever you want.  
If Handler return true, gonebot will block following plugins and immediately send the result.

## ActiveHandler

ActiveHandler will continuing running in a goroutine. You can listen in some online service, for example RSS.  
After returning, the function restarts immediatly.  
You need to construct the result message yourself.

BE CAREFUL don't eat up all the cpu time here. Sleep when necessary!

```go
type GoneActiveHandler struct{
  Handler func() Message
}
```
