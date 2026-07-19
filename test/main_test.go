package main

import (
	"reflect"
	"testing"

	"github.com/rojack96/mcc"
	"github.com/rojack96/mcc/models"
)

func TestMccInfoAndFieldHelpers(t *testing.T) {
	reader := mcc.NewMccReader()

	info, err := reader.MccInfo("222")
	if err != nil {
		t.Fatalf("MccInfo returned error: %v", err)
	}

	if info.Code != "222" {
		t.Fatalf("Code = %q, want %q", info.Code, "222")
	}
	if info.Iso != "IT" {
		t.Fatalf("Iso = %q, want %q", info.Iso, "IT")
	}
	if info.Country != "Italy" {
		t.Fatalf("Country = %q, want %q", info.Country, "Italy")
	}
	if info.CountryCode != 39 {
		t.Fatalf("CountryCode = %d, want %d", info.CountryCode, 39)
	}
	if len(info.Mnc) == 0 {
		t.Fatal("MccInfo returned no MNC entries for Italy")
	}

	iso, err := reader.Iso("222")
	if err != nil {
		t.Fatalf("Iso returned error: %v", err)
	}
	if iso != info.Iso {
		t.Fatalf("Iso = %q, want %q", iso, info.Iso)
	}

	country, err := reader.Country("222")
	if err != nil {
		t.Fatalf("Country returned error: %v", err)
	}
	if country != info.Country {
		t.Fatalf("Country = %q, want %q", country, info.Country)
	}

	countryCode, err := reader.CountryCode("222")
	if err != nil {
		t.Fatalf("CountryCode returned error: %v", err)
	}
	if countryCode != info.CountryCode {
		t.Fatalf("CountryCode = %d, want %d", countryCode, info.CountryCode)
	}
}

func TestMccLookupErrors(t *testing.T) {
	reader := mcc.NewMccReader()

	tests := []struct {
		name    string
		mccCode string
		wantErr string
	}{
		{name: "not numeric", mccCode: "abc", wantErr: "MCC not valid"},
		{name: "unknown", mccCode: "999", wantErr: "MCC not found"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := reader.MccInfo(tt.mccCode)
			if err == nil {
				t.Fatal("MccInfo returned nil error")
			}
			if err.Error() != tt.wantErr {
				t.Fatalf("error = %q, want %q", err.Error(), tt.wantErr)
			}
		})
	}
}

func TestMncList(t *testing.T) {
	reader := mcc.NewMccReader()

	list, err := reader.MncList("262")
	if err != nil {
		t.Fatalf("MncList returned error: %v", err)
	}
	if !containsMnc(list, "01", "T-mobile") {
		t.Fatal("MncList for Germany does not contain T-mobile with code 01")
	}
	if !containsMnc(list, "02", "Vodafone") {
		t.Fatal("MncList for Germany does not contain Vodafone with code 02")
	}

	_, err = reader.MncList("330")
	if err == nil {
		t.Fatal("MncList returned nil error for MCC without MNC entries")
	}
	if err.Error() != "MNC not found" {
		t.Fatalf("error = %q, want %q", err.Error(), "MNC not found")
	}
}

func TestMncMapGroupedByMncCode(t *testing.T) {
	reader := mcc.NewMccReader()

	byCode, err := reader.MncMap("262", mcc.MncCode)
	if err != nil {
		t.Fatalf("MncMap returned error: %v", err)
	}

	if got := byCode["01"]; !reflect.DeepEqual(got, []string{"T-mobile"}) {
		t.Fatalf("byCode[01] = %#v, want %#v", got, []string{"T-mobile"})
	}
	if got := byCode["299"]; !containsAll(got, "Argon Networks", "Failed Calls", "Tismi") {
		t.Fatalf("byCode[299] = %#v, want all expected networks", got)
	}
}

func TestMncMapGroupedByNetwork(t *testing.T) {
	reader := mcc.NewMccReader()

	byNetwork, err := reader.MncMap("262", mcc.Network)
	if err != nil {
		t.Fatalf("MncMap returned error: %v", err)
	}

	if got := byNetwork["T-mobile"]; !containsAll(got, "01", "06", "13", "78") {
		t.Fatalf("byNetwork[T-mobile] = %#v, want all expected MNC codes", got)
	}
	if got := byNetwork["Vodafone"]; !containsAll(got, "02", "04", "09") {
		t.Fatalf("byNetwork[Vodafone] = %#v, want all expected MNC codes", got)
	}
}

func TestHniListByCode(t *testing.T) {
	reader := mcc.NewMccReader()

	hniList := reader.HniListByCode("412")
	if !containsAll(hniList, "41201", "41220", "41280") {
		t.Fatalf("HniListByCode = %#v, want expected Afghanistan HNI values", hniList)
	}

	if got := reader.HniListByCode("abc"); got != nil {
		t.Fatalf("HniListByCode for invalid MCC = %#v, want nil", got)
	}
}

func TestImsiList(t *testing.T) {
	reader := mcc.NewMccReader()

	imsiList := reader.ImsiList("412", "123456789")
	if !containsAll(imsiList, "41201123456789", "41220123456789", "41280123456789") {
		t.Fatalf("ImsiList = %#v, want expected Afghanistan IMSI values", imsiList)
	}

	if got := reader.ImsiList("999", "123456789"); got != nil {
		t.Fatalf("ImsiList for unknown MCC = %#v, want nil", got)
	}
}

func containsMnc(list []models.Mnc, code, network string) bool {
	for _, item := range list {
		if item.Code == code && item.Network == network {
			return true
		}
	}
	return false
}

func containsAll(values []string, want ...string) bool {
	for _, expected := range want {
		found := false
		for _, value := range values {
			if value == expected {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}
