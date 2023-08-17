package container

type ContainerInterface interface {
	BuildUseCase(code string) (interface{}, error)
	Get(code string) (interface{}, bool)
	Put(code string, value interface{})
}