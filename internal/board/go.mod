module board

go 1.18

require internal/marble v1.0.0 // indirect

replace internal/marble => ../marble

require internal/block v1.0.0

replace internal/block => ../block
