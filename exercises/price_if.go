package exercises

// AdmissionPrice erwartet ein Alter in Jahren und liefert den
// Eintrittspreis in Cent: Unter `6` Jahren ist der Eintritt frei,
// bis einschließlich `17` Jahren kostet er `300`, ab `18` Jahren `500`.
func AdmissionPrice(age int) int {
	// Hinweis:
	// Prüfen Sie mit `if`/`else if`/`else` nacheinander, ob das Alter
	// unter `6` liegt, dann ob es unter `18` liegt.
	// - Die zweite Bedingung wird nur geprüft, wenn die erste nicht
	//   zutrifft. Sie braucht deshalb keine untere Grenze.
	// - Achten Sie auf die Reihenfolge der Bedingungen.

	if age < 6 {
		return 0
	} else if age < 18 {
		return 300
	} else {
		return 500
	}
	// Anmerkung:
	// Reihenfolge der Bedingungen:
	// Stünde `age < 18` vor `age < 6`, träfe sie auch für `5` zu.
	// Der freie Eintritt würde dann nie erreicht.
}
