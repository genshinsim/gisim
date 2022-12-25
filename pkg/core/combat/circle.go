package combat

import (
	"fmt"
	"math"
)

type Circle struct {
	center   Point
	r        float64
	dir      Point
	fanAngle float64
	segments []Point
}

func NewCircle(center Point, r float64, dir Point, fanAngle float64) *Circle {
	var segments []Point
	if fanAngle > 0 && fanAngle < 360 {
		segments = calcSegments(center, r, dir, fanAngle)
	}
	return &Circle{
		center:   center,
		r:        r,
		dir:      dir,
		fanAngle: fanAngle,
		segments: segments,
	}
}

func (c *Circle) Pos() Point {
	return c.center
}

func (c *Circle) SetPos(p Point) {
	if c.center == p {
		return
	}
	for i := 0; i < len(c.segments); i++ {
		c.segments[i] = c.segments[i].Add(p.Sub(c.center))
	}
	c.center = p
}

func (c *Circle) String() string {
	return fmt.Sprintf(
		"r: %v x: %v y: %v dir: %v fanAngle: %v segments: %v",
		c.r, c.center.X, c.center.Y, c.dir, c.fanAngle, c.segments,
	)
}

func calcSegments(center Point, r float64, dir Point, fanAngle float64) []Point {
	fanAngleRadian := fanAngle * math.Pi / 180
	// assume circle center is origin at first to do the rotation stuff
	segmentStart := Point{X: 0, Y: r}.Rotate(dir)
	segmentLeft := segmentStart.Rotate(Point{X: math.Sin(-fanAngleRadian / 2), Y: math.Cos(-fanAngleRadian / 2)})
	segmentRight := segmentStart.Rotate(Point{X: math.Sin(fanAngleRadian / 2), Y: math.Cos(fanAngleRadian / 2)})
	// save segment points (the circle center and segment point make up a line segment)
	// need to move segment to where the actual circle center is
	return []Point{segmentLeft.Add(center), segmentRight.Add(center)}
}

// TODO: this ignores the possibility of c1 also having a fanAngle (target with a partial circle hitbox...)
func (c1 *Circle) IntersectCircle(c2 Circle) bool {
	// https://stackoverflow.com/a/4226473
	// A: full circles have to be intersecting
	// (R0 - R1)^2 <= (x0 - x1)^2 + (y0 - y1)^2 <= (R0 + R1)^2
	radiusSum := c1.r + c2.r
	if c1.center.Sub(c2.center).MagnitudeSquared() >= radiusSum*radiusSum {
		return false
	}

	// c2 has no fanAngle -> there's an intersection if A
	if c2.segments == nil {
		return true
	}

	// c2 has a fanAngle -> there's an intersection if A && (B || C)
	// https://www.baeldung.com/cs/circle-line-segment-collision-detection
	// B: check if c1 intersects any of c2's segments, if yes we can exit early
	// (it's necessary to check for this because c1 can collide with c2's fanAngle area
	// even if c1's circle center isn't in c2's fanAngle range)
	for _, segment := range c2.segments {
		o := c1.center
		p := c2.center
		q := segment

		op := p.Sub(o)
		qp := p.Sub(q)
		oq := q.Sub(o)
		pq := q.Sub(p)

		opDist := o.Distance(p)
		oqDist := o.Distance(q)
		pqDist := p.Distance(q)

		minDist := math.Min(opDist, oqDist)
		maxDist := math.Max(opDist, oqDist)
		if op.Dot(qp) > 0 && oq.Dot(pq) > 0 {
			minDist = math.Abs(op.Cross(oq)) / pqDist
		}
		if minDist <= c1.r && maxDist >= c1.r {
			return true
		}
	}

	// C: check if the angle between the vector pointing from c2 to c1 and the y axis lies within the fanAngle of c2
	return fanAngleAreaCheck(c2.center, c1.center, c2.dir, c2.fanAngle)
}

func (c *Circle) IntersectRectangle(r Rectangle) bool {
	return IntersectRectangle(r, *c)
}
