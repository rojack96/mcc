package countries

import "github.com/rojack96/mcc/models"

var Malaysia = models.Mcc{
	Code:        502,
	Iso:         "MY",
	Country:     "Malaysia",
	CountryCode: 60,
	Mnc: []models.Mnc{
		{Code: "10", Network: "Digi Telecommunications"},
		{Code: "12", Network: "Maxis"},
		{Code: "13", Network: "CelCom"},
		{Code: "150", Network: "TuneTalk"},
		{Code: "152", Network: "YES"},
		{Code: "153", Network: "TM/Unify"},
		{Code: "16", Network: "Digi Telecommunications"},
		{Code: "17", Network: "Maxis"},
		{Code: "18", Network: "U Mobile"},
		{Code: "19", Network: "CelCom"},
		{Code: "195", Network: "XOX Com Sdn Bhd"},
		{Code: "198", Network: "CelCom"},
		{Code: "999", Network: "Fix Line"},
	},
}
