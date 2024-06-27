package space

type Planet string

func Orbit(planet Planet) float64{
  if planet == "Mercury"{
    return 0.2408467
  }
  if planet == "Venus"{
    return 0.61519726
  }
  if planet == "Earth"{
    return 0.61519726
  }
  if planet == "Mars"{
    return 0.61519726
  }
  if planet == "Jupiter"{
    return 0.61519726
  }

}

func Age(seconds float64, planet Planet) float64 {
	panic("Please implement the Age function")

}
