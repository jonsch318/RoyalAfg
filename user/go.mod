module github.com/JohnnyS318/RoyalAfgInGo/user

go 1.15

replace github.com/JohnnyS318/RoyalAfgInGo/shared => ../shared

require (
	github.com/JohnnyS318/RoyalAfgInGo/shared v0.0.0-20200812205550-fe1d79422453
	github.com/Kamva/mgm v1.2.3
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/golang/protobuf v1.4.1
	github.com/gorilla/mux v1.7.4
	github.com/justinas/alice v1.2.0
	github.com/mitchellh/go-homedir v1.1.0
	github.com/mitchellh/mapstructure v1.3.3 // indirect
	github.com/pelletier/go-toml v1.8.0 // indirect
	github.com/spf13/afero v1.3.4 // indirect
	github.com/spf13/cast v1.3.1 // indirect
	github.com/spf13/cobra v1.0.0
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.7.1
	go.mongodb.org/mongo-driver v1.4.0
	go.uber.org/zap v1.15.0
	golang.org/x/net v0.0.0-20200707034311-ab3426394381 // indirect
	golang.org/x/sys v0.0.0-20200812155832-6a926be9bd1d // indirect
	google.golang.org/protobuf v1.25.0
	gopkg.in/ini.v1 v1.57.0 // indirect
)
