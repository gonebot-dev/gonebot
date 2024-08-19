package onebot11

import (
	"encoding/json"
	"gonebot/messages"
	"log"
	"strconv"

	"github.com/gorilla/websocket"
)

// The third thread, Sender.
func SendingMessage(ws *websocket.Conn) {
	log.Printf("Sender started.\n")
	for { //get a result
		result, succ := messages.PopResult()
		if !succ {
			continue
		}
		var payload APIPayload
		//send_private_message
		if result.MessageType == "private" {
			userid, _ := strconv.Atoi(result.To)
			params := APISendPrivateMessage{}
			params.UserID = userid
			params.Message = append(params.Message,
				MessageSegmentText{
					Type: "text",
					Data: struct {
						Text string `json:"text"`
					}{Text: result.Text}})

			payload.Action = "send_private_msg"
			payload.Params = params

		}
		jsonResult, _ := json.Marshal(payload)
		err := ws.WriteMessage(websocket.TextMessage, jsonResult)
		if err != nil {
			log.Println(err)
		}

		log.Printf("Sending message: %s\n", jsonResult)
	}
}
