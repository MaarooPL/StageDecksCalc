/*
Program do przeliczania zapotrzebowania na materiał potrzebny do budowy sceny.
Program przelicza ile podestow lekkich i ciezkich potrzeba do zbudowania sceny, ile nog itp.
Program przyjmuje, ze najpierw budowane są podesty lekkie a potem ciezkie.
*/
package main

import (
	"Poniat/PoniatPakiet"
	"fmt"
)

func main() {
	fmt.Println("Podaj szerokość sceny:")
	szerokosc := pakiet.Input()
	fmt.Println("Podaj głębokość sceny:")
	glebokosc := pakiet.Input()
	fmt.Println("Podaj ile rzedow podestow ciezkich chcesz zastosowac:")
	ileRzedowCiezkich := pakiet.Input()
	fmt.Printf("Podany szerokosc sceny to %0.1f metrów a głębokość to %0.1f metrów, uzyto %0.0f rzedow podestow ciezkich.\n", szerokosc, glebokosc, ileRzedowCiezkich)
	fmt.Printf("Program zaklada, ze od przodu idą podesty lekkie a z tylu idą podesty ciężkie. \n\n\n")

	fmt.Printf("Podano %0.0f rzędow podestów ciezkich.\n", ileRzedowCiezkich)
	pakiet.SelekcjaSzerokosci(glebokosc, szerokosc, ileRzedowCiezkich) // obliczanie pelnego zapotrzeboewania

}
