package exercises

// AdmissionPriceFromTable erwartet ein Alter in Jahren und eine
// Preistabelle, deren Zeilen jeweils aus einem Mindestalter und
// einem Preis bestehen, z.B. `{{0, 0}, {6, 300}, {18, 500}}`.
// Die Zeilen sind aufsteigend nach dem Mindestalter sortiert.
// Ansonsten gilt dasselbe wie für `AdmissionPriceFromLists`.
func AdmissionPriceFromTable(age int, table [][]int) int {
	// Hinweis:
	// Gehen Sie vor wie in `AdmissionPriceFromLists`, aber mit einer
	// Schleife über die Zeilen der Tabelle.
	// - Jede Zeile ist selbst eine Liste mit zwei Einträgen:
	//   `row[0]` ist das Mindestalter, `row[1]` der Preis.

	price := 0
	for _, row := range table {
		if age >= row[0] {
			price = row[1]
		}
	}
	return price
}
