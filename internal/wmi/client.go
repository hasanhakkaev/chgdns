package wmi

import (
	"github.com/yusufpapurcu/wmi"
	"log"
)

type Win32Process struct {
	name string
}

func NewClient() *wmi.Client {
	var DefaultClient = &wmi.Client{}
	return DefaultClient
}
func listProcess(c wmi.Client) {
	var dst []Win32Process
	q := wmi.CreateQuery(&dst, "")
	err := wmi.Query(q, &dst)
	if err != nil {
		log.Fatal(err)
	}
	for i, v := range dst {
		println(i, v.name)
	}
}
