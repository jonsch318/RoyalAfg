package serviceconfig

import (
	"github.com/spf13/viper"
)

const SwaggerUrl = "swagger_url"
const SwaggerTitle = "swagger_title"
const SwaggerFile = "swagger_file"

func ConfigureDefaults() {
	// HttpServer configuration
	viper.SetDefault(SwaggerTitle, "Royalafg Docs")
}
