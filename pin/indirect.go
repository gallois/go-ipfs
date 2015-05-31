package pin

import "github.com/ipfs/go-ipfs/util"

type indirectPin struct {
	refCounts map[util.Key]uint64
}

func newIndirectPin() *indirectPin {
	return &indirectPin{
		refCounts: make(map[util.Key]uint64),
	}
}

func (i *indirectPin) Increment(k util.Key) {
	i.refCounts[k]++
}

func (i *indirectPin) Decrement(k util.Key) {
	if i.refCounts[k] == 0 {
		log.Warningf("pinning: bad call: asked to unpin nonexistent indirect key: %v", k)
		return
	}
	i.refCounts[k]--
	if i.refCounts[k] == 0 {
		delete(i.refCounts, k)
	}
}

func (i *indirectPin) HasKey(k util.Key) bool {
	_, found := i.refCounts[k]
	return found
}

func (i *indirectPin) GetRefs() map[util.Key]uint64 {
	return i.refCounts
}
