package main

import (
	"fmt"
	"strings"
)

type Biodata struct {
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}

func main() {

	filterName := "Thomas"
	result := filterBiodataByName(generateBiodata(), filterName)
	fmt.Println(result)

}

func generateBiodata() []Biodata {

	listBiodata := []Biodata{
		{
			nama:      "Thomas",
			alamat:    "Jalan Satu",
			pekerjaan: "Kerja Satu",
			alasan:    "Alasan Satu",
		},
		{
			nama:      "Rio",
			alamat:    "Jalan Dua",
			pekerjaan: "Kerja Dua",
			alasan:    "Alasan Dua",
		},
		{
			nama:      "Janiero",
			alamat:    "Jalan Tiga",
			pekerjaan: "Kerja Tiga",
			alasan:    "Alasan Tiga",
		},
		{
			nama:      "Thomas",
			alamat:    "Jalan Empat",
			pekerjaan: "Kerja Empat",
			alasan:    "Alasan Empat",
		},
	}

	return listBiodata

}

func filterBiodataByName(list []Biodata, nameFilter string) []Biodata {
	var result []Biodata
	for _, v := range list {
		if strings.EqualFold(nameFilter, v.nama) {
			result = append(result, v)
		}
	}
	return result
}
