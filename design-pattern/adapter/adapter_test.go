package adapter

import "testing"

func Test_adapter(t *testing.T) {
	client := &client{}
	mac, windows := &mac{}, &windows{}

	// client(LightningPort)  --O-->  mac(LightninPort)
	if !client.insertLightningConnectorIntoComputer(mac) {
		t.Error("client(LightningPort)  --X-->  mac(LightninPort)")
	}
	// client(LightningPort)  --X-->  windows(USB)
	if client.insertLightningConnectorIntoComputer(windows) {
		t.Error("client(LightningPort)  --O-->  mac(LightninPort)")
	}

	// windowsAdapter(LightningPort 🔀 USB)
	windowsAdapter := &windowsAdapter{
		windowsMachine: windows,
	}

	// client(LightningPort)  --O-->  windowsAdapter(LightningPort 🔀 USB)  --O--> windows(USB)
	if !client.insertLightningConnectorIntoComputer(windowsAdapter) {
		t.Error("Invalid adapter function.")
	}
}
