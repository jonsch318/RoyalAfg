check_swagger_install:
	which swagger || go get -u github.com/go-swagger/go-swagger/cmd/swagger

swagger: check_swagger_install
	swagger generate spec -o ./services/docs/swagger.yaml --scan-models

update_go_deps: 
	bazel run //:gazelle -- update-repos -from_file=go.mod -to_macro=deps.bzl%go_dependencies

update:
	bazel run //:gazelle update

proto:
	protoc --go_out=plugins=. --go_opt=paths=source_relative ./pkg/user/pkg/user/protos/user.proto   

protos:
	cd ./pkg/protos && make protos

push_user:
	bazel run //pkg/user:push_image --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64

push_auth:
	bazel run //pkg/auth:push_image --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64

docker_build_search:
	docker build -t ryl_search --build-arg service=./services/search/main.go .

docker_build_user:
	docker build -t ryl_user --build-arg service=./services/user/main.go .