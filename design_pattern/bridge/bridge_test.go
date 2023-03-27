package bridge

import "testing"

func Test_bridge(t *testing.T) {
	epsonPrinter := &epson{}
	hpPrinter := &hp{}

	windowsCumputer := &windows{}
	windowsCumputer.setPrinter(epsonPrinter)
	windowsCumputer.print()
	windowsCumputer.setPrinter(hpPrinter)
	windowsCumputer.print()

	macComputer := &mac{}
	macComputer.setPrinter(epsonPrinter)
	macComputer.print()
	macComputer.setPrinter(hpPrinter)
	macComputer.print()
}