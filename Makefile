check_swagger_install:
	which swagger || go get -u github.com/go-swagger/go-swagger/cmd/swagger

swagger: check_swagger_install
	swagger generate spec -o ./docs/swagger.yaml --scan-models

update_go_deps: 
	bazel run //:gazelle -- update-repos -from_file=go.mod
