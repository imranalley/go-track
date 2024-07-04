package darts

func Score(x, y float64) int {
  abs_diff := x-y
  if abs_diff<0 {
    abs_diff=-abs_diff
  }
  if x<1 && y<1{
    return 10
  }
  if x>10 || y>10{
    return 0
  }
  if (x>5 && x<=10) || (y>5 && y<=10){
    return 1
  }
  if abs_diff>10{
    return 0
  }
  return 0
}
