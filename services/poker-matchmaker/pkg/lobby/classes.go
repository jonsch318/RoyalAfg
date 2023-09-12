package lobby

import (
	"github.com/jonsch318/royalafg/pkg/poker/models"
)

func (m *Manager) GetRegisteredClasses() []models.Class {
	return m.classes
}
