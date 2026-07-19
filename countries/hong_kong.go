package countries

import "github.com/rojack96/mcc/models"

var HongKong = models.Mcc{
	Code:        454,
	Iso:         "HK",
	Country:     "Hongkong",
	CountryCode: 852,
	Mnc: []models.Mnc{
		{Code: "00", Network: "CSL Ltd."},
		{Code: "02", Network: "CSL Ltd."},
		{Code: "03", Network: "H3G/Hutchinson"},
		{Code: "04", Network: "H3G/Hutchinson"},
		{Code: "05", Network: "H3G/Hutchinson"},
		{Code: "06", Network: "Vodafone/SmarTone"},
		{Code: "08", Network: "Trident Telecom Ventures Ltd."},
		{Code: "09", Network: "China Motion"},
		{Code: "10", Network: "CSL/New World PCS"},
		{Code: "12", Network: "China Mobile/Peoples"},
		{Code: "13", Network: "China Mobile/Peoples"},
		{Code: "14", Network: "H3G/Hutchinson"},
		{Code: "15", Network: "Vodafone/SmarTone"},
		{Code: "16", Network: "HKT/PCCW"},
		{Code: "17", Network: "Vodafone/SmarTone"},
		{Code: "18", Network: "CSL Ltd."},
		{Code: "19", Network: "HKT/PCCW"},
		{Code: "28", Network: "China Mobile/Peoples"},
		{Code: "29", Network: "HKT/PCCW"},
		{Code: "31", Network: "CTExcel"},
	},
}
