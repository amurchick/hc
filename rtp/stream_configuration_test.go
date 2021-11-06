package rtp

import (
	"fmt"
	"testing"

	"github.com/amurchick/hc/characteristic"
	"github.com/amurchick/hc/tlv8"
)

func TestSelectedStreamConfiguration(t *testing.T) {
	c := characteristic.NewSelectedStreamConfiguration()
	c.Value = "ARUCAQABEHW8tiJ9E0F4tLlvOURdFCc="

	b := c.GetValue()

	var cfg StreamConfiguration
	err := tlv8.Unmarshal(b, &cfg)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%+v", cfg)
}
