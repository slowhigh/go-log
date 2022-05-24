// Client code
package command

import "testing"

func Test_Command(t *testing.T) {
	tv := &tv{}

	onCommand := &onCommand{
		device: tv,
	}

	offCommand := &offCommand{
		device: tv,
	}

	onButton := &button{
		command: onCommand,
	}
	onButton.press()

	if !tv.isRunning {
		t.Error("Not working onCommand")
	}

	offButton := &button{
		command: offCommand,
	}
	offButton.press()

	if tv.isRunning {
		t.Error("Not working offCommand")
	}
}