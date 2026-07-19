package mcc

import (
	"errors"
	"github.com/rojack96/mcc/models"
	"strconv"
)

type GroupBy string

const (
	MncCode GroupBy = "mncCode"
	Network GroupBy = "network"
)

// MccInfo returns the full MCC record for the given Mobile Country Code.
func (r *Mcc) MccInfo(mccCode string) (models.MccResult, error) {
	result, err := r.findByCode(mccCode)
	if err != nil {
		return models.MccResult{}, err
	}

	return result, nil
}

// Iso returns the ISO country code associated with the given MCC.
func (r *Mcc) Iso(mccCode string) (string, error) {
	result, err := r.findByCode(mccCode)
	if err != nil {
		return "", err
	}

	return result.Iso, nil
}

// Country returns the country name associated with the given MCC.
func (r *Mcc) Country(mccCode string) (string, error) {
	result, err := r.findByCode(mccCode)
	if err != nil {
		return "", err
	}

	return result.Country, nil
}

// CountryCode returns the international calling code associated with the given MCC.
func (r *Mcc) CountryCode(mccCode string) (uint, error) {
	result, err := r.findByCode(mccCode)
	if err != nil {
		return 0, err
	}

	return result.CountryCode, nil
}

// MncList - return a list of MNC of relative MCC
func (r *Mcc) MncList(mccCode string) ([]models.Mnc, error) {
	result, err := r.findByCode(mccCode)
	if err != nil {
		return nil, err
	}

	if len(result.Mnc) == 0 {
		return nil, errors.New("MNC not found")
	}

	return result.Mnc, nil
}

// MncMap returns MNC entries grouped either by MNC code or by network name.
func (r *Mcc) MncMap(mccCode string, groupBy GroupBy) (map[string][]string, error) {
	result := make(map[string][]string)

	list, err := r.MncList(mccCode)
	if err != nil {
		return nil, err
	}

	for _, mnc := range list {
		var (
			key, value = mnc.Code, mnc.Network
		)

		switch groupBy {
		case MncCode:
			key = mnc.Code
			value = mnc.Network
		case Network:
			key = mnc.Network
			value = mnc.Code
		}

		result[key] = append(result[key], value)
	}

	return result, nil
}

// HniListByCode - The combination of MCC and MNC is called HNI (Home network identity) and is the combination of both in one string
// (e.g. MCC= 262 and MNC = 01 results in an HNI of 26201)
func (r *Mcc) HniListByCode(mccCode string) []string {
	result := make([]string, 0)
	mccTemp, err := r.findByCode(mccCode)
	if err != nil {
		return nil
	}

	for _, m := range mccTemp.Mnc {
		result = append(result, mccTemp.Code+m.Code)
	}

	return result
}

// ImsiList - If you combine the HNI with the MSIN (Mobile Subscriber Identification Number)
// the result is the so called IMSI (integrated mobile subscriber identify).
func (r *Mcc) ImsiList(mccCode, msin string) []string {
	result := make([]string, 0)
	hniList := r.HniListByCode(mccCode)

	if len(hniList) == 0 {
		return nil
	}

	for _, h := range hniList {
		result = append(result, h+msin)
	}

	return result
}

func (r *Mcc) findByCode(mccCode string) (models.MccResult, error) {
	mccCodeString, err := strconv.Atoi(mccCode)
	if err != nil {
		return models.MccResult{}, errors.New("MCC not valid")
	}

	result := r.binarySearchByCode(uint16(mccCodeString), mcc, 0, len(mcc)-1)
	if result.Code == "" {
		return models.MccResult{}, errors.New("MCC not found")
	}

	return result, nil
}

func (r *Mcc) binarySearchByCode(target uint16, mccSlice []models.Mcc, lowIdx, highIdx int) models.MccResult {
	if lowIdx > highIdx {
		return models.MccResult{}
	}

	midIdx := (lowIdx + highIdx) / 2

	if target == mccSlice[midIdx].Code {
		res := mccSlice[midIdx]
		return models.MccResult{
			Code:        strconv.Itoa(int(res.Code)),
			Iso:         res.Iso,
			Country:     res.Country,
			CountryCode: res.CountryCode,
			Mnc:         res.Mnc,
		}
	} else if target < mccSlice[midIdx].Code {
		return r.binarySearchByCode(target, mccSlice, lowIdx, midIdx-1)
	} else if target > mccSlice[midIdx].Code {
		return r.binarySearchByCode(target, mccSlice, midIdx+1, highIdx)
	}

	return models.MccResult{}
}
