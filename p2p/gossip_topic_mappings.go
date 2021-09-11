package p2p

import (
	twopcpb "github.com/RosettaFlow/Carrier-Go/lib/netmsg/consensus/twopc"
	taskmngpb "github.com/RosettaFlow/Carrier-Go/lib/netmsg/taskmng"
	librpcpb "github.com/RosettaFlow/Carrier-Go/lib/rpc/v1"
	"github.com/gogo/protobuf/proto"
	"reflect"
)

// GossipTopicMappings represent the protocol ID to protobuf message type map for easy lookup.
var GossipTopicMappings = map[string]proto.Message{
	GossipTestDataTopicFormat:   &librpcpb.GossipTestData{},
	TwoPcPrepareMsgTopicFormat:  &twopcpb.PrepareMsg{},
	TwoPcPrepareVoteTopicFormat: &twopcpb.PrepareVote{},
	TwoPcConfirmMsgTopicFormat:  &twopcpb.ConfirmMsg{},
	TwoPcConfirmVoteTopicFormat: &twopcpb.ConfirmVote{},
	TwoPcCommitMsgTopicFormat:   &twopcpb.CommitMsg{},
	TaskResultMsgTopicFormat:    &taskmngpb.TaskResultMsg{},
}

// GossipTypeMapping is the inverse of GossipTopicMappings so that an arbitrary protobuf message
// can be mapped to a protocol ID string.
var GossipTypeMapping = make(map[reflect.Type]string, len(GossipTopicMappings))

func init() {
	for k, v := range GossipTopicMappings {
		GossipTypeMapping[reflect.TypeOf(v)] = k
	}
}
