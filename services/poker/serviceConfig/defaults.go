package serviceConfig

import (
	"time"

	"github.com/spf13/viper"
)

const Port = "port"
const PlayersRequiredForStart = "player_required_for_start"
const PlayerActionTimeout = "player_action_timeout"
const BuyInOptions = "buy_in_options"

//SetDefaults configures all defaults
func SetDefaults() {
	viper.SetDefault(Port, 5000)
	//viper.SetDefault(BuyInOptions, []int{500, 1000})
	viper.SetDefault(BuyInOptions, [][]int{{500, 1500, 25}, {1500, 5000, 100}, {5000, 15000, 250}, {15000, 50000, 1000}, {50000, 200000, 2500}})
	viper.SetDefault(PlayersRequiredForStart, 2)
	viper.SetDefault(PlayerActionTimeout, 1*time.Minute)
}
