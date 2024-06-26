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
    return 1
  }
  if planet == "Mars"{
    return 1.8808158
  }
  if planet == "Jupiter"{
    return 11.862615
  }
  if planet == "Saturn"{
    return 29.447498
  }
  if planet == "Uranus"{
    return 84.016846
  }
  if planet == "Neptune"{
    return 164.79132
  }
  return -1.000000
}

func Age(seconds float64, planet Planet) float64 {
  orbit_multiplier := Orbit(planet)
  if orbit_multiplier != -1.000000{
    return seconds/(31557600*orbit_multiplier)
  }
  return orbit_multiplier
}
