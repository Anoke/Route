package main

import (
	"fmt"
)

type Point struct {
	x, y int
}

func main() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var n, m int
		fmt.Scan(&n, &m)

		var robotA, robotB Point
		var field []string

		for j := 0; j < n; j++ {
			var row string
			fmt.Scan(&row)
			field = append(field, row)

			for k, chr := range row {
				switch chr {
				case 'A':
					robotA = Point{j, k}
				case 'B':
					robotB = Point{j, k}
				}
			}
		}

		if (robotA == Point{0, 0} && robotB == Point{n - 1, m - 1}) || (robotA == Point{n - 1, m - 1} && robotB == Point{0, 0}) {
			for j := 0; j < n; j++ {
				fmt.Println(field[j])
			}
		} else {
			switch {
			case robotA.x < robotB.x:
				moveRobo(&field, robotA, 'a', n, m, 0)
				moveRobo(&field, robotB, 'b', n, m, 1)
			case robotB.x < robotA.x:
				moveRobo(&field, robotB, 'b', n, m, 0)
				moveRobo(&field, robotA, 'a', n, m, 1)
			case robotA.x == robotB.x:
				if robotA.y < robotB.y {
					moveRobo(&field, robotA, 'a', n, m, 0)
					moveRobo(&field, robotB, 'b', n, m, 1)
				} else {
					moveRobo(&field, robotB, 'b', n, m, 0)
					moveRobo(&field, robotA, 'a', n, m, 1)
				}
			}

			for j := 0; j < n; j++ {
				fmt.Println(field[j])
			}
		}
	}
}

func moveRobo(field *[]string, robot Point, moveChar byte, n, m, direction int) {
	if direction == 0 {
		if robot.y%2 != 0 {
			robot.y--
			(*field)[robot.x] = (*field)[robot.x][:robot.y] + string(moveChar) + (*field)[robot.x][robot.y+1:]
		}
		robot.x--
		for ; robot.x >= 0; robot.x-- {
			(*field)[robot.x] = (*field)[robot.x][:robot.y] + string(moveChar) + (*field)[robot.x][robot.y+1:]
		}
		robot.x++
		robot.y--
		for ; robot.y >= 0; robot.y-- {
			(*field)[robot.x] = (*field)[robot.x][:robot.y] + string(moveChar) + (*field)[robot.x][robot.y+1:]
		}
	} else {
		if robot.y%2 != 0 {
			robot.y++
			(*field)[robot.x] = (*field)[robot.x][:robot.y] + string(moveChar) + (*field)[robot.x][robot.y+1:]
		}
		robot.x++
		for ; robot.x < n; robot.x++ {
			(*field)[robot.x] = (*field)[robot.x][:robot.y] + string(moveChar) + (*field)[robot.x][robot.y+1:]
		}
		robot.x--
		robot.y++
		for ; robot.y < m; robot.y++ {
			(*field)[robot.x] = (*field)[robot.x][:robot.y] + string(moveChar) + (*field)[robot.x][robot.y+1:]
		}
	}
}
