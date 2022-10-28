package domain

import (
	"testing"

	// "hexagonal/internal/core/domain"
	"github.com/stretchr/testify/assert"
)

// BOARD TEST

func TestNewBoard(t *testing.T) {
	size := uint(10)
	bombs := uint(50)

	board := NewBoard(size, bombs)

	countBombs := uint(0)
	for row := range board {
		for col := range board[0] {
			if board[row][col] == CELL_BOMB {
				countBombs++
			}
		}
	}
}

func TestBoard_HideBombs(t *testing.T) {
	board := NewBoard(10, 50).HideBombs()

	hasBombHided := true
	for i := range board {
		if !hasBombHided {
			break
		}

		for j := range board[0] {
			if board[i][j] != CELL_EMPTY {
				hasBombHided = false
				break
			}
		}
	}

	assert.True(t, hasBombHided)
}

func TestBoard_IsValidPosition(t *testing.T) {
	board := NewBoard(10, 50)

	assert.True(t, board.IsValidPosition(5, 5))
	assert.False(t, board.IsValidPosition(11, 5))
	assert.False(t, board.IsValidPosition(5, 11))
}

func TestBoard_Contains(t *testing.T) {
	board := NewBoard(10, 50)
	board[1][2] = CELL_BOMB

	assert.True(t, board.Contains(1, 2, CELL_BOMB))
}

func TestBoard_HasEmptyCell(t *testing.T) {
	boardWithEmptyCells := NewBoard(10, 50)
	boardWithoutEmptyCells := NewBoard(10, 100)

	assert.True(t, boardWithEmptyCells.HasEmptyCells())
	assert.False(t, boardWithoutEmptyCells.HasEmptyCells())
}

// ··· GAME TESTS ··· //

func TestNewGame(t *testing.T) {
	game := NewGame("1001", "my game", 10, 50)

	assert.Equal(t, "1001", game.ID)
	assert.Equal(t, "my game", game.Name)
	assert.Equal(t, BoardSettings{Size: 10, Bombs: 50}, game.BoardSettings)
	assert.Equal(t, GAME_STATE_NEW, game.State)
	assert.Equal(t, 10, len(game.Board))
	assert.Equal(t, 10, len(game.Board[0]))
}

func TestGame_IsOver(t *testing.T) {
	gameNew := NewGame("1001", "my game", 10, 50)

	gameWon := NewGame("1001", "my game", 10, 50)
	gameWon.State = GAME_STATE_WON

	gameLost := NewGame("1001", "my game", 10, 50)
	gameLost.State = GAME_STATE_LOST

	assert.False(t, gameNew.IsOver())
	assert.True(t, gameWon.IsOver())
	assert.True(t, gameLost.IsOver())
}