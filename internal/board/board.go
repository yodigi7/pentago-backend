package board

import (
	"fmt"
	"internal/block"
	"internal/marble"
)

type Board struct {
	Blocks [2][2]block.Block
}

// TODO: have t be type Marble rather than type int
func (b Board) PlaceMarble(c, d, x, y, t int) Board {
	// FIXME: c,d are the coords of the block, and x,y are coords within the block, t is type of marble (b/w). Dunno if there's better convention for these var names.
	// TODO: Make it so a player can't try to put a marble where another marble has already been placed.
	marbleTypeUsed := marble.EmptyMarble
	if t == 0 {
		marbleTypeUsed = marble.BlackMarble
	} else if t == 1 {
		marbleTypeUsed = marble.WhiteMarble
	} else {
		fmt.Println("Error! internal/board/board.go func PlaceMarble got passed in an invalid value for t (b/w marble!)")
		// FIXME: Why won't this error message work instead of the simple Println?
		// t.Errorf("Error! internal/board/board.go func PlaceMarble got passed in an invalid value for t (b/w marble!)")
	}
	blockUsed := b.Blocks[c][d]
	blockUsed.Places[x][y] = marbleTypeUsed

	return b
}

