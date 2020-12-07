package pakiet

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

var nogiCiezkie float64

var blachyciezkie float64
var blachylekkie float64
var nogi float64

// funkcja Input sluzy do przyjmowania danych wejsciowych z klawiatury do programu
func Input() float64 { // czytnik danych wejsciowych
	reader := bufio.NewReader(os.Stdin)            // tworzenie czytnika
	scanner, _ := reader.ReadString('\n')          // uzycie czytnika do spacji
	scanner = strings.TrimSpace(scanner)           // kasowanie bialych znakow z wpisu
	liczba, err := strconv.ParseFloat(scanner, 64) // zamiana wczytanego stringu na float64
	if err != nil {
		log.Fatal("ERROR podano nieprawidlowa wartosc")
	}
	return liczba // zwrot liczby podanej do klawiatury
}

// funkcja sprawdza szerokosc frontu i na tej podstawie okresla dalsze dzialanie
func SelekcjaSzerokosci(glebokosc, szerokosc, ileRzedowCiezkich float64) {
	badanieFr := math.Mod(szerokosc, 2)
	if badanieFr == 0 { // szerokosc frontu typu 8 metrow
		FrontSTD(glebokosc, szerokosc, ileRzedowCiezkich)
	}
	if badanieFr == 0.5 { // szerokosc frontu typu 8.5 metra
		frontPlusPolowka(glebokosc, szerokosc, ileRzedowCiezkich)
	}
	if badanieFr == 1 { // szerokosc frontu typu 9 metrow ( osiem plus jeden )
		frontPlusJedynka(glebokosc, szerokosc, ileRzedowCiezkich)
	}
	if badanieFr == 1.5 { // szerokosc frontu typu 9.5 metra ( osiem plus jeden plus pol )
		frontPlusPoltora(glebokosc, szerokosc, ileRzedowCiezkich)
	}

}

//  funkcja podstawowa liczy bazowy wymiar sceny dla wymiaru prostokatnego
func Podstawowa(glebokosc, szerokosc, ileRzedowCiezkich float64) (iloscPodestowLekkich, blachylekkie, spinki, nogi float64) { // obliczanie ilosci blatow 2x1 dla sceny lekkiej lub lekkiej i ciezkiej
	ileRzedowLekkich := glebokosc - ileRzedowCiezkich                    // obliczanie ile realizator chcce uzyc rzedow lekkich
	iloscPodestowLekkich = math.Trunc(szerokosc/2) * ileRzedowLekkich    // obliczanie potrzebnej ilosci podestow lekkich z odrzuceniem polowek i metrowego w przypadku scen nieparzystych
	iloscPodestowCiezkich := math.Trunc(szerokosc/2) * ileRzedowCiezkich // obliczanie potrzebnej ilosci podestow ciezkich
	// jesli sa podesty ciezkie to je wypisz
	if iloscPodestowCiezkich > 0 {
		fmt.Printf("Potrzeba %0.0f podestow ciezkich o wymiarze 2x1.\n", iloscPodestowCiezkich)

	}
	blachylekkie = (math.Trunc(szerokosc/2) - 1) * (ileRzedowLekkich - 1) // obliczanie ile potrzeba blach, rur i stop

	if iloscPodestowCiezkich > 0 { // funkcja liczaca ewentualnosc blach dla ciezkich
		blachyciezkie = ((math.Trunc(szerokosc / 2)) - 1) * (ileRzedowCiezkich - 1)
	} else {
		blachyciezkie = 0
	}
	//ilosc spinek
	spinki = ((iloscPodestowCiezkich + iloscPodestowLekkich) * 2)
	// ilosc nog
	nogiLekkie := ((2 * szerokosc) + (4 * ileRzedowLekkich) - 4)

	if iloscPodestowCiezkich > 0 {
		nogiCiezkie = ((2 * szerokosc) + (4 * ileRzedowCiezkich) - 4.0)
	} else {
		nogiCiezkie = 0
	}

	nogi = nogiCiezkie + nogiLekkie

	return iloscPodestowLekkich, spinki, nogi, blachylekkie

}

func FrontSTD(glebokosc, szerokosc, ileRzedowCiezkich float64) {
	//Podstawowa(glebokosc, szerokosc, ileRzedowCiezkich)
	ileBlatow, ileSpinek, ileNog, ileBlaszek := Podstawowa(glebokosc, szerokosc, ileRzedowCiezkich)
	fmt.Printf("Potrzeba %0.0f podestow lekkich\n, Potrzeba %0.0f nog\n", ileBlatow, ileNog)
	fmt.Printf("Potrzeba %0.0f zestawow rura + blacha + stopa,\n Potrzeba %0.0f spinek do podestów.", ileBlaszek, ileSpinek)
}
func frontPlusJedynka(glebokosc, szerokosc, ileRzedowCiezkich float64) {
	ileBlatow, ileSpinek, ileNog, ileBlaszek := Podstawowa(glebokosc, szerokosc, ileRzedowCiezkich)

	ileMetrowegoPodestu := math.Mod(glebokosc, 2)
	if ileMetrowegoPodestu == 0 { // opcja jesli same podesty dwumetrowe na boku dokladasz
		podestyDwuMetrowe := glebokosc / 2
		nozki := 4 * podestyDwuMetrowe
		klamry := 2 * podestyDwuMetrowe
		fmt.Printf("Potrzeba %0.0f podestow lekkich\n, Potrzeba %0.0f nog\n", ileBlatow+podestyDwuMetrowe, ileNog+nozki)
		fmt.Printf("Potrzeba %0.0f zestawow rura + blacha + stopa,\n Potrzeba %0.0f spinek do podestów.", ileBlaszek, ileSpinek+klamry)

	} else { // opcja jesli jeszcze jeden podest metrowy wpada
		podestyMetrowePrzedZmiana := glebokosc / 2
		podestyDwuMetrowe := math.Trunc(podestyMetrowePrzedZmiana)
		nozki := (4 * podestyDwuMetrowe) + 4
		ileblacikow := math.Trunc(glebokosc / 2)
		klamerki := (2 * podestyDwuMetrowe) + 2
		fmt.Println("Potrzeby jest jeden podest 1x1 metra.")
		fmt.Printf("Potrzeba %0.0f podestow lekkich,\nPotrzeba %0.0f nog\n", ileBlatow+ileblacikow, ileNog+nozki)
		fmt.Printf("Potrzeba %0.0f zestawow rura + blacha + stopa,\n Potrzeba %0.0f spinek do podestów.", ileBlaszek, ileSpinek+klamerki)

	}
}

func frontPlusPolowka(glebokosc float64, szerokosc float64, ileRzedowCiezkich float64) {
	ileBlatow, ileSpinek, ileNog, ileBlaszek := Podstawowa(glebokosc, szerokosc, ileRzedowCiezkich)
	ilePolmetrowegoPodestu := math.Mod(glebokosc, 2)
	if ilePolmetrowegoPodestu == 0 {
		podestyPolmetrowe := glebokosc / 2
		nozki := 4 * podestyPolmetrowe
		klamry := 2 * podestyPolmetrowe
		fmt.Printf("Potrzeba %0.0f podestów 2x0.5 metra ", podestyPolmetrowe)
		fmt.Printf("Potrzeba %0.0f podestow lekkich\n, Potrzeba %0.0f nog\n", ileBlatow, ileNog+nozki)
		fmt.Printf("Potrzeba %0.0f zestawow rura + blacha + stopa,\n Potrzeba %0.0f spinek do podestów.", ileBlaszek, ileSpinek+klamry)

	} else {
		podestyPolMetrowePrzedZmiana := glebokosc / 2
		podestyPolMetrowe := math.Trunc(podestyPolMetrowePrzedZmiana)
		ileBlacikow := math.Trunc(glebokosc / 2) // podest metrowy zaklamuje dzielenie glebokosci na dwa, stad ta metoda
		nozki := 4*podestyPolMetrowe + 4
		klamry := 2*podestyPolMetrowe + 2
		fmt.Printf("Potrzeba JEDEN podest 1x1 metra oraz %0.0f podestów 2x0.5 metra ", ileBlacikow)
		fmt.Printf("Potrzeba %0.0f podestow lekkich\n, Potrzeba %0.0f nog\n", ileBlatow, ileNog+nozki)
		fmt.Printf("Potrzeba %0.0f zestawow rura + blacha + stopa,\n Potrzeba %0.0f spinek do podestów.", ileBlaszek, ileSpinek+klamry)

	}
}

func frontPlusPoltora(glebokosc float64, szerokosc float64, ileRzedowCiezkich float64) {
	ileBlatow, ileSpinek, ileNog, ileBlaszek := Podstawowa(glebokosc, szerokosc, ileRzedowCiezkich)
	ilePoltorametrowegoPodestu := math.Mod(glebokosc, 2)
	if ilePoltorametrowegoPodestu == 0 {
		podestyPolmetrowe := glebokosc / 2
		fmt.Printf("Potrzeba %0.0f podestow lekkich 2x0.5 metra\n", podestyPolmetrowe)
		fmt.Printf("Potrzeba %0.0f podestow lekkich,\nPotrzeba %0.0f nog\n", ileBlatow+podestyPolmetrowe, ileNog+8*podestyPolmetrowe)
		fmt.Printf("Potrzeba %0.0f zestawow rura + blacha + stopa,\nPotrzeba %0.0f spinek do podestów.", ileBlaszek, ileSpinek+4*podestyPolmetrowe)
	} else {
		glebokoscPodzielona := math.Trunc(glebokosc / 2)
		podestyPolMetrowePrzedZmiana := glebokosc / 2
		podestyPolMetrowe := math.Trunc(podestyPolMetrowePrzedZmiana)
		fmt.Println("Potrzeba jeden podest o wymiarze 1x1 metra oraz jeden podest 1x.05 metra. ")
		fmt.Printf("Potrzeba %0.0f podestow lekkich 2x0.5 metra\n", glebokoscPodzielona)
		fmt.Printf("Potrzeba %0.0f podestow lekkich,\nPotrzeba %0.0f nog\n", ileBlatow+glebokoscPodzielona, ileNog+(8*podestyPolMetrowe)+4)
		fmt.Printf("Potrzeba %0.0f zestawow rura + blacha + stopa,\nPotrzeba %0.0f spinek do podestów.", ileBlaszek, ileSpinek+4*podestyPolMetrowe+(2*podestyPolMetrowe+2))

	}
}
