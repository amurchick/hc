// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/amurchick/hc/characteristic"
)

const TypeFan = "40"

type Fan struct {
	*Service

	On *characteristic.On
}

func NewFan() *Fan {
	svc := Fan{}
	svc.Service = New(TypeFan)

	svc.On = characteristic.NewOn()
	svc.AddCharacteristic(svc.On.Characteristic)

	return &svc
}
