package gateway

import "github.com.br/cristian.scherer/eda-balance/internal/entity"

type BallanceGateway interface {
	Save(ballance *entity.Balance) error
}
