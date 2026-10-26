package exercises

// AdmissionPriceFromLists erwartet ein Alter in Jahren und eine
// Preistabelle als zwei gleich lange Listen: `prices[i]` gilt ab dem
// Alter `minAges[i]`. Die Mindestalter sind aufsteigend sortiert.
// Geliefert wird der Preis zum größten Mindestalter, das das Alter
// nicht übersteigt. Liegt das Alter unter allen Mindestaltern,
// ist der Eintritt frei.
//
// Die Preise aus `AdmissionPrice` ergeben sich also mit
// `minAges = []int{0, 6, 18}` und `prices = []int{0, 300, 500}`.
func AdmissionPriceFromLists(age int, minAges []int, prices []int) int {
	// Hinweis:
	// Gehen Sie die Mindestalter in einer Schleife durch.
	// - Merken Sie sich in einer Variablen den bisher gültigen Preis.
	//   Mit welchem Wert muss sie starten?
	// - Für den Preis zum Mindestalter benötigen Sie den Index.
	// - Wegen der Sortierung gilt am Ende der Preis des letzten
	//   Mindestalters, das das Alter nicht übersteigt.

	price := 0
	for i, minAge := range minAges {
		if age >= minAge {
			price = prices[i]
		}
	}
	return price
}
