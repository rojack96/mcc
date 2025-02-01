package europe

import (
	"github.com/rojack96/mcc/models"
)

var Greece = models.Mcc{
	Code:        202,
	Iso:         "GR",
	Country:     "Greece",
	CountryCode: 30,
	Mnc: []models.Mnc{
		{Code: "299", Network: "AMD Telecom"},
		{Code: "299", Network: "Apifon"},
		{Code: "15", Network: "BWS"},
		{Code: "2", Network: "Cosmote"},
		{Code: "1", Network: "Cosmote"},
		{Code: "299", Network: "Failed Calls"},
		{Code: "999", Network: "Fix Line"},
		{Code: "16", Network: "Inter Telecom"},
		{Code: "299", Network: "Interconnect"},
		{Code: "299", Network: "M-STAT"},
		{Code: "299", Network: "Nova"},
		{Code: "4", Network: "OSE"},
		{Code: "3", Network: "OTE"},
		{Code: "299", Network: "OTEGLOBE"},
		{Code: "299", Network: "Premium Net"},
		{Code: "5", Network: "Vodafone"},
		{Code: "10", Network: "Wind"},
		{Code: "9", Network: "Wind"},
		{Code: "12", Network: "Yuboto"},
	}}
