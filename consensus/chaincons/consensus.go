package chaincons

import (
	"github.com/datumtechs/datum-network-carrier/consensus/twopc"
	"github.com/datumtechs/datum-network-carrier/types"
	"github.com/libp2p/go-libp2p-core/peer"
	log "github.com/sirupsen/logrus"
)

type Chaincons struct {

}

func New() *Chaincons {return &Chaincons{}}
func (c *Chaincons)Start() error {
	log.Info("Started chainCons consensus engine ...")
	return nil
}
func (c *Chaincons) Stop() error {
	log.Info("Stopped chainCons consensus engine ...")
	return nil
}
func (c *Chaincons)OnPrepare(task *types.NeedConsensusTask) error {return nil}
func (c *Chaincons)OnHandle(task *types.NeedConsensusTask) error {return nil}
func (c *Chaincons) OnConsensusMsg(pid peer.ID, msg types.ConsensusMsg) error {return nil}
func (c *Chaincons) GetConsensusStateInfo() *twopc.ConsensusStateInfo {return &twopc.ConsensusStateInfo{}}
func (c *Chaincons) OnError() error  {return nil}