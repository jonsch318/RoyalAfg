check_swagger_install:
	which swagger || go get -u github.com/go-swagger/go-swagger/cmd/swagger

swagger: check_swagger_install
	swagger generate spec -o ./swagger.yaml --scan-models
