package grid

import (
	"log"

	"gitlab.com/andraantariksa/grid-shortest-path-visualizer/data-structure/queue"
)

func (grd *Grid) SolveBFS() *queue.ItemQueue {
	q := queue.New()

	visited := make([][]bool, grd.Size, grd.Size)
	for i := range visited {
		visited[i] = make([]bool, grd.Size, grd.Size)
		for j := range visited[i] {
			if grd.BoxesState[i][j] == BOX_STATE_WALL {
				visited[i][j] = true
			} else {
				visited[i][j] = false
			}
		}
	}

	visited[0][0] = true
	q.Enqueue(coord{
		x: 0,
		y: 0,
	})

	for !q.IsEmpty() {
		p := (*q.Front()).(coord)
		q.Dequeue()

		if grd.BoxesState[p.y][p.x] == BOX_STATE_END {
			log.Println("end!")
			log.Println(p.y, p.x, q)
			return q
		}

		if p.y != 0 {
			if !visited[p.y-1][p.x] {
				q.Enqueue(coord{
					x: p.x,
					y: p.y - 1,
				})
				visited[p.y-1][p.x] = true
			}
		}

		if p.y+1 < grd.Size {
			if !visited[p.y+1][p.x] {
				q.Enqueue(coord{
					x: p.x,
					y: p.y + 1,
				})
				visited[p.y+1][p.x] = true
			}
		}

		if p.x != 0 {
			if !visited[p.y][p.x-1] {
				q.Enqueue(coord{
					x: p.x - 1,
					y: p.y,
				})
				visited[p.y][p.x-1] = true
			}
		}

		if p.x+1 < grd.Size {
			if !visited[p.y][p.x+1] {
				q.Enqueue(coord{
					x: p.x + 1,
					y: p.y,
				})
				visited[p.y][p.x+1] = true
			}
		}

	}

	return q
}
