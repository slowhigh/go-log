// Concrete handler
package chainofresponsibility

import (
	"fmt"
)

type doctor struct {
	next department
}

func (d *doctor) execute(p *patient) {
	if p.doctorCheckUpDone {
		fmt.Println("Doctor checkup already done")
	} else {	
		fmt.Println("Doctor checking patient")
		p.doctorCheckUpDone = true
	}

	if d.next != nil {
		d.next.execute(p)
	}
}

func (d *doctor) setNext(next department) {
	d.next = next
}