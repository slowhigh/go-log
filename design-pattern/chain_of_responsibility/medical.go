// Concrete handler
package chainofresponsibility

import "fmt"

type medical struct {
	next department
}

func (m *medical) execute(p *patient) {
	if p.medicineDone {
		fmt.Println("Medical already given to patient")
	} else {
		fmt.Println("Medical giving medicine to patient")
		p.medicineDone = true
	}

	if m.next != nil {
		m.next.execute(p)
	}
}

func (m *medical) setNext(next department) {
	m.next = next
}