# Eintrittspreise

Ein Schwimmbad verlangt je nach Alter unterschiedliche Eintrittspreise:

| Alter          | Preis (Cent) |
| -------------- | ------------ |
| unter 6        | 0            |
| 6 bis 17       | 300          |
| ab 18          | 500          |

Die Übungsaufgaben bestimmen den Preis zu einem Alter,
zuerst mit fest eingebauter Preistabelle, dann mit wechselnden Tabellen:

- [`exercises/price_if.go`](exercises/price_if.go): Preistabelle fest im Code,
  mit `if`/`else if`/`else`.
- [`exercises/price_lists.go`](exercises/price_lists.go): Preistabelle als
  zwei Listen (Mindestalter und Preise).
- [`exercises/price_table.go`](exercises/price_table.go): Preistabelle als
  2D-Slice mit einer Zeile pro Altersgruppe.
- [`exercises/price_map.go`](exercises/price_map.go) (Bonus): Preistabelle als Map.

Die Aufgaben stehen in den Kommentaren der Funktionen, die erwarteten
Ergebnisse in den zugehörigen Tests. Prüfen mit `go test ./...` in diesem Verzeichnis.
