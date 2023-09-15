package block

import "internal/marble"

type Block struct {
    Places [3][3]marble.Marble
}

func (b Block) RotateCounterClockwise() Block {
    return b.transpose().reverseRows()
}

func (b Block) RotateClockwise() Block {
    return b.reverseRows().transpose()
}

func reverseRow(row [3]marble.Marble) [3]marble.Marble {
    row[0], row[2] = row[2], row[0]
    return row
}

func (b Block) reverseRows() Block {
    b.Places[0], b.Places[2] = b.Places[2], b.Places[0]
    return b
}

func (b Block) reverseColumns() Block {
    for i := 0; i < 3; i++ {
        b.Places[i] = reverseRow(b.Places[i])
    }
    return b
}


func (b Block) transpose() Block {
    for i := 0; i < 3; i++ {
        for j := i; j < 3; j++ {
            if i == j {
                continue
            }
            b.Places[i][j], b.Places[j][i] = b.Places[j][i], b.Places[i][j]
        }
    }
    return b
}
