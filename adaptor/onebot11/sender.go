package onebot11

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/gonebot-dev/gonebot/messages"

	"github.com/gorilla/websocket"
)

// The third thread, Sender.
func SendingMessage(ws *websocket.Conn) {
	log.Printf("Sender started.\n")
	for { //get a result
		result := messages.PopResult()
		var payload APIPayload
		//private message
		if result.MessageType == "private" {
			userid, _ := strconv.Atoi(result.To)
			params := APISendPrivateMessage{}
			params.UserID = userid
			//text part
			if result.Text != "" {
				params.Message = append(params.Message,
					MessageSegmentText{
						Type: "text",
						Data: struct {
							Text string `json:"text"`
						}{Text: result.Text}})
			}
			//image part
			if len(result.Imgs) > 0 {
				for _, img := range result.Imgs {
					params.Message = append(params.Message,
						MessageSegmentImg{
							Type: "image",
							Data: struct {
								Uri string "json:\"file\""
							}{Uri: img},
						})
				}
			}

			payload.Action = "send_private_msg"
			payload.Params = params

		}
		//group message
		if result.MessageType == "group" {
			groupID, _ := strconv.Atoi(result.To)
			params := APISendGroupMessage{}
			params.GroupID = groupID
			//text part
			if result.Text != "" {
				params.Message = append(params.Message,
					MessageSegmentText{
						Type: "text",
						Data: struct {
							Text string "json:\"text\""
						}{Text: result.Text}})
			}
			//image part
			if len(result.Imgs) > 0 {
				for _, img := range result.Imgs {
					params.Message = append(params.Message,
						MessageSegmentImg{
							Type: "image",
							Data: struct {
								Uri string "json:\"file\""
							}{Uri: img},
						})
				}
			}
			payload.Action = "send_group_msg"
			payload.Params = params

		}

		//send
		jsonResult, _ := json.Marshal(payload)
		err := ws.WriteMessage(websocket.TextMessage, jsonResult)
		if err != nil {
			log.Println(err)
		}

		log.Printf("Sending message: %s\n", jsonResult)
	}
}
