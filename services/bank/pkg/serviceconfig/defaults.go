package serviceconfig

import "github.com/spf13/viper"

const EventstoreDbUrl = "eventstore_url"

func ConfigureDefault() {
	viper.SetDefault(EventstoreDbUrl, "http://localhost:2113?tls=false")
}
