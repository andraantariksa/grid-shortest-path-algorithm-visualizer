package grid

type BoxState int32

const (
	BOX_STATE_EMPTY BoxState = 0
	BOX_STATE_WALL  BoxState = 1
	BOX_STATE_START BoxState = 2
	BOX_STATE_END   BoxState = 3
)
