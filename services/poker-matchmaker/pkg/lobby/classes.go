package lobby

import (
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/poker/models"
)

func (m *Manager) GetRegisteredClasses() []models.Class {
	return m.classes
}
