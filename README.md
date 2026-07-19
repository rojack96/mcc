# mcc

[![Go Reference](https://pkg.go.dev/badge/github.com/rojack96/mcc.svg)](https://pkg.go.dev/github.com/rojack96/mcc)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

`mcc` is a small Go library for reading Mobile Country Code (MCC) and Mobile
Network Code (MNC) information.

It can be used to look up country metadata, list mobile network operators,
group MNCs by code or network name, and generate HNI/IMSI prefixes.

## Installation

```bash
go get github.com/rojack96/mcc
```

## Quick Start

```go
package main

import (
	"fmt"
	"log"

	"github.com/rojack96/mcc"
)

func main() {
	reader := mcc.NewMccReader()

	info, err := reader.MccInfo("222")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(info.Country)     // Italy
	fmt.Println(info.Iso)         // IT
	fmt.Println(info.CountryCode) // 39
	fmt.Println(info.Mnc)         // [{299 1Mobile} {299 A-Tono} ...]
}
```

## Usage

### Get MCC Information

```go
reader := mcc.NewMccReader()

info, err := reader.MccInfo("262")
if err != nil {
	// handle invalid or unknown MCC
}

fmt.Println(info.Code)        // 262
fmt.Println(info.Iso)         // DE
fmt.Println(info.Country)     // Germany
fmt.Println(info.CountryCode) // 49
```

### Read Individual Fields

```go
iso, err := reader.Iso("222")
country, err := reader.Country("222")
countryCode, err := reader.CountryCode("222")

fmt.Println(iso)         // IT
fmt.Println(country)     // Italy
fmt.Println(countryCode) // 39
```

### List MNCs for an MCC

```go
mncList, err := reader.MncList("262")
if err != nil {
	// returns an error when the MCC is invalid, unknown, or has no MNC entries
}

for _, mnc := range mncList {
	fmt.Printf("MNC: %s, Network: %s\n", mnc.Code, mnc.Network)
}
```

### Group MNCs by Code

```go
byCode, err := reader.MncMap("262", mcc.MncCode)
if err != nil {
	// handle error
}

fmt.Println(byCode["01"]) // [T-mobile]
```

### Group MNCs by Network

```go
byNetwork, err := reader.MncMap("262", mcc.Network)
if err != nil {
	// handle error
}

fmt.Println(byNetwork["Vodafone"]) // [02 09 04]
```

### Generate HNI Values

The Home Network Identity (HNI) is the MCC and MNC combined into one string.

```go
hniList := reader.HniListByCode("412")

fmt.Println(hniList) // [41201 41220 41230 ...]
```

### Generate IMSI Prefixes

An IMSI value is composed by combining an HNI with an MSIN.

```go
imsiList := reader.ImsiList("412", "123456789")

fmt.Println(imsiList) // [41201123456789 41220123456789 ...]
```

## Public API

Create a reader with `NewMccReader`:

```go
reader := mcc.NewMccReader()
```

Available methods on `mcc.Mcc`:

```go
func NewMccReader() Mcc

func (r *Mcc) MccInfo(mccCode string) (models.MccResult, error)
func (r *Mcc) Iso(mccCode string) (string, error)
func (r *Mcc) Country(mccCode string) (string, error)
func (r *Mcc) CountryCode(mccCode string) (uint, error)
func (r *Mcc) MncList(mccCode string) ([]models.Mnc, error)
func (r *Mcc) MncMap(mccCode string, groupBy GroupBy) (map[string][]string, error)
func (r *Mcc) HniListByCode(mccCode string) []string
func (r *Mcc) ImsiList(mccCode, msin string) []string
```

Supported grouping values for `MncMap`:

| Value | Description |
| --- | --- |
| `mcc.MncCode` | Groups by MNC code and returns network names. |
| `mcc.Network` | Groups by network name and returns MNC codes. |

## Data Models

```go
type MccResult struct {
	Code        string
	Iso         string
	Country     string
	CountryCode uint
	Mnc         []Mnc
}

type Mnc struct {
	Code    string
	Network string
}
```

## Errors

The library returns errors for invalid or missing MCC data:

| Case | Error |
| --- | --- |
| MCC is not numeric | `MCC not valid` |
| MCC is not present in the database | `MCC not found` |
| MCC has no MNC entries | `MNC not found` |

`HniListByCode` and `ImsiList` return `nil` when the MCC cannot be found or no
HNI values can be built.

## Data Source

The MCC/MNC data used by this library is based on the public database available
at [MCC-MNC.com](https://mcc-mnc.com/#database).

The generated country data also references the site API endpoint exposed by the
web application:

```text
https://mcc-mnc.com/api/v1/mcc-mnc.php
```

Please refer to [MCC-MNC.com](https://mcc-mnc.com/) for the original database,
updates, and any data usage terms that may apply.

## Testing

Run the test suite with:

```bash
go test ./...
```

## License

This project is licensed under the [MIT License](LICENSE).

The software license applies to the code in this repository. The MCC/MNC facts
are sourced from [MCC-MNC.com](https://mcc-mnc.com/#database); please review
their website for data attribution and usage requirements.
