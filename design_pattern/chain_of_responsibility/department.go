// Handler interface
package chainofresponsibility

type department interface {
	execute(*patient)
	setNext(department)
}