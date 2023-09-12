package database

import "github.com/jonsch318/royalafg/pkg/models"

//############ GameBuffer ############

type GameBuffer struct {
	buffer       []*models.SlotGame
	bufferLength int
	Callback     func(gameBuffer []*models.SlotGame) error
}

func NewGameBuffer(callback func(gameBuffer []*models.SlotGame) error) *GameBuffer {
	return &GameBuffer{
		buffer:       make([]*models.SlotGame, 100),
		bufferLength: 0,
		Callback:     callback,
	}
}

func (b *GameBuffer) BufferGame(game *models.SlotGame) error {

	err := game.Validate()
	if err != nil {
		return err
	}

	b.buffer[b.bufferLength] = game
	b.bufferLength++

	if b.bufferLength != 0 && b.bufferLength%100 == 0 {
		err := b.Callback(b.buffer)
		if err != nil {
			b.bufferLength = 100

			// could not save buffer to db so we increase the buffered elements until we can save them.
			tmp := make([]*models.SlotGame, len(b.buffer)+100)
			copy(tmp, b.buffer)
			b.buffer = tmp
		}
		b.buffer = make([]*models.SlotGame, 100)
		b.bufferLength = 0
	}

	return nil
}

func (b *GameBuffer) GetBuffer() []*models.SlotGame {
	return b.buffer
}
