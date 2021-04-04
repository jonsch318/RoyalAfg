package serviceconfig

import "github.com/spf13/viper"

const ElasticSearchAddress = "elasticsearch_address"
const ElasticSearchUsername = "elasticsearch_username"
const ElasticSearchPassword = "elasticsearch_password"

func ConfigureDefaults() {
	viper.SetDefault(ElasticSearchUsername, "")
	viper.SetDefault(ElasticSearchPassword, "")
}
