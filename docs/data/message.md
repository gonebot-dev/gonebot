# Message

A message struct contains multiple message segments.

```go
type Message struct{
  // UID of sender.
  SenderID string
  // UID of the bot.
  SelfID string

  // Whether it is a group message.
  IsGroup bool
  // If the message contains @at_bot
  IsToMe bool
  // Group ID
  GroupID bool

  // Message segments
  Segments []MessageSegment
}
```

## Message segments

Message segments have some types: `text`,`image`,`video`,`sound`,`action`.

```go
type MessageSegments struct{
  // Type: text,image,video,sound,action.
  Type string
  // Content of the message.
  Content string
}
```

## Utils

`getText()` will return all the text content in string.  
`getImageSegs()` `getVideoSegs()` `getSoundSegs()` `getActionSegs()` will return all the corresponding segments in array.

`addTextSeg()` `addImageSeg()` `addVideoSeg()` `addSoundSeg()` `addActionSeg()` will add all the corresponding segments.
