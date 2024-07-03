package darts

func Score(x, y float64) int {
  if x<1 && y<1{
    return 10
  }
  return 0
}
