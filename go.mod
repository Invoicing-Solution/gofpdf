module github.com/Invoicing-Solution/gofpdf

go 1.22

require (
	github.com/boombuler/barcode v1.0.2
	github.com/phpdave11/gofpdi v1.0.13
	github.com/ruudk/golang-pdf417 v0.0.0-20201230142125-a7e3863a1245
	golang.org/x/image v0.20.0
)

require github.com/pkg/errors v0.9.1 // indirect

replace gofpdf => ./
