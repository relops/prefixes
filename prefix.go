package prefixes

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"github.com/relops/csvb"
	"io"
	"speter.net/go/exp/math/dec/inf"
	"strings"
)

type Country struct {
	Name      string
	Prefix    string
	Relevance *inf.Dec
}

type areaCode struct {
	name        string
	code        string
	countryName string
}

var (
	prefixes = make(map[string]Country)
	nanp     = make(map[string]areaCode)
)

func init() {
	initWorldData()
	initNanpData()
}

func initNanpData() {
	data, _ := data_canada_csv()
	readNanp("Canada", data)
	data, _ = data_usa_csv()
	readNanp("United States", data)
	data, _ = data_caribbean_csv()
	readNanp("Carribean", data)
	data, _ = data_non_geo_csv()
	readNanp("Non Geo", data)
}

func readNanp(country string, data []byte) {
	r := bytes.NewReader(data)

	csv := csv.NewReader(r)
	csv.FieldsPerRecord = -1

	for {
		row, err := csv.Read()
		if err == io.EOF {
			break
		}

		for i := 1; i < len(row); i++ {
			area := areaCode{
				name:        row[0],
				code:        row[i],
				countryName: country,
			}
			nanp[row[i]] = area
		}

	}
}

func initWorldData() {
	data, _ := data_countries_csv()
	r := bytes.NewReader(data)

	opts := &csvb.Options{Separator: ';'}

	strategy := make(map[string]string)
	strategy["name"] = "Name"
	strategy["callingCode"] = "Prefix"
	strategy["relevance"] = "Relevance"

	b, _ := csvb.NewBinder(r, opts)
	b.ForEach(func(r csvb.Row) (bool, error) {

		var c Country
		if err := r.Bind(&c, strategy); err != nil {
			return false, err
		}

		if strings.Contains(c.Prefix, ",") {
			for _, part := range strings.Split(c.Prefix, ",") {
				c.Prefix = part
				updateCountry(c)
			}
		} else {
			updateCountry(c)
		}

		return true, nil
	})
}

func updateCountry(c Country) {
	_, ok := prefixes[c.Prefix]
	if ok {
		if prefixes[c.Prefix].Relevance.Cmp(c.Relevance) < 0 {
			prefixes[c.Prefix] = c
		}
	} else {
		prefixes[c.Prefix] = c
	}
}

func Lookup(number string) (Country, error) {
	p := number[0:2]
	c, ok := prefixes[p]

	if ok {
		if p == "39" && number[0:5] == "39066" {
			c, ok = prefixes["39066"]
		}
		return c, nil
	}

	p = number[0:1]

	switch p {
	case "7":
		{
			c, ok = prefixes[p]
		}
	case "1":
		{
			p = number[1:4]
			areaCode, ok := nanp[p]
			if !ok {
				return Country{}, fmt.Errorf("Resolved NANP code %s for number %s", p, number)
			} else {
				c.Name = areaCode.countryName
				c.Prefix = areaCode.code
				return c, nil
			}

		}
	}

	if !ok {
		p = number[0:3]
		c, ok = prefixes[p]
	}

	if !ok {
		return Country{}, fmt.Errorf("Resolved world code %s for number %s", p, number)
	} else {
		return c, nil
	}
}
