package usecases

import "github.com/unawaretub86/leal-challenge/src/domain/ports"

type LealUseCase struct {
	rest ports.LealPort
}

func NewLealUseCase(rest ports.RepositoryPort) *LealUseCase {
	return &LealUseCase{
		rest,
	}
}
