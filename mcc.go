package mcc

import (
	"errors"
	"github.com/rojack96/mcc/mnc"
)

// Reference for Mcc and Mnc link https://mcc-mnc.com/

type Mcc struct {
	Code        uint16
	Iso         string
	Country     string
	CountryCode uint
	Mnc         []mnc.Mnc
}

// GenerateHni The combination of MCC and MNC is called HNI (Home network identity) and is the combination of both in one string
// (e.g. MCC= 262 and MNC = 01 results in an HNI of 26201)
func GenerateHni() {
	return
}

func binarySearch(code uint16, mccSlice []Mcc) Mcc {
	indexLow := 0
	indexHigh := len(mccSlice) - 1

	if indexLow > indexHigh {
		return Mcc{}
	}

	indexMid := (indexLow + indexHigh) / 2

	if code == mccSlice[indexMid].Code {
		return mccSlice[indexMid]
	} else if code < mccSlice[indexMid].Code {
		newSlice := mccSlice[:indexMid-1]
		return binarySearch(code, newSlice)
	} else if code > mccSlice[indexMid].Code {
		newSlice := mccSlice[indexMid+1:]
		return binarySearch(code, newSlice)
	}

	return Mcc{}
}

func FindByCode(code uint16) (Mcc, error) {
	result := binarySearch(code, mcc)
	if result.Code == 0 {
		return Mcc{}, errors.New("mcc not found")
	}

	return result, nil
}

var mcc = []Mcc{
	{Code: 202, Iso: "GR", Country: "Greece", CountryCode: 30, Mnc: mnc.Greece},
	{Code: 204, Iso: "NL", Country: "Netherlands", CountryCode: 31, Mnc: mnc.Netherlands},
	{Code: 206, Iso: "BE", Country: "Belgium", CountryCode: 32, Mnc: mnc.Belgium},
	{Code: 208, Iso: "FR", Country: "France", CountryCode: 33, Mnc: mnc.France},
	{Code: 212, Iso: "MC", Country: "Monaco", CountryCode: 377, Mnc: mnc.Monaco},
	{Code: 213, Iso: "AD", Country: "Andorra", CountryCode: 376, Mnc: mnc.Andorra},
	{Code: 214, Iso: "ES", Country: "Spain", CountryCode: 34, Mnc: mnc.Spain},
	{Code: 216, Iso: "HU", Country: "Hungary", CountryCode: 36, Mnc: mnc.Hungary},
	{Code: 218, Iso: "BA", Country: "Bosnia and Herzegovina", CountryCode: 387, Mnc: mnc.BosniaAndHerzegovina},
	{Code: 219, Iso: "HR", Country: "Croatia", CountryCode: 385, Mnc: mnc.Croatia},
	{Code: 220, Iso: "RS", Country: "Serbia", CountryCode: 381, Mnc: mnc.Serbia},
	{Code: 221, Iso: "XK", Country: "Kosovo", CountryCode: 383, Mnc: mnc.Kosovo},
	{Code: 222, Iso: "IT", Country: "Italy", CountryCode: 39, Mnc: mnc.Italy},
	{Code: 225, Iso: "VA", Country: "Vatican", Mnc: mnc.Vatican},
	{Code: 226, Iso: "RO", Country: "Romania", CountryCode: 40, Mnc: mnc.Romania},
	{Code: 228, Iso: "CH", Country: "Switzerland", CountryCode: 41, Mnc: mnc.Switzerland},
	{Code: 230, Iso: "CZ", Country: "Czech Republic", CountryCode: 420, Mnc: mnc.CzechRepublic},
	{Code: 231, Iso: "SK", Country: "Slovakia", CountryCode: 421, Mnc: mnc.Slovakia},
	{Code: 232, Iso: "AT", Country: "Austria", CountryCode: 43},
	{Code: 234, Iso: "GB", Country: "United Kingdom", CountryCode: 44},
	{Code: 235, Iso: "GB", Country: "United Kingdom", CountryCode: 44},
	{Code: 238, Iso: "DK", Country: "Denmark", CountryCode: 45},
	{Code: 240, Iso: "SE", Country: "Sweden", CountryCode: 46},
	{Code: 242, Iso: "NO", Country: "Norway", CountryCode: 47},
	{Code: 244, Iso: "FI", Country: "Finland", CountryCode: 358},
	{Code: 246, Iso: "LT", Country: "Lithuania", CountryCode: 370},
	{Code: 247, Iso: "LV", Country: "Latvia", CountryCode: 371},
	{Code: 248, Iso: "EE", Country: "Estonia", CountryCode: 372},
	{Code: 250, Iso: "RU", Country: "Russia", CountryCode: 79},
	{Code: 255, Iso: "UA", Country: "Ukraine", CountryCode: 380},
	{Code: 257, Iso: "BY", Country: "Belarus", CountryCode: 375},
	{Code: 259, Iso: "MD", Country: "Moldova", CountryCode: 373},
	{Code: 260, Iso: "PL", Country: "Poland", CountryCode: 48},
	{Code: 262, Iso: "DE", Country: "Germany", CountryCode: 49},
	{Code: 266, Iso: "GI", Country: "Gibraltar", CountryCode: 350},
	{Code: 268, Iso: "PT", Country: "Portugal", CountryCode: 351},
	{Code: 270, Iso: "LU", Country: "Luxembourg", CountryCode: 352},
	{Code: 272, Iso: "IE", Country: "Ireland", CountryCode: 353},
	{Code: 274, Iso: "IS", Country: "Iceland", CountryCode: 354},
	{Code: 276, Iso: "AL", Country: "Albania", CountryCode: 355},
	{Code: 278, Iso: "MT", Country: "Malta", CountryCode: 356},
	{Code: 280, Iso: "CY", Country: "Cyprus", CountryCode: 357},
	{Code: 282, Iso: "GE", Country: "Georgia", CountryCode: 995},
	{Code: 283, Iso: "AM", Country: "Armenia", CountryCode: 374},
	{Code: 284, Iso: "BG", Country: "Bulgaria", CountryCode: 359},
	{Code: 286, Iso: "TR", Country: "Turkey", CountryCode: 90},
	{Code: 288, Iso: "FO", Country: "Faroe Island", CountryCode: 298},
	{Code: 289, Iso: "GE", Country: "Abkhazia", CountryCode: 7},
	{Code: 290, Iso: "GL", Country: "Greenland", CountryCode: 299},
	{Code: 292, Iso: "SM", Country: "San Marino", CountryCode: 378},
	{Code: 293, Iso: "SI", Country: "Slovenia", CountryCode: 386},
	{Code: 294, Iso: "MK", Country: "North Macedonia", CountryCode: 389},
	{Code: 295, Iso: "LI", Country: "Liechtenstein", CountryCode: 423},
	{Code: 297, Iso: "ME", Country: "Montenegro", CountryCode: 423},
	{Code: 302, Iso: "CA", Country: "Canada", CountryCode: 1},
	{Code: 308, Iso: "PM", Country: "Saint Pierre and Miquelon", CountryCode: 508},
	{Code: 310, Iso: "US", Country: "United States of America", CountryCode: 1},
	{Code: 311, Iso: "US", Country: "United States of America", CountryCode: 1},
	{Code: 312, Iso: "US", Country: "United States of America", CountryCode: 1},
	{Code: 313, Iso: "US", Country: "United States of America", CountryCode: 1},
	{Code: 314, Iso: "US", Country: "United States of America", CountryCode: 1},
	{Code: 315, Iso: "US", Country: "United States of America", CountryCode: 1},
	{Code: 316, Iso: "US", Country: "United States of America", CountryCode: 1},
	{Code: 330, Iso: "PR", Country: "Puerto Rico"},
	{Code: 332, Iso: "VI", Country: "United States Virgin Islands"},
	{Code: 334, Iso: "MX", Country: "Mexico", CountryCode: 52},
	{Code: 338, Iso: "JM", Country: "Jamaica", CountryCode: 1876},
	{Code: 340, Iso: "GF", Country: "Guadeloupe and Martinique and French Guiana", CountryCode: 594},
	{Code: 342, Iso: "BB", Country: "Barbados", CountryCode: 1246},
	{Code: 344, Iso: "AG", Country: "Antigua and Barbuda", CountryCode: 1268},
	{Code: 346, Iso: "KY", Country: "Cayman Islands", CountryCode: 1345},
	{Code: 348, Iso: "VG", Country: "British Virgin Islands", CountryCode: 284},
	{Code: 350, Iso: "BM", Country: "Bermuda", CountryCode: 1441},
	{Code: 352, Iso: "GD", Country: "Grenada", CountryCode: 1473},
	{Code: 354, Iso: "MS", Country: "Montserrat", CountryCode: 1666},
	{Code: 356, Iso: "KN", Country: "Saint Kitts and Nevis", CountryCode: 1869},
	{Code: 358, Iso: "LC", Country: "Saint Lucia", CountryCode: 1758},
	{Code: 360, Iso: "VC", Country: "Saint Vincent and the Grenadines", CountryCode: 1784},
	{Code: 362, Iso: "AN", Country: "Netherlands Antilles", CountryCode: 599},
	{Code: 363, Iso: "AW", Country: "Aruba", CountryCode: 297},
	{Code: 364, Iso: "BS", Country: "Bahamas", CountryCode: 1242},
	{Code: 365, Iso: "AI", Country: "Anguilla", CountryCode: 1264},
	{Code: 366, Iso: "DM", Country: "Dominica", CountryCode: 1767},
	{Code: 368, Iso: "CU", Country: "Cuba", CountryCode: 53},
	{Code: 370, Iso: "DO", Country: "Dominican Republic", CountryCode: 1809},
	{Code: 372, Iso: "HT", Country: "Haiti", CountryCode: 509},
	{Code: 374, Iso: "TT", Country: "Trinidad and Tobago", CountryCode: 1868},
	{Code: 376, Iso: "TC", Country: "Turks and Caicos Islands"},
	{Code: 400, Iso: "AZ", Country: "Azerbaijan", CountryCode: 994},
	{Code: 401, Iso: "KZ", Country: "Kazakhstan", CountryCode: 7},
	{Code: 402, Iso: "BT", Country: "Bhutan", CountryCode: 975},
	{Code: 404, Iso: "IN", Country: "India", CountryCode: 91},
	{Code: 405, Iso: "IN", Country: "India", CountryCode: 91},
	{Code: 410, Iso: "PK", Country: "Pakistan", CountryCode: 92},
	{Code: 412, Iso: "AF", Country: "Afghanistan", CountryCode: 93},
	{Code: 413, Iso: "LK", Country: "Sri Lanka", CountryCode: 94},
	{Code: 414, Iso: "MM", Country: "Myanmar", CountryCode: 95},
	{Code: 415, Iso: "LB", Country: "Lebanon", CountryCode: 961},
	{Code: 416, Iso: "JO", Country: "Jordan", CountryCode: 962},
	{Code: 417, Iso: "SY", Country: "Syria", CountryCode: 963},
	{Code: 418, Iso: "IQ", Country: "Iraq", CountryCode: 964},
	{Code: 419, Iso: "KW", Country: "Kuwait", CountryCode: 965},
	{Code: 420, Iso: "SA", Country: "Saudi Arabia", CountryCode: 966},
	{Code: 421, Iso: "YE", Country: "Yemen", CountryCode: 967},
	{Code: 422, Iso: "OM", Country: "Oman", CountryCode: 968},
	{Code: 423, Iso: "PS", Country: "Palestine", CountryCode: 970},
	{Code: 424, Iso: "AE", Country: "United Arab Emirates", CountryCode: 971},
	{Code: 425, Iso: "IL", Country: "Israel", CountryCode: 972},
	{Code: 426, Iso: "BH", Country: "Bahrain", CountryCode: 973},
	{Code: 427, Iso: "QA", Country: "Qatar", CountryCode: 974},
	{Code: 428, Iso: "MN", Country: "Mongolia", CountryCode: 976},
	{Code: 429, Iso: "NP", Country: "Nepal", CountryCode: 977},
	{Code: 430, Iso: "AE"},
	{Code: 431, Iso: "AE"},
	{Code: 432, Iso: "IR", Country: "Iran", CountryCode: 98},
	{Code: 434, Iso: "UZ", Country: "Uzbekistan", CountryCode: 998},
	{Code: 436, Iso: "TJ", Country: "Tajikistan", CountryCode: 992},
	{Code: 437, Iso: "KG", Country: "Kyrgyzstan", CountryCode: 996},
	{Code: 438, Iso: "TM", Country: "Turkmenistan", CountryCode: 99},
	{Code: 440, Iso: "JP", Country: "Japan", CountryCode: 81},
	{Code: 441, Iso: "JP", Country: "Japan", CountryCode: 81},
	{Code: 450, Iso: "KR", Country: "South Korea", CountryCode: 82},
	{Code: 452, Iso: "VN"},
	{Code: 454, Iso: "HK"},
	{Code: 455, Iso: "MO"},
	{Code: 456, Iso: "KH"},
	{Code: 457, Iso: "LA"},
	{Code: 460, Iso: "CN"},
	{Code: 466, Iso: "TW"},
	{Code: 467, Iso: "KP"},
	{Code: 470, Iso: "BD"},
	{Code: 472, Iso: "MV"},
	{Code: 502, Iso: "MY"},
	{Code: 505, Iso: "AU"},
	{Code: 510, Iso: "ID"},
	{Code: 514, Iso: "TL"},
	{Code: 515, Iso: "PH"},
	{Code: 520, Iso: "TH"},
	{Code: 525, Iso: "SG"},
	{Code: 528, Iso: "BN"},
	{Code: 530, Iso: "NZ"},
	{Code: 534, Iso: "MP"},
	{Code: 535, Iso: "GU"},
	{Code: 536, Iso: "NR"},
	{Code: 537, Iso: "PG"},
	{Code: 539, Iso: "TO"},
	{Code: 540, Iso: "SB"},
	{Code: 541, Iso: "VU"},
	{Code: 542, Iso: "FJ"},
	{Code: 543, Iso: "WF"},
	{Code: 544, Iso: "AS"},
	{Code: 545, Iso: "KI"},
	{Code: 546, Iso: "NC"},
	{Code: 547, Iso: "PF"},
	{Code: 548, Iso: "CK"},
	{Code: 549, Iso: "WS"},
	{Code: 550, Iso: "FM"},
	{Code: 551, Iso: "MH"},
	{Code: 552, Iso: "PW"},
	{Code: 602, Iso: "EG"},
	{Code: 603, Iso: "DZ"},
	{Code: 604, Iso: "MA"},
	{Code: 605, Iso: "TN"},
	{Code: 606, Iso: "LY"},
	{Code: 607, Iso: "GM"},
	{Code: 608, Iso: "SN"},
	{Code: 609, Iso: "MR"},
	{Code: 610, Iso: "ML"},
	{Code: 611, Iso: "GN"},
	{Code: 612, Iso: "CI"},
	{Code: 613, Iso: "BF"},
	{Code: 614, Iso: "NE"},
	{Code: 615, Iso: "TG"},
	{Code: 616, Iso: "BJ"},
	{Code: 617, Iso: "MU"},
	{Code: 618, Iso: "LR"},
	{Code: 619, Iso: "SL"},
	{Code: 620, Iso: "GH"},
	{Code: 621, Iso: "NG"},
	{Code: 622, Iso: "TD"},
	{Code: 623, Iso: "CF"},
	{Code: 624, Iso: "CM"},
	{Code: 625, Iso: "CV"},
	{Code: 626, Iso: "ST"},
	{Code: 627, Iso: "GQ"},
	{Code: 628, Iso: "GA"},
	{Code: 629, Iso: "CG"},
	{Code: 630, Iso: "CD"},
	{Code: 631, Iso: "AO"},
	{Code: 632, Iso: "GW"},
	{Code: 633, Iso: "SC"},
	{Code: 634, Iso: "SD"},
	{Code: 635, Iso: "RW"},
	{Code: 636, Iso: "ET"},
	{Code: 637, Iso: "SO"},
	{Code: 638, Iso: "DJ"},
	{Code: 639, Iso: "KE"},
	{Code: 640, Iso: "TZ"},
	{Code: 641, Iso: "UG"},
	{Code: 642, Iso: "BI"},
	{Code: 643, Iso: "MZ"},
	{Code: 645, Iso: "ZM"},
	{Code: 646, Iso: "MG"},
	{Code: 647, Iso: "RE"},
	{Code: 648, Iso: "ZW"},
	{Code: 649, Iso: "NA"},
	{Code: 650, Iso: "MW"},
	{Code: 651, Iso: "LS"},
	{Code: 652, Iso: "BW"},
	{Code: 653, Iso: "SZ"},
	{Code: 654, Iso: "KM"},
	{Code: 655, Iso: "ZA"},
	{Code: 657, Iso: "ER"},
	{Code: 702, Iso: "BZ"},
	{Code: 704, Iso: "GT"},
	{Code: 706, Iso: "SV"},
	{Code: 708, Iso: "HN"},
	{Code: 710, Iso: "NI"},
	{Code: 712, Iso: "CR"},
	{Code: 714, Iso: "PA"},
	{Code: 716, Iso: "PE"},
	{Code: 722, Iso: "AR"},
	{Code: 724, Iso: "BR"},
	{Code: 730, Iso: "CL"},
	{Code: 732, Iso: "CO"},
	{Code: 734, Iso: "VE"},
	{Code: 736, Iso: "BO"},
	{Code: 738, Iso: "GY"},
	{Code: 740, Iso: "EC"},
	{Code: 742, Iso: "GF"},
	{Code: 744, Iso: "PY"},
	{Code: 746, Iso: "SR"},
	{Code: 748, Iso: "UY"},
	{Code: 750, Iso: "FK"},
}

/*func main() {
	fmt.Println(findByCode(555))
}*/
