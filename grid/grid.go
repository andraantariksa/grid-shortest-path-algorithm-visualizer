package grid

import "github.com/gotk3/gotk3/gtk"

type BoxState int32

const (
	BOX_STATE_EMPTY BoxState = 0
	BOX_STATE_WALL  BoxState = 1
	BOX_STATE_START BoxState = 2
	BOX_STATE_END   BoxState = 3
)

type Grid struct {
	BoxesState [16][16]BoxState
	BoxSize    float64
	Size       uint
	Window     *gtk.ApplicationWindow
}

type coord struct {
	x uint
	y uint
}

func New(win *gtk.ApplicationWindow) *Grid {
	grd := Grid{
		Size:    16,
		BoxSize: 25.0,
		Window:  win,
	}
	for y := uint(0); y < grd.Size; y++ {
		for x := uint(0); x < grd.Size; x++ {
			grd.BoxesState[y][x] = BOX_STATE_EMPTY
		}
	}
	return &grd
}

func (grd *Grid) SetStart(x, y uint) {
	grd.BoxesState[y][x] = BOX_STATE_START
}

func (grd *Grid) SetEnd(x, y uint) {
	grd.BoxesState[y][x] = BOX_STATE_END
}
