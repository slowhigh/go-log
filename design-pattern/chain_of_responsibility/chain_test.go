// Client code
package chainofresponsibility

import "testing"

func Test_Chain(t *testing.T) {
	cashier := &cashier{}

	medical := &medical{}
	medical.setNext(cashier)

	doctor := &doctor{}
	doctor.setNext(medical)

	reception := &reception{}
	reception.setNext(doctor)

	patient := &patient{name: "abc"}
	reception.execute(patient)

	if !patient.paymentDone || !patient.medicineDone || !patient.doctorCheckUpDone || !patient.registrationDone {
		t.Error("Invalid chain of chainofresponsibility")
	}
}