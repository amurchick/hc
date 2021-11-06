package rtp

import (
	"testing"

	"github.com/amurchick/hc/characteristic"
	"github.com/amurchick/hc/tlv8"
)

func TestSetupEndpoints(t *testing.T) {
	c := characteristic.NewSetupEndpoints()
	c.Value = "ARBz21VuCupGZre3A62biD8XAxkBAQACDDE5Mi4xNjguMC4xMwMC+OUEAurRBCUCEPpLBUWQEzkfFiGd1qkieqoDDi8cIMO0Vl1+kegzGgnpAQEABSUCEJQ27Ze9EEmuxcIVPhDEs68DDlaHwww6f6d5+NSClT7TAQEA"

	b := c.GetValue()

	var setup SetupEndpoints
	err := tlv8.Unmarshal(b, &setup)
	if err != nil {
		t.Fatal(err)
	}
}
