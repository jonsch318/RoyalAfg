package metrics

//UseCase contains all use cases of an app for metrics (HTTP, GRPC, CLI,...)
type UseCase interface {
	SaveHTTP(h *HTTP)
}
