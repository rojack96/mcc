package countries

import "github.com/rojack96/mcc/models"

var Botswana = models.Mcc{
	Code:        652,
	Iso:         "BW",
	Country:     "Botswana",
	CountryCode: 267,
	Mnc: []models.Mnc{
		{Code: "01", Network: "Mascom Wireless (Pty) Ltd."},
		{Code: "02", Network: "Orange"},
		{Code: "04", Network: "beMOBILE"},
	},
}
