package basics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGame(t *testing.T) {
	g := newGame(3, "minecraft", 20, "sandbox")

	assert.Equal(t, g.id, 3, "id")
	assert.Equal(t, g.name, "minecraft", "name")
	assert.Equal(t, g.price, 20, "price")
	assert.Equal(t, g.genre, "sandbox", "genre")
}

func TestString(t *testing.T) {
	g := newGame(4, "warcraft", 40, "strategy")

	assert.Equal(t, g.item.String(), "4: warcraft costs 40", "item string")
	assert.Equal(t, g.String(), "Game 4: warcraft costs 40 of genre strategy", "game string")
}

func TestList(t *testing.T) {
	games := newGameList()

	assert.Len(t, games, 3, "gamelist len")

 	assert.Equal(t, games[0].id, 1, "game 0 id")
	assert.Equal(t, games[0].name, "god of war", "game 0 name")
	assert.Equal(t, games[0].price, 50, "game 0 price")
	assert.Equal(t, games[0].genre, "action adventure", "game 0 genre")

 	assert.Equal(t, games[1].id, 3, "game 1 id")
	assert.Equal(t, games[1].name, "minecraft", "game 1 name")
	assert.Equal(t, games[1].price, 20, "game 1 price")
	assert.Equal(t, games[1].genre, "sandbox", "game 1 genre")

 	assert.Equal(t, games[2].id, 4, "game 2 id")
	assert.Equal(t, games[2].name, "warcraft", "game 2 name")
	assert.Equal(t, games[2].price, 40, "game 2 price")
	assert.Equal(t, games[2].genre, "strategy", "game 2 genre")
}

func TestById(t *testing.T) {
	g, err := queryById(newGameList(), 1)
	assert.NoError(t, err, "no error")
 	assert.Equal(t, g.id, 1, "game id")
	assert.Equal(t, g.name, "god of war", "game name")
	assert.Equal(t, g.price, 50, "game price")
	assert.Equal(t, g.genre, "action adventure", "game genre")

	g, err = queryById(newGameList(), 11)
	assert.EqualError(t, err, "No such game", "error")
}

func TestNameByPrice(t *testing.T) {
	names := listNameByPrice(newGameList(), 40)

	assert.Len(t, names, 2, "namelist len")
	if 2 > 0 {
	 	assert.Equal(t, names[0], "minecraft", "names 1")
	}
	if 2 > 1 {
	 	assert.Equal(t, names[1], "warcraft", "names 2")
	}
	if 2 > 2 {
	 	assert.Equal(t, names[2], "N/A", "names 3")
	}
	if 2 > 3 {
	 	assert.Equal(t, names[3], "N/A", "names 4")
	}
}

