// Strategy interface
package strategy

type evictionAlgo interface {
	evict(c *cache)
}
