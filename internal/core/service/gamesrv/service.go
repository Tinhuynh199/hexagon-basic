package gamesrv

import (
	"github.com/matiasvarela/errors"
	"hexagonal/internal/core/domain"
	"hexagonal/internal/core/ports"
	"hexagonal/pkg/apperrors"
	"hexagonal/pkg/uidgen"
)

type service struct {
	gamesRepository ports.GamesRepository
	uidGen uidgen.UIDGen
}

func New(gamesRepository ports.GamesRepository, uidgen uidgen.UIDGen) *service {
	return &service{
		gamesRepository: gamesRepository,
		uidGen: uidgen,
	}
}

func (srv *service) Get(id string) (domain.Game, error) {
	game, err := srv.gamesRepository.Get(id)
	if err != nil {
		if errors.Is(err, apperrors.NotFound) {
			return domain.Game{}, errors.New(apperrors.NotFound, err, "Game not found !!!", "")
		}

		return domain.Game{}, errors.New(apperrors.Internal, err, "Get game from repository has failed !!!", "")
	}

	game.Board = game.Board.HideBombs()

	return game, nil
}

func (srv *service) Create(name string, size uint, bombs uint) (domain.Game, error) {
	if bombs >= size*size {
		return domain.Game{}, errors.New(apperrors.InvalidInput, nil, "The number of bombs is too high !!!", "")
	}

	game := domain.NewGame(srv.uidGen.New(), name, size, bombs)

	if err := srv.gamesRepository.Save(game); err != nil {
		return domain.Game{}, errors.New(apperrors.Internal, err, "Create game into repository has failed !!!", "")
	}

	game.Board = game.Board.HideBombs()

	return game, nil
}

func (srv *service) Reveal(id string, row uint, col uint) (domain.Game, error) {
	game, err := srv.gamesRepository.Get(id)
	if err != nil {
		if errors.Is(err, apperrors.NotFound) {
			return domain.Game{}, errors.New(apperrors.NotFound, err, "Game not found !!!", "")
		}

		return domain.Game{}, errors.New(apperrors.Internal, err, "Get game from repository has failed !!!", "")
	}

	if !game.Board.IsValidPosition(row, col) {
		return domain.Game{}, errors.New(apperrors.InvalidInput, nil, "Invalid position !!!", "")
	}

	if game.IsOver() {
		return domain.Game{}, errors.New(apperrors.IllegalOperation, nil, "Game is over !!!", "")
	}

	if game.Board.Contains(row, col, domain.CELL_BOMB) {
		game.State = domain.GAME_STATE_LOST
	} else {
		game.Board.Set(row, col, domain.CELL_REVEALED)

		if !game.Board.HasEmptyCells() {
			game.State = domain.GAME_STATE_WON
		}
	}

	if err := srv.gamesRepository.Save(game); err != nil {
		return domain.Game{}, errors.New(apperrors.Internal, err, "Update game into repository has failed !!!", "")
	}

	game.Board = game.Board.HideBombs()

	return game, nil
}

