package enums

type Season int

const (
	Spring Season = iota
	Summer
	Autumn
	Winter
)

func GetWeatherAdvisory(s Season) string {
	var stateName string

	switch s {
	case Spring:
		stateName = "Bring an umbrella"
	case Summer:
		stateName = "Wear sunscreen"
	case Autumn:
		stateName = "Bring a light jacket"
	case Winter:
		stateName = "Wear a heavy coat"
	default:
		stateName = "Unknown season"
	}

	return stateName
}
