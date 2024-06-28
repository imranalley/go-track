package space

type Planet string

func Orbit(planet Planet) float64{
  if planet == "Mercury"{
    return 0.2408467*31557600
  }
  if planet == "Venus"{
    return 0.61519726*31557600
  }
  if planet == "Earth"{
    return 1*31557600
  }
  if planet == "Mars"{
    return 1.8808158*31557600
  }
  if planet == "Jupiter"{
    return 11.862615*31557600
  }
  if planet == "Saturn"{
    return 29.447498*31557600
  }
  if planet == "Uranus"{
    return 84.016846*31557600
  }
  if planet == "Neptune"{
    return 164.79132*31557600
  }
  return 0
}

func Age(seconds float64, planet Planet) float64 {
  orbit_multiplier := Orbit(planet)
  return orbit_multiplier/seconds
}
