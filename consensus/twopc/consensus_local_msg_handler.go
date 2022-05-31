package twopc

import (
	carriertwopcpb "github.com/datumtechs/datum-network-carrier/pb/carrier/netmsg/consensus/twopc"
	"github.com/datumtechs/datum-network-carrier/types"
	"github.com/libp2p/go-libp2p-core/peer"
)

func (t *Twopc) sendLocalPrepareMsg(pid peer.ID, req *carriertwopcpb.PrepareMsg) error {
	return t.onPrepareMsg(pid, &types.PrepareMsgWrap{PrepareMsg: req}, types.LocalNetworkMsg)
}

func (t *Twopc) sendLocalPrepareVote(pid peer.ID, req *carriertwopcpb.PrepareVote) error {
	return t.onPrepareVote(pid, &types.PrepareVoteWrap{PrepareVote: req}, types.LocalNetworkMsg)
}

func (t *Twopc) sendLocalConfirmMsg(pid peer.ID, req *carriertwopcpb.ConfirmMsg) error {
	return t.onConfirmMsg(pid, &types.ConfirmMsgWrap{ConfirmMsg: req}, types.LocalNetworkMsg)
}

func (t *Twopc) sendLocalConfirmVote(pid peer.ID, req *carriertwopcpb.ConfirmVote) error {
	return t.onConfirmVote(pid, &types.ConfirmVoteWrap{ConfirmVote: req}, types.LocalNetworkMsg)
}

func (t *Twopc) sendLocalCommitMsg(pid peer.ID, req *carriertwopcpb.CommitMsg) error {
	return t.onCommitMsg(pid, &types.CommitMsgWrap{CommitMsg: req}, types.LocalNetworkMsg)
}