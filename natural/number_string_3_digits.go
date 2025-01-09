package main
import (
"fmt"
"strings"
)
// Funktion, die eine Zahl zwischen 20 und 999 in Text umwandelt
func zahlInText(n int) string {
if n < 20 || n > 999 {
return "Ungültige Zahl"
}
// Arrays für die Zahlwörter
"neunzehn"}
einheiten := []string{"", "eins", "zwei", "drei", "vier", "fünf", "sechs", "sieben", "acht", "neun", "zehn", "elf", "zwölf", "dreizehn", "vierzehn", "fünfzehn", "sechzehn", "siebzehn", "achtzehn",
ziger := []string{"", "", "zwanzig", "dreißig", "vierzig", "fünfzig", "sechzig", "siebzig", "achtzig", "neunzig"}
var result strings.Builder
if n >= 100 {
hunderter := n / 100
result.WriteString(einheiten[hunderter])
result.WriteString("hundert")
n %= 100
if n > 0 {
result.WriteString(" und ")
}
}
if n >= 20 {
zehner := n / 10
result.WriteString(ziger[zehner])
n %= 10
if n > 0 {
result.WriteString("-")
}
}
}
if n > 0 {
result.WriteString(einheiten[n])
return result.String()
}
func main() {
// Alle Zahlen von 20 bis 999 ausgeben
for i := 20; i <= 999; i++ {
fmt.Println(i, ":", zahlInText(i))
}
}