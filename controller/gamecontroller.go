package controller

import (
	"game21/logger"
	"game21/model"
	"github.com/gorilla/websocket"
	"math/rand"
	"net/http"
)

type GameController struct {
}

var logr logger.Logger

var upgrader = websocket.Upgrader{
	CheckOrigin:       func(r *http.Request) bool {
		return true
	},
}

func (controller *GameController) GameHandler() http.HandlerFunc {

	return func(writer http.ResponseWriter, request *http.Request) {
		println("hello")

		ws, err := upgrader.Upgrade(writer, request, nil)
		if err != nil {
			logr.LogErr(err)
			ws.Close()
			return
		}

		msg := model.GameModelResponse{}
		if err = ws.ReadJSON(msg); err != nil {
			logr.LogErr(err)
			ws.Close()
			return
		}

		if msg.Type != 0 { return }

		go handleGameLogic(ws)
	}
}

func handleGameLogic(ws *websocket.Conn) {
	min := 6
	max := 11
	for {
		msg := model.GameModelResponse{}
		if err := ws.ReadJSON(msg); err != nil {
			logr.LogErr(err)
			ws.Close()
			return
		}

		if msg.Type != 1 {
			errorModel := struct {
				Error string `json:"error"`
			}{}
			errorModel.Error = "Wrong command!"
			errorResponse := model.GameModelResponse{
				Type:  3,
				Model: errorModel,
			}
			ws.WriteJSON(errorResponse)
			return
		}

		//player card
		cardValue := rand.Intn(max - min) + min
		msgResponse := model.GameModelResponse {
			Type:  1,
			Model: model.ScoreModel{Score: cardValue},
		}

		if err := ws.WriteJSON(msgResponse); err != nil {
			logr.LogErr(err)
			ws.Close()
			return
		}

		//server card
		cardValue = rand.Intn(max - min) + min
		msgResponse.Type = 2
		msgResponse.Model = model.ScoreModel{Score: cardValue}
		if err := ws.WriteJSON(msgResponse); err != nil {
			logr.LogErr(err)
			ws.Close()
			return
		}
	}
}