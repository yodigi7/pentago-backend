package block

import (
    "testing"
    "internal/marble"
)

func TestReverseRows(t *testing.T) {
    var p [3][3]marble.Marble
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if i == 0 {
                p[i][j] = marble.WhiteMarble
            } else {
                p[i][j] = marble.EmptyMarble
            }
        }
    }
    b := Block{p}
    b = b.reverseRows()
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if i == 2 && b.Places[i][j] != marble.WhiteMarble {
                t.Errorf("White marble not found at %d, %d", i, j)
            } else if i != 2 && b.Places[i][j] != marble.EmptyMarble {
                t.Errorf("Empty marble not found at %d, %d", i, j)
            }
        }
    }
}

func TestTranspose(t *testing.T) {
    var p [3][3]marble.Marble
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if i == 0 {
                p[i][j] = marble.WhiteMarble
            } else {
                p[i][j] = marble.EmptyMarble
            }
        }
    }
    b := Block{p}
    b = b.transpose()
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if j == 0 && b.Places[i][j] != marble.WhiteMarble {
                t.Errorf("White marble not found at %d, %d", i, j)
            } else if j != 0 && b.Places[i][j] != marble.EmptyMarble {
                t.Errorf("Empty marble not found at %d, %d", i, j)
            }
        }
    }
}

func TestRotateClockwise(t *testing.T) {
    var p [3][3]marble.Marble
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if i == 0 {
                p[i][j] = marble.WhiteMarble
            } else {
                p[i][j] = marble.EmptyMarble
            }
        }
    }
    b := Block{p}
    b = b.RotateClockwise()
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if j == 2 && b.Places[i][j] != marble.WhiteMarble {
                t.Errorf("White marble not found at %d, %d", i, j)
            } else if j != 2 && b.Places[i][j] != marble.EmptyMarble {
                t.Errorf("Empty marble not found at %d, %d", i, j)
            }
        }
    }
}

func TestRotateCounterClockwise(t *testing.T) {
    var p [3][3]marble.Marble
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if i == 0 {
                p[i][j] = marble.WhiteMarble
            } else {
                p[i][j] = marble.EmptyMarble
            }
        }
    }
    b := Block{p}
    b = b.RotateCounterClockwise()
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if j == 0 && b.Places[i][j] != marble.WhiteMarble {
                t.Errorf("White marble not found at %d, %d", i, j)
            } else if j != 0 && b.Places[i][j] != marble.EmptyMarble {
                t.Errorf("Empty marble not found at %d, %d", i, j)
            }
        }
    }
}

func TestRotatesCancelEachOtherOut(t *testing.T) {
    var p [3][3]marble.Marble
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if i == 0 {
                p[i][j] = marble.WhiteMarble
            } else {
                p[i][j] = marble.EmptyMarble
            }
        }
    }
    b := Block{p}
    b = b.RotateCounterClockwise().RotateClockwise()
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if i == 0 && b.Places[i][j] != marble.WhiteMarble {
                t.Errorf("White marble not found at %d, %d", i, j)
            } else if i != 0 && b.Places[i][j] != marble.EmptyMarble {
                t.Errorf("Empty marble not found at %d, %d", i, j)
            }
        }
    }
}
