// Collection
package iterator

type collection interface {
	createIterator() iterator
}