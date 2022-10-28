package gamesrepo 

import (
	"encoding/json"

	"github.com/matiasvarela/errors"

	"hexagonal/internal/core/domain"
	"hexagonal/pkg/apperrors"
)

type memkvs struct {
	kvs map[string][]byte
}

func NewMemKVS() *memkvs {
	return &memkvs{kvs: map[string][]byte{}}
}

func (repo *memkvs) Get(id string) (domain.Game, error) {
	if value, ok := repo.kvs[id]; ok {
		game := domain.Game{}
		err := json.Unmarshal(value, &game)
		if err != nil {
			return domain.Game{}, errors.New(apperrors.Internal, err, "Failed to get value from kvs", "")
		}

		return game, nil
	}

	return domain.Game{}, errors.New(apperrors.Internal, nil, "Game not found from kvs", "")
}


func (repo *memkvs) Save(game domain.Game) error {
	bytes, err := json.Marshal(game)
	if err != nil {
		return errors.New(apperrors.InvalidInput, err, "Game fails at marshal into json string", "")
	}

	repo.kvs[game.ID] = bytes

	return nil
}