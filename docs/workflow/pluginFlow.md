# Plugin Flow

## Handler

After receiving message, gonebot will start a goroutine to process it. Gonebot creates an empty result message, runs plugin's matcher in order of plugin loading. If matcher returns true, handler function will be called. The goroutine wait for its result to see whether it is blocked and continue processing.  

The result message will be sent to cache channel.

## Active Handler

You can send message by yourself through the active handler.
