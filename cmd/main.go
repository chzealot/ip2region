package main

import (
	"fmt"
	"github.com/chzealot/ip2region"
	"github.com/chzealot/ip2region/assets"
)

func main() {
	fmt.Println("data size", len(assets.QQWryContent))

	r, err := ip2region.NewRequestor()
	if err != nil {
		panic(err)
	}
	for _, ip := range []string{"8.8.8.8", "155.69.203.4"} {
		loc, err := r.Query(ip)
		if err != nil {
			panic(err)
		}
		fmt.Println(loc.String())
	}
	loc, err := r.Query("114.114.114.114")
	if err != nil {
		panic(err)
	}
	fmt.Println(loc.String())
}
