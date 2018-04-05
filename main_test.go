package main

import (
	"testing"
)

func TestCzyQuicksortDobrzeSortuje(t *testing.T) {
	tablica := generujTabliceIntow(15)
	QuickSort(tablica)
	tablicaDobrzePosortowana := tablicaDobrzePosortowanaRosnaco(tablica)
	if !tablicaDobrzePosortowana {
		t.Error("Tablica zle posortowana")
	}
}

func tablicaDobrzePosortowanaRosnaco(tablica []int) bool {
	if len(tablica) < 2 {
		return true
	}
	for indeks := range tablica {
		if indeks > 0 {
			if tablica[indeks-1] > tablica[indeks] {
				return false
			}
		}
	}
	return true
}
