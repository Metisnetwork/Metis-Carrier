package p2p

import (
	pb "github.com/Metisnetwork/Metis-Carrier/lib/p2p/v1"
	"github.com/Metisnetwork/Metis-Carrier/params"
	"github.com/ethereum/go-ethereum/p2p/enode"
	"github.com/ethereum/go-ethereum/p2p/enr"
	"time"
)

var enr2ENRKey = params.CarrierNetworkConfig().EnrKey

// forkDigest returns the current fork digest of the node.
func (s *Service) forkDigest() ([4]byte, error) {
	if s.currentForkDigest != [4]byte{} {
		return s.currentForkDigest, nil
	}
	return [4]byte{0x1, 0x1, 0x1, 0x1,},nil
}


// Compares fork ENRs between an incoming peer's record and our node's
// local record values for current and next fork version/epoch.
func (s *Service) compareForkENR(record *enr.Record) error {
	return nil
}

// Adds a fork entry as an ENR record under the xxxxx for
// the local node. The fork entry is an ssz-encoded enrForkID type
// which takes into account the current fork version from the current
// epoch to create a fork digest, the next fork version,
// and the next fork epoch.
func addForkEntry(
	node *enode.LocalNode,
	genesisTime time.Time,
	genesisValidatorsRoot []byte,
) (*enode.LocalNode, error) {
	return node, nil
}

// Retrieves an enrForkID from an ENR record by key lookup under the xxxxx.
func forkEntry(record *enr.Record) (*pb.ENRForkID, error) {
	/*sszEncodedForkEntry := make([]byte, 16)
	entry := enr.WithEntry(enr2ENRKey, &sszEncodedForkEntry)
	err := record.Load(entry)
	if err != nil {
		return nil, err
	}
	forkEntry := &pb.ENRForkID{}
	if err := forkEntry.UnmarshalSSZ(sszEncodedForkEntry); err != nil {
		return nil, err
	}
	return forkEntry, nil*/
	return nil, nil
}

