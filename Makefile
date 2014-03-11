bindata: data.go

data.go: data/countries.csv \
		 data/usa.csv \
		 data/canada.csv \
		 data/caribbean.csv \
		 data/non_geo.csv
	go-bindata -pkg=prefixes -o=./data.go data

default: bindata
