package countries

import "github.com/rojack96/mcc/models"

var Australia = models.Mcc{
	Code:        505,
	Iso:         "AU",
	Country:     "Australia",
	CountryCode: 61,
	Mnc: []models.Mnc{
		{Code: "01", Network: "Telstra Corp. Ltd."},
		{Code: "02", Network: "Singtel Optus"},
		{Code: "03", Network: "Vodafone"},
		{Code: "06", Network: "Hutchison 3G Australia Pty. Ltd."},
		{Code: "11", Network: "Telstra Corp. Ltd."},
		{Code: "12", Network: "Hutchison"},
		{Code: "19", Network: "Lycamobile Pty Ltd"},
		{Code: "71", Network: "Telstra Corp. Ltd."},
		{Code: "72", Network: "Telstra Corp. Ltd."},
		{Code: "90", Network: "Singtel Optus"},
		{Code: "999", Network: "Fix Line Australia"},
	},
}
