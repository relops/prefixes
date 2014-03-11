package prefixes

import (
	"bytes"
	"fmt"
	"github.com/relops/csvb"
	"speter.net/go/exp/math/dec/inf"
)

type Country struct {
	Name      string
	Prefix    string
	Relevance *inf.Dec
}

var (
	prefixes = make(map[string]Country)
)

func init() {
	initWorldData()
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

		_, ok := prefixes[c.Prefix]
		if ok {
			if prefixes[c.Prefix].Relevance.Cmp(c.Relevance) < 0 {
				prefixes[c.Prefix] = c
			}
		} else {
			prefixes[c.Prefix] = c
		}

		return true, nil
	})
}

func Lookup(number string) (Country, error) {
	p := number[0:2]
	c, ok := prefixes[p]

	if ok {
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
			p = number[0:4]
			c, ok = prefixes[p]
			if !ok {
				p = number[0:1]
				c, ok = prefixes[p]
			}
		}
	}

	if !ok {
		return Country{}, fmt.Errorf("Resolved %s for number %s", p, number)
	} else {
		return c, nil
	}
}
