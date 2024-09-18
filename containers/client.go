package containers

type ContainerCreateArgs struct {
	Image     string
	Name      string
	Port      int
	IPAddress string
	Env       []string
}

type Client interface {
	Create(args ContainerCreateArgs) (string, error)
	Start(id string) error
	Stop(id string) error
}
