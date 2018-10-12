package main

import (
  "fmt"
)

const (
  PLAYER_X = 1
  PLAYER_O = 2
)

var (
  grid = [3][3]int{{3, 4, 5}, {6, 7, 8}, {9, 10, 11}}
)

func main() {
  isX := true

  for i := 0; i < 9; i++ {
    fmt.Print("Enter x, y coordinates: ")
    var x, y int
    fmt.Scan(&x, &y)

    if (isX) {
      grid[y][x] = PLAYER_X
    } else {
      grid[y][x] = PLAYER_O
    }

    isX = !isX

    fmt.Println("Grid: ")
    drawGrid()

    if (checkWin()) {
      fmt.Println("Win")
      return
    }
  }
}

func drawGrid() {
  for y := 0; y < 3; y++ {
		for x := 0; x < 3; x++ {
      if (grid[x][y] == PLAYER_X) {
        fmt.Print("X")
      } else if (grid[x][y] == PLAYER_O) {
        fmt.Print("O")
      } else {
        fmt.Print(" ")
      }
      fmt.Print(" | ")
    }
    fmt.Println()
  }
}

func checkWin() (status bool) {
  return (
    // Horisontal lines
    (grid[0][0] == grid[0][1] && grid[0][1] == grid[0][2]) ||
    (grid[1][0] == grid[1][1] && grid[1][1] == grid[1][2]) ||
    (grid[2][0] == grid[2][1] && grid[2][1] == grid[2][2]) ||

    // Vertical lines 
    (grid[0][0] == grid[1][0] && grid[1][0] == grid[2][0]) ||
    (grid[0][1] == grid[1][1] && grid[1][1] == grid[2][1]) ||
    (grid[0][2] == grid[1][2] && grid[1][2] == grid[2][2]) ||

    // X
    (grid[0][0] == grid[1][1] && grid[1][1] == grid[2][2]) ||
    (grid[2][0] == grid[1][1] && grid[1][1] == grid[0][2]))
}