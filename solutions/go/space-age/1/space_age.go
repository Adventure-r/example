package spaceage

type Planet string

const ageOnEarthInSecs = 31557600

func Age(seconds float64, planet Planet) float64 {
	var multipliar float64
	switch planet {
	case "Mercury":
		multipliar = 0.2408467
	case "Venus":
		multipliar = 0.61519726
	case "Earth":
		multipliar = 1
	case "Mars":
		multipliar = 1.8808158
	case "Jupiter":
		multipliar = 11.862615
	case "Saturn":
		multipliar = 29.447498
	case "Uranus":
		multipliar = 84.016846
	case "Neptune":
		multipliar = 164.79132
	default:
		return -1
	}
	return seconds / ageOnEarthInSecs / multipliar
}
