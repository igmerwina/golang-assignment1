package function

import "pgd-sesi3/models"

func SeedData() []models.Absen {
	listPeserta := []models.Absen{
		{
			Nama:      "andi",
			Alamat:    "bogor",
			Pekerjaan: "petani",
			Alasan:    "sehat",
		},
		{
			Nama:      "kuji",
			Alamat:    "jakarta",
			Pekerjaan: "pejabat",
			Alasan:    "mantap",
		}, {
			Nama:      "pamo",
			Alamat:    "melawai",
			Pekerjaan: "gubernur",
			Alasan:    "baik",
		}, {
			Nama:      "juan",
			Alamat:    "kramat",
			Pekerjaan: "politisi",
			Alasan:    "baik",
		},
	}

	return listPeserta
}
