package darts

func Score(x, y float64) int {
  if x<1 && y<1{
    return 10
  }
  if x>10 || y>10{
    return 0
  }
  if (x>5 && x<=10) || (y>5 && y<=10){
    return 1
  }
  return 0
}
