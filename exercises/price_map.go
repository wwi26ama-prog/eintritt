package exercises

// AdmissionPriceFromMap erwartet ein Alter in Jahren und eine
// Preistabelle als Map vom Mindestalter auf den Preis, z.B.
// `map[int]int{0: 0, 6: 300, 18: 500}`. Alter und Mindestalter
// sind nicht negativ. Ansonsten gilt dasselbe wie für
// `AdmissionPriceFromLists`.
//
// Bonusaufgabe.
func AdmissionPriceFromMap(age int, prices map[int]int) int {
	// Hinweis:
	// Eine Schleife mit `range` über eine Map liefert Schlüssel und Wert,
	// aber in keiner festen Reihenfolge.
	// - Sie können sich deshalb nicht darauf verlassen, dass das passende
	//   Mindestalter als letztes kommt.
	// - Merken Sie sich neben dem Preis auch das größte bisher passende
	//   Mindestalter, ähnlich wie beim Maximum einer Liste.

	price := 0
	bestMinAge := -1
	for minAge, minAgePrice := range prices {
		if minAge <= age && minAge > bestMinAge {
			bestMinAge = minAge
			price = minAgePrice
		}
	}
	return price
	// Anmerkung:
	// Startwert `-1`:
	// Da die Mindestalter nicht negativ sind, ist jedes passende
	// Mindestalter größer als `-1`, auch `0`.
}
