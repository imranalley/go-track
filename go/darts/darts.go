package darts

import "math"

func Score(x, y float64) int {
  abs_diff := math.Sqrt(x*x + y*y)
  
  if abs_diff<1 {
    return 10
  }
  if abs_diff <=5 {
    return 5
  }
  if abs_diff <=10 {
    return 1
  }
  return 0
}
