package types

import (
	"github.com/RosettaFlow/Carrier-Go/common"
	libTypes "github.com/RosettaFlow/Carrier-Go/lib/types"
	"sync/atomic"
)

type Metadata struct {
	data *libTypes.MetaData

	// caches
	hash atomic.Value
	size atomic.Value
}

func NewMetadata(data *libTypes.MetaData) *Metadata {
	return &Metadata{data: data}
}

func (m *Metadata) EncodePb() ([]byte, error) {
	return m.data.Marshal()
}

func (m *Metadata) DecodePb(data []byte) error {
	m.size.Store(common.StorageSize(len(data)))
	return m.data.Unmarshal(data)
}

func (m *Metadata) Hash() common.Hash {
	if hash := m.hash.Load(); hash != nil {
		return hash.(common.Hash)
	}
	data, _ := m.EncodePb()
	v := protoBufHash(data)
	m.hash.Store(v)
	return v
}

// MetadataArray is a Transaction slice type for basic sorting.
type MetadataArray []*libTypes.MetaData

// Len returns the length of s.
func (s MetadataArray) Len() int { return len(s) }

// Swap swaps the i'th and the j'th element in s.
func (s MetadataArray) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s MetadataArray) GetPb(i int) []byte {
	enc, _ := s[i].Marshal()
	return enc
}

