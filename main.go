package main

import (
	"fmt"
	"math/rand"
)

func main() {
	tablica := generujTabliceIntow(15)
	fmt.Println("tablica nieposortowana: ", tablica)
	QuickSort(tablica)
	fmt.Println("tablica posortowana: ", tablica)
}

func generujTabliceIntow(wielkoscTablicy int) (tablica []int) {
	tablica = make([]int, wielkoscTablicy)
	for i := 0; i < wielkoscTablicy; i++ {
		tablica[i] = rand.Intn(100)
	}
	return
}

func QuickSort(tablica []int) {
	if len(tablica) < 2 {
		return
	}
	wielkoscTablicy := len(tablica)
	lewy, prawy := 0, wielkoscTablicy-1

	piwot := rand.Intn(wielkoscTablicy)

	tablica[piwot], tablica[prawy] = tablica[prawy], tablica[piwot]

	for indeks := 0; indeks < prawy; indeks++ {
		if tablica[indeks] < tablica[prawy] {
			tablica[indeks], tablica[lewy] = tablica[lewy], tablica[indeks]
			lewy++
		}
	}
	tablica[prawy], tablica[lewy] = tablica[lewy], tablica[prawy]
	QuickSort(tablica[:lewy])
	QuickSort(tablica[lewy+1:])
}
