package board

import (
	"pentago-backend/internal/block"
	"pentago-backend/internal/marble"
	"testing"
)

func TestPlaceMarble(t *testing.T) {
	var s [2][2]block.Block
	b := Board{s}
	b.PlaceMarble(0, 1, 2, 0, 0) // *Should* put marble in the top right corner of the bottom left block.

	// Added this check temporarily since your test wasn't correctly working
	if b.Blocks[0][1].Places[2][0] != marble.BlackMarble {
		t.Errorf("Black marble not found")
	}

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 3; k++ {
				for l := 0; l < 3; l++ {
					// placement := b.Blocks[i][j].Places[k][l] // FIXME: Is it easier to use shorthand like this, or
					if i == 0 && j == 1 {
						if k == 2 && j == 0 && b.Blocks[i][j].Places[k][l] != marble.BlackMarble { // FIXME  This if statement should never be hit because the previous if statement checks to make sure that j=1 and this checks for j=0 so one will be false making this if statement impossible to enter
							t.Errorf("Black marble not found at %d, %d", k, l)
						} else if b.Blocks[i][j].Places[k][l] != marble.EmptyMarble {
							t.Errorf("Empty marble not found at %d, %d", k, l)
						}
					} else {
						if b.Blocks[i][j].Places[k][l] != marble.EmptyMarble {
							t.Errorf("Empty marble not found at %d, %d", k, l)
						}
					}
				}
			}
		}
	}
}
