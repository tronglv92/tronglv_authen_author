package registry

import "github/tronglv_authen_author/internal/config"

type ServiceContext struct {
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{}
}
