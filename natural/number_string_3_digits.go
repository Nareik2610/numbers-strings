package natural

import (
	"fmt"
	"strings"
)

// Arrays für die deutsche Übersetzung von Zahlen
var einsen = []string{
	"null", "eins", "zwei", "drei", "vier", "fünf", "sechs", "sieben", "acht", "neun",
}

var zehner = []string{
	"", "", "zwanzig", "dreißig", "vierzig", "fünfzig", "sechzig", "siebzig", "achtzig", "neunzig",
}

var zweistelligeEinheiten = []string{
	"", "zehn", "elf", "zwölf", "dreizehn", "vierzehn", "fünfzehn", "sechzehn", "siebzehn", "achtzehn", "neunzehn",
}

var hunderter = []string{
	"", "hundert", "zweihundert", "dreihundert", "vierhundert", "fünfhundert", "sechshundert", "siebenhundert", "achthundert", "neunhundert",
}

func main() {
	// Iteriere über alle Zahlen von 0 bis 999
	for i := 0; i <= 999; i++ {
		fmt.Println(NumberString3Digits(i))
	}
}

// Funktion, um eine Zahl in ihre deutsche Schreibweise zu übersetzen
func NumberString3Digits(n int) string {
	if n == 0 {
		return einsen[0]
	}

	var result strings.Builder

	// Hundertstellen
	h := n / 100
	if h > 0 {
		result.WriteString(einsen[h])
		result.WriteString("hundert")
	}

	// Zehnerstellen
	z := (n % 100) / 10
	e := n % 10

	// Wenn die Zahl zwischen 10 und 19 liegt
	if n >= 10 && n <= 19 {
		result.WriteString(zweistelligeEinheiten[n-10])
	} else {
		// Bei anderen Zahlen die Zehner und Einer separat behandeln
		if z > 1 {
			result.WriteString(zehner[z])
		} else if z == 1 {
			// Zehner „zehn“ muss speziell behandelt werden
			if e > 0 {
				result.WriteString("und")
			}
		}
		// Einerstelle
		if e > 0 || n == 0 {
			result.WriteString(einsen[e])
		}
	}

	return result.String()
}
