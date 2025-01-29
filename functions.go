package mcc

import (
	"errors"
	"github.com/rojack96/mcc/models"
	"strconv"
)

func (r *Mcc) FindByCode(mccCode string) (models.MccResult, error) {
	mccCodeString, err := strconv.Atoi(mccCode)
	if err != nil {
		return models.MccResult{}, err
	}

	result := r.binarySearchByCode(uint16(mccCodeString), mcc, 0, len(mcc)-1)
	if result.Code == "" {
		return models.MccResult{}, errors.New("mcc not found")
	}

	return result, nil
}

// HniListByCode The combination of MCC and MNC is called HNI (Home network identity) and is the combination of both in one string
// (e.g. MCC= 262 and MNC = 01 results in an HNI of 26201)
func (r *Mcc) HniListByCode(mccCode string) []string {
	result := make([]string, 0)
	mccTemp, err := r.FindByCode(mccCode)
	if err != nil {
		return nil
	}

	for _, m := range mccTemp.Mnc {
		result = append(result, mccTemp.Code+m.Code)
	}

	return result
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
