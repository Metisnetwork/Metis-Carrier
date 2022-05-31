package types

import (
	"bytes"
	"fmt"
	"github.com/datumtechs/datum-network-carrier/common"
	"github.com/datumtechs/datum-network-carrier/common/bytesutil"
	"github.com/datumtechs/datum-network-carrier/common/rlputil"
	carriernetmsgtaskmngpb "github.com/datumtechs/datum-network-carrier/pb/carrier/netmsg/taskmng"
	carriertypespb "github.com/datumtechs/datum-network-carrier/pb/carrier/types"
)

type TaskResultMsg struct {
	MsgOption     *MsgOption
	TaskEventList []*carriertypespb.TaskEvent
	CreateAt      uint64
	Sign          []byte
}

func (msg *TaskResultMsg) String() string {
	return fmt.Sprintf(`{"msgOption": %s, "createAt": %d, "sign": %v}`,
		msg.GetMsgOption().String(), msg.GetCreateAt(), msg.GetSign())
}
func (msg *TaskResultMsg) Hash() common.Hash {

	/**
	MsgOption     *MsgOption
	TaskEventList []*libtypes.TaskEvent
	CreateAt      uint64
	*/

	var buf bytes.Buffer

	buf.Write(msg.GetMsgOption().Hash().Bytes())

	eventBytes := make([]byte, 0)
	for _, event := range msg.GetTaskEventList() {

		/**
		Type                 string
		TaskId               string
		IdentityId           string
		PartyId              string
		Content              string
		CreateAt             uint64
		*/
		eventBytes = append(eventBytes, []byte(event.GetType())...)
		eventBytes = append(eventBytes, []byte(event.GetTaskId())...)
		eventBytes = append(eventBytes, []byte(event.GetIdentityId())...)
		eventBytes = append(eventBytes, []byte(event.GetPartyId())...)
		eventBytes = append(eventBytes, []byte(event.GetContent())...)
		eventBytes = append(eventBytes, bytesutil.Uint64ToBytes(event.GetCreateAt())...)
	}

	buf.Write(eventBytes)
	buf.Write(bytesutil.Uint64ToBytes(msg.GetCreateAt()))

	v := rlputil.RlpHash(buf.Bytes())
	return v
}


func (msg *TaskResultMsg) GetMsgOption() *MsgOption                { return msg.MsgOption }
func (msg *TaskResultMsg) GetTaskEventList() []*carriertypespb.TaskEvent { return msg.TaskEventList }
func (msg *TaskResultMsg) GetCreateAt() uint64                     { return msg.CreateAt }
func (msg *TaskResultMsg) GetSign() []byte                         { return msg.Sign }

func ConvertTaskResultMsg(msg *TaskResultMsg) *carriernetmsgtaskmngpb.TaskResultMsg {
	return &carriernetmsgtaskmngpb.TaskResultMsg{
		MsgOption:     ConvertMsgOption(msg.GetMsgOption()),
		TaskEventList: ConvertTaskEventArr(msg.GetTaskEventList()),
		CreateAt:      msg.GetCreateAt(),
		Sign:          msg.GetSign(),
	}
}

func FetchTaskResultMsg(msg *carriernetmsgtaskmngpb.TaskResultMsg) *TaskResultMsg {
	return &TaskResultMsg{
		MsgOption:     FetchMsgOption(msg.GetMsgOption()),
		TaskEventList: FetchTaskEventArr(msg.GetTaskEventList()),
		CreateAt:      msg.GetCreateAt(),
		Sign:          msg.GetSign(),
	}
}

type TaskResourceUsageMsg struct {
	MsgOption *MsgOption
	Usage     *TaskResuorceUsage
	CreateAt  uint64
	Sign      []byte
}

func (msg *TaskResourceUsageMsg) String() string {
	return fmt.Sprintf(`{"msgOption": %s, "usage": %s, "createAt": %d, "sign": %v}`,
		msg.GetMsgOption().String(), msg.GetUsage().String(), msg.GetCreateAt(), msg.GetSign())
}
func (msg *TaskResourceUsageMsg) Hash() common.Hash {

	/**
	MsgOption *MsgOption
	Usage     *TaskResuorceUsage
	CreateAt  uint64
	*/

	var buf bytes.Buffer

	buf.Write(msg.GetMsgOption().Hash().Bytes())
	buf.Write(msg.GetUsage().Hash().Bytes())
	buf.Write(bytesutil.Uint64ToBytes(msg.GetCreateAt()))

	v := rlputil.RlpHash(buf.Bytes())
	return v
}

func (msg *TaskResourceUsageMsg) GetMsgOption() *MsgOption     { return msg.MsgOption }
func (msg *TaskResourceUsageMsg) GetUsage() *TaskResuorceUsage { return msg.Usage }
func (msg *TaskResourceUsageMsg) GetCreateAt() uint64          { return msg.CreateAt }
func (msg *TaskResourceUsageMsg) GetSign() []byte              { return msg.Sign }

func FetchTaskResourceUsageMsg(msg *carriernetmsgtaskmngpb.TaskResourceUsageMsg) *TaskResourceUsageMsg {
	return &TaskResourceUsageMsg{
		MsgOption: FetchMsgOption(msg.GetMsgOption()),
		Usage: NewTaskResuorceUsage(
			string(msg.GetTaskId()),
			string(msg.GetMsgOption().GetSenderPartyId()),
			msg.GetUsage().GetTotalMem(),
			msg.GetUsage().GetTotalBandwidth(),
			msg.GetUsage().GetTotalDisk(),
			msg.GetUsage().GetUsedMem(),
			msg.GetUsage().GetUsedBandwidth(),
			msg.GetUsage().GetUsedDisk(),
			uint32(msg.GetUsage().GetTotalProcessor()),
			uint32(msg.GetUsage().GetUsedProcessor()),
		),
		CreateAt: msg.GetCreateAt(),
		Sign:     msg.GetSign(),
	}
}

type TaskTerminateTaskMngMsg struct {
	MsgOption *MsgOption
	TaskId    string
	CreateAt  uint64
	Sign      []byte
}

func (msg *TaskTerminateTaskMngMsg) String() string {
	return fmt.Sprintf(`{"msgOption": %s, "taskId": %s, "createAt": %d, "sign": %v}`,
		msg.GetMsgOption().String(), msg.GetTaskId(), msg.GetCreateAt(), msg.GetSign())
}
func (msg *TaskTerminateTaskMngMsg) Hash() common.Hash {

	/**
	MsgOption *MsgOption
	TaskId    string
	CreateAt  uint64
	*/

	var buf bytes.Buffer

	buf.Write(msg.GetMsgOption().Hash().Bytes())
	buf.Write([]byte(msg.GetTaskId()))
	buf.Write(bytesutil.Uint64ToBytes(msg.GetCreateAt()))

	v := rlputil.RlpHash(buf.Bytes())
	return v
}

func (msg *TaskTerminateTaskMngMsg) GetMsgOption() *MsgOption { return msg.MsgOption }
func (msg *TaskTerminateTaskMngMsg) GetTaskId() string        { return msg.TaskId }
func (msg *TaskTerminateTaskMngMsg) GetCreateAt() uint64      { return msg.CreateAt }
func (msg *TaskTerminateTaskMngMsg) GetSign() []byte          { return msg.Sign }

func FetchTaskTerminateTaskMngMsg(msg *carriernetmsgtaskmngpb.TaskTerminateMsg) *TaskTerminateTaskMngMsg {
	return &TaskTerminateTaskMngMsg{
		MsgOption: FetchMsgOption(msg.GetMsgOption()),
		TaskId:    string(msg.GetTaskId()),
		CreateAt:  msg.GetCreateAt(),
		Sign:      msg.GetSign(),
	}
}