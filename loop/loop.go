package main

import "fmt"

func main() {
  x := 7
  table := [][]int{ {1, 5, 6}, {2, 6, 3, 12}, {5, 3, 7, 9} }
  found := false

  for row := 0; row < len(table); row++ {
    for col := 0; col < len(table[row]); col++ {
      if table[row][col] == x {
        found = true;
        fmt.Printf("found %d(row:%d, col:%d)\n", x, row, col)
        break
      }
      if found {
        break
      }
    }
  }

}
