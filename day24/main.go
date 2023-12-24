package day24

import (
	"fmt"
	aoc "github.com/eveenendaal/advent-of-code-2023/aoc"
	"math"
	"strconv"
	"strings"
)

type Hailstone struct {
	point    Vector3D
	velocity Vector3D
}

func NewHailstone(line string) *Hailstone {
	// split on @
	parts := strings.Split(line, "@")
	// Parse position
	positionParts := strings.Split(parts[0], ",")
	px, _ := strconv.Atoi(strings.TrimSpace(positionParts[0]))
	py, _ := strconv.Atoi(strings.TrimSpace(positionParts[1]))
	pz, _ := strconv.Atoi(strings.TrimSpace(positionParts[2]))

	// Parse velocity
	velocityParts := strings.Split(parts[1], ",")
	vx, _ := strconv.Atoi(strings.TrimSpace(velocityParts[0]))
	vy, _ := strconv.Atoi(strings.TrimSpace(velocityParts[1]))
	vz, _ := strconv.Atoi(strings.TrimSpace(velocityParts[2]))

	// Create Hailstone
	return &Hailstone{
		point:    Vector3D{float64(px), float64(py), float64(pz)},
		velocity: Vector3D{float64(vx), float64(vy), float64(vz)},
	}
}

type Vector3D struct {
	x, y, z float64
}

type Vector2D struct {
	x, y float64
}

func (v Vector3D) to2D() Vector2D {
	return Vector2D{v.x, v.y}
}

func (v Vector3D) Add(w Vector3D) Vector3D {
	return Vector3D{v.x + w.x, v.y + w.y, v.z + w.z}
}

func (v Vector3D) Subtract(w Vector3D) Vector3D {
	return Vector3D{v.x - w.x, v.y - w.y, v.z - w.z}
}

func (v Vector3D) Multiply(scalar float64) Vector3D {
	return Vector3D{v.x * scalar, v.y * scalar, v.z * scalar}
}

func (v Vector3D) Dot(w Vector3D) float64 {
	return v.x*w.x + v.y*w.y + v.z*w.z
}

func (v Vector3D) Cross(w Vector3D) Vector3D {
	return Vector3D{
		v.y*w.z - v.z*w.y,
		v.z*w.x - v.x*w.z,
		v.x*w.y - v.y*w.x,
	}
}

func (v Vector3D) Length() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y + v.z*v.z)
}

func (v Vector2D) Add(w Vector2D) Vector2D {
	return Vector2D{v.x + w.x, v.y + w.y}
}

func (v Vector2D) Subtract(w Vector2D) Vector2D {
	return Vector2D{v.x - w.x, v.y - w.y}
}

func (v Vector2D) Multiply(scalar float64) Vector2D {
	return Vector2D{v.x * scalar, v.y * scalar}
}

func (v Vector2D) Dot(w Vector2D) float64 {
	return v.x*w.x + v.y*w.y
}

func (v Vector2D) Length() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func IntersectionPoint2D(P1, P2, D1, D2 Vector2D) (Vector2D, bool) {
	// Calculate the direction cross product.
	Dcross := D1.x*D2.y - D1.y*D2.x

	// If the cross product is zero, the lines are parallel and there's no intersection.
	if Dcross == 0 {
		return Vector2D{}, false
	}

	// Calculate the parameter for the intersection point.
	t := ((P2.x-P1.x)*D2.y - (P2.y-P1.y)*D2.x) / Dcross

	// Calculate the intersection point.
	P := P1.Add(D1.Multiply(t))

	return P, true
}

func IntersectionPoint3D(P1, P2, D1, D2 Vector3D) (Vector3D, bool) {
	// Calculate the direction cross product.
	Dcross := D1.Cross(D2)

	// If the cross product is zero, the lines are parallel and there's no intersection.
	if Dcross.Length() == 0 {
		return Vector3D{}, false
	}

	// Calculate the parameter for the intersection point.
	t := ((P2.Subtract(P1)).Cross(D2)).Dot(Dcross) / math.Pow(Dcross.Length(), 2)

	// Calculate the intersection point.
	P := P1.Add(D1.Multiply(t))

	return P, true
}

func Part1(filePath string) int {
	lines := aoc.ReadFileToLines(filePath)
	hailstones := make([]*Hailstone, 0)

	for _, line := range lines {
		hailstone := NewHailstone(line)
		hailstones = append(hailstones, hailstone)
	}

	minX := float64(7)
	maxX := float64(27)
	minY := float64(7)
	maxY := float64(27)

	intersections := 0
	// compare every pair of hailstones
	for i := 0; i < len(hailstones); i++ {
		for j := i + 1; j < len(hailstones); j++ {
			// calculate intersection point
			P, ok := IntersectionPoint2D(hailstones[i].point.to2D(), hailstones[j].point.to2D(), hailstones[i].velocity.to2D(), hailstones[j].velocity.to2D())
			if ok {
				// if point with x and y coordinates is within the bounds of the hailstones, there is an intersection
				if P.x >= minX && P.x <= maxX && P.y >= minY && P.y <= maxY {
					intersections++
					fmt.Printf("Intersection point: %v\n", P)
				} else {
					fmt.Println("Intersection outside range")
				}
			} else {
				fmt.Println("No intersection")
			}
		}
	}

	return intersections
}

func main() {
	P1 := Vector3D{1, 0, 0}
	D1 := Vector3D{1, 1, 1}

	P2 := Vector3D{0, 1, 0}
	D2 := Vector3D{-1, -1, -1}

	P, ok := IntersectionPoint3D(P1, P2, D1, D2)
	if ok {
		fmt.Printf("Intersection point: %v\n", P)
	} else {
		fmt.Println("No intersection")
	}
}
