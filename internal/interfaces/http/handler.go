package http

import "heating-prediction/internal/domain/ports/service"

type Handlers struct {
	Service service.ComponentService
	Auth    service.AuthService
}
