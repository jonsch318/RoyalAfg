package handlers

import (
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/models"
	"log"

	"github.com/gorilla/websocket"
)

type PlayerConn struct {
	conn    *websocket.Conn
	Out     chan []byte
	In      chan *models.Event
	OnClose func(error)
	Close   chan bool
}

func NewPlayerConn(conn *websocket.Conn) *PlayerConn {
	return &PlayerConn{
		conn:  conn,
		Out:   make(chan []byte),
		In:    make(chan *models.Event),
		Close: make(chan bool),
	}
}

func (p *PlayerConn) reader() {
	defer func() {
		p.conn.Close()
	}()
	for {
		_, message, err := p.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseNormalClosure) {
				log.Printf("WS Message Error: %v", err)
				p.Close <- true
				break
			}
			p.Close <- false
			break
		}

		event, err := models.NewEventFromRaw(message)

		if err != nil {
			log.Printf("Error parsing message %v", err)
			continue
		}

		select {
		case p.In <- event:
		default:
		}

	}
}

func (p *PlayerConn) writer() {
	for {
		select {
		case message, ok := <-p.Out:
			if !ok {
				log.Printf("Closing conn due to sending error")
				p.conn.WriteMessage(websocket.CloseMessage, make([]byte, 0))
				p.conn.Close()

				p.Close <- true
				return
			}

			w, err := p.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				p.conn.WriteMessage(websocket.CloseMessage, make([]byte, 0))
				p.conn.Close()
				p.Close <- false
				w.Close()
				return
			}

			w.Write(message)
			w.Close()
		}
	}
}
