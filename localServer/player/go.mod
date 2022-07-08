module github.com/mux-mul-mem/golang-practice/localServer/player

go 1.18

replace local.package/api => ./api

require local.package/api v0.0.0-00010101000000-000000000000

require (
	github.com/PuerkitoBio/goquery v1.8.0 // indirect
	github.com/andybalholm/cascadia v1.3.1 // indirect
	github.com/tidwall/gjson v1.14.1 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.0 // indirect
	golang.org/x/net v0.0.0-20210916014120-12bc252f5db8 // indirect
)
