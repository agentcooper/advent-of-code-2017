package Day20

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"regexp"

	"github.com/agentcooper/advent-of-code-2017/utils"
)

// Part is either Part1 or Part2
type Part int

const (
	// Part1 is part 1
	Part1 Part = 1
	// Part2 is part 2
	Part2 Part = 2
)

type Vector3 struct {
	x int
	y int
	z int
}

type Particle struct {
	position     Vector3
	velocity     Vector3
	acceleration Vector3

	dead bool
}

func (p *Particle) tick() {
	p.velocity.x += p.acceleration.x
	p.velocity.y += p.acceleration.y
	p.velocity.z += p.acceleration.z

	p.position.x += p.velocity.x
	p.position.y += p.velocity.y
	p.position.z += p.velocity.z
}

func (p *Particle) distance() int {
	return utils.Abs(p.position.x) + utils.Abs(p.position.y) + utils.Abs(p.position.z)
}

func initParticle(s string) Particle {
	ParticleRx :=
		regexp.MustCompile(`p=<(-?\d+),(-?\d+),(-?\d+)>, v=<(-?\d+),(-?\d+),(-?\d+)>, a=<(-?\d+),(-?\d+),(-?\d+)>`)

	submatch := ParticleRx.FindStringSubmatch(s)

	if len(submatch) != 1+3*3 {
		panic(fmt.Errorf("Bad input (can't match): %s", s))
	}

	ints, err := utils.StringsToInts(submatch[1:])
	if err != nil {
		panic(fmt.Errorf("Bad input (can't convert to ints): %s", submatch))
	}

	return Particle{
		position:     Vector3{x: ints[0], y: ints[1], z: ints[2]},
		velocity:     Vector3{x: ints[3], y: ints[4], z: ints[5]},
		acceleration: Vector3{x: ints[6], y: ints[7], z: ints[8]},
	}
}

// Solve solves the puzzle
func Solve(r io.Reader, part Part) int {
	const debug = false

	scanner := bufio.NewScanner(r)

	particles := []Particle{}

	for scanner.Scan() {
		line := scanner.Text()
		particle := initParticle(line)
		particles = append(particles, particle)
	}

	const iterations = 10000

	min := math.MaxInt64
	var winner int

	for k := 0; k < iterations; k++ {
		if debug {
			for i := range particles {
				fmt.Printf("[%d], Particle %d: %+v\n", k, i, particles[i])
			}
			fmt.Println("---")
		}

		for i := range particles {
			if particles[i].dead {
				continue
			}

			particles[i].tick()
		}

		if part == Part2 {
			pos := map[Vector3][]int{}

			for i := range particles {
				if particles[i].dead {
					continue
				}

				pos[particles[i].position] = append(pos[particles[i].position], i)
			}

			for _, v := range pos {
				if len(v) > 1 {
					// fmt.Printf("Dead: %v\n", v)
					for _, j := range v {
						particles[j].dead = true
					}
				}
			}
		}
	}

	if part == Part1 {
		for i := range particles {
			d := particles[i].distance()
			if d < min {
				min = d
				winner = i
			}
		}

		return winner
	}

	if part == Part2 {
		count := 0

		for _, p := range particles {
			if !p.dead {
				count++
			}
		}

		return count
	}

	panic("Unknown part")
}
