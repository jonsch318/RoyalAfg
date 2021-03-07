package handlers

import (
	"sync"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/log"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/models"

	"github.com/gorilla/websocket"
)

type PlayerConn struct {
	Status byte
	Out     chan []byte
	In      chan *models.Event
	Close   chan bool
	conn    *websocket.Conn
	lock sync.Mutex
}

func NewPlayerConn(conn *websocket.Conn) *PlayerConn {
	return &PlayerConn{
		Status: 0b0,
		conn:  conn,
		Out:   make(chan []byte),
		In:    make(chan *models.Event),
		Close: make(chan bool),
	}
}

func (p *PlayerConn) reader() {
	defer log.Logger.Debugf("Conn reader closed")
	log.Logger.Debugf("Conn reader started")
	for {
		//we dont need message type. We detect closing in the error
		_, message, err := p.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseNormalClosure) {
				//Close was *unexpected* (abnormal)
				log.Logger.Errorw("WS Message Error", "error",  err)
				p.CloseConnection(true)
				break
			}
			//close was *expected* (normal) and messaging to close channel.
			p.CloseConnection(false)
			break
		}
		//Decode message
		event, err := models.NewEventFromRaw(message)
		if err != nil {
			log.Logger.Infow("Error parsing message", "error", err)
			//message could not be decoded and will not be passed through.
			continue
		}

		// Send to anyone listening on that channel. We use select here to defeat the blocking nature of the channel. If no one is listening the we dont pass through anything
		select {
		case p.In <- event:
		default:
		}

	}
}

func (p *PlayerConn) writer() {
	defer log.Logger.Debugf("Conn writer closed")
	log.Logger.Debugf("Conn writer started")
	for true {
		select {
		case message, ok := <-p.Out:
			if !ok {
				p.CloseConnection(false)
				return
			}

			//Create Websocket writer
			w, err := p.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				p.CloseConnection(true)
				return
			}

			//write message with the writer
			_, err = w.Write(message)
			if err != nil {
				p.CloseConnection(true)
				err = w.Close()
				if err != nil {
					log.Logger.Warnw("error during websocket writer closing", "error", err)
				}
				return
			}

			err = w.Close()
			if err != nil {
				log.Logger.Warnw("error during websocket writer closing", "error", err)
			}
		}
	}
}

func (p *PlayerConn) CloseConnection(unexpected bool){
	defer func() {
		if r := recover(); r != nil {
			log.Logger.Debugf("recovering in round start from %v", r)
		}
	}()

	log.Logger.Warnw("Closing Connection", "unexpected", unexpected)


	err := p.conn.WriteMessage(websocket.CloseMessage, make([]byte, 0))
	if err != nil {
		log.Logger.Warnw("Error during websocket closing message", "error", err)
	}

	err = p.conn.Close()
	if err != nil {
		log.Logger.Warnw("Error during websocket closing", "error", err)
	}

	p.Status = 0

	select {
	case <-p.Close:
	default:
		//Message close event
		log.Logger.Warnf("Closing channel")
		close(p.Close)
	}


}

