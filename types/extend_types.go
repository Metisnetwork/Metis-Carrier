package types

import (
	pb "github.com/RosettaFlow/Carrier-Go/lib/api"
	apipb "github.com/RosettaFlow/Carrier-Go/lib/common"
	libTypes "github.com/RosettaFlow/Carrier-Go/lib/types"
)

func NewTaskDetailShowFromTaskData(input *Task, role string) *TaskDetailShow {
	taskData := input.TaskData()
	detailShow := &TaskDetailShow{
		TaskId:   taskData.GetTaskId(),
		TaskName: taskData.GetTaskName(),
		Role:     role,
		Owner: &TaskNodeAlias{
			PartyId:    taskData.GetPartyId(),
			Name:       taskData.GetNodeName(),
			NodeId:     taskData.GetNodeId(),
			IdentityId: taskData.GetIdentityId(),
		},
		AlgoSupplier: &TaskNodeAlias{
			PartyId:    taskData.GetPartyId(),
			Name:       taskData.GetNodeName(),
			NodeId:     taskData.GetNodeId(),
			IdentityId: taskData.GetIdentityId(),
		},
		DataSupplier:  make([]*TaskDataSupplierShow, 0, len(taskData.GetDataSupplier())),
		PowerSupplier: make([]*TaskPowerSupplierShow, 0, len(taskData.GetPowerSupplier())),
		Receivers:     make([]*TaskNodeAlias, 0, len(taskData.GetReceivers())),
		CreateAt:      taskData.GetCreateAt(),
		StartAt:       taskData.GetStartAt(),
		EndAt:         taskData.GetEndAt(),
		State:         taskData.GetState(),
		OperationCost: &TaskOperationCost{
			Processor: uint64(taskData.GetOperationCost().GetCostProcessor()),
			Mem:       taskData.GetOperationCost().GetCostMem(),
			Bandwidth: taskData.GetOperationCost().GetCostBandwidth(),
			Duration:  taskData.GetOperationCost().GetDuration(),
		},
	}
	// DataSupplier
	for _, metadataSupplier := range taskData.GetDataSupplier() {
		dataSupplier := &TaskDataSupplierShow{
			MemberInfo: &TaskNodeAlias{
				PartyId:    metadataSupplier.GetMemberInfo().GetPartyId(),
				Name:       metadataSupplier.GetMemberInfo().GetNodeName(),
				NodeId:     metadataSupplier.GetMemberInfo().GetNodeId(),
				IdentityId: metadataSupplier.GetMemberInfo().GetIdentityId(),
			},
			MetaDataId:   metadataSupplier.GetMetadataId(),
			MetaDataName: metadataSupplier.GetMetadataName(),
		}
		detailShow.DataSupplier = append(detailShow.DataSupplier, dataSupplier)
	}
	// powerSupplier
	for _, data := range taskData.GetPowerSupplier() {
		detailShow.PowerSupplier = append(detailShow.PowerSupplier, &TaskPowerSupplierShow{
			MemberInfo: &TaskNodeAlias{
				PartyId:    data.GetOrganization().GetPartyId(),
				Name:       data.GetOrganization().GetNodeName(),
				NodeId:     data.GetOrganization().GetNodeId(),
				IdentityId: data.GetOrganization().GetIdentityId(),
			},
			ResourceUsage: &ResourceUsage{
				TotalMem:       data.GetResourceUsedOverview().GetTotalMem(),
				UsedMem:        data.GetResourceUsedOverview().GetUsedMem(),
				TotalProcessor: uint64(data.GetResourceUsedOverview().GetTotalProcessor()),
				UsedProcessor:  uint64(data.GetResourceUsedOverview().GetUsedProcessor()),
				TotalBandwidth: data.GetResourceUsedOverview().GetTotalBandwidth(),
				UsedBandwidth:  data.GetResourceUsedOverview().GetUsedBandwidth(),
			},
		})
	}
	// Receivers
	for _, receiver := range taskData.GetReceivers() {
		detailShow.Receivers = append(detailShow.Receivers, &TaskNodeAlias{
			PartyId:    receiver.GetReceiver().GetPartyId(),
			Name:       receiver.GetReceiver().GetNodeName(),
			NodeId:     receiver.GetReceiver().GetNodeId(),
			IdentityId: receiver.GetReceiver().GetIdentityId(),
		})
	}
	return detailShow
}

func NewTaskEventFromAPIEvent(input []*libTypes.TaskEvent) []*pb.TaskEventShow {
	result := make([]*pb.TaskEventShow, 0, len(input))
	for _, event := range input {
		result = append(result, &pb.TaskEventShow{
			TaskId:   event.GetTaskId(),
			Type:     event.GetType(),
			CreateAt: event.GetCreateAt(),
			Content:  event.GetContent(),
			Owner:    &apipb.Organization{
				IdentityId: event.GetIdentityId(),
			},
		})
	}
	return result
}

func NewOrgMetaDataInfoFromMetadata(input *Metadata) *pb.GetMetaDataDetailResponse {
	response := &pb.GetMetaDataDetailResponse{
		Owner: &apipb.Organization{
			NodeName:   input.data.GetNodeName(),
			NodeId:     input.data.GetNodeId(),
			IdentityId: input.data.GetIdentityId(),
		},
		Information: &libTypes.MetadataDetail{
			MetaDataSummary: &libTypes.MetaDataSummary{
				MetaDataId: input.data.GetDataId(),
				OriginId:   input.data.GetOriginId(),
				TableName:  input.data.GetTableName(),
				Desc:       input.data.GetDesc(),
				FilePath:   input.data.GetFilePath(),
				Rows:       uint32(input.data.GetRows()),
				Columns:    uint32(input.data.GetColumns()),
				Size_:      uint32(input.data.GetSize_()),
				FileType:   input.data.GetFileType(),
				HasTitle:   input.data.GetHasTitle(),
				State:      input.data.GetState(),
			},
			MetadataColumnList: input.data.GetMetadataColumnList(),
		},
	}
	return response
}

func NewOrgMetaDataInfoArrayFromMetadataArray(input MetadataArray) []*pb.GetMetaDataDetailResponse {
	result := make([]*pb.GetMetaDataDetailResponse, 0, input.Len())
	for _, metadata := range input {
		if metadata == nil {
			continue
		}
		result = append(result, NewOrgMetaDataInfoFromMetadata(metadata))
	}
	return result
}

func NewOrgResourceFromResource(input *Resource) *RemoteResourceTable {
	return &RemoteResourceTable{
		identityId: input.data.IdentityId,
		total: &resource{
			mem:       input.data.TotalMem,
			processor: input.data.TotalProcessor,
			bandwidth: input.data.TotalBandWidth,
		},
		used: &resource{
			mem:       input.data.UsedMem,
			processor: input.data.UsedProcessor,
			bandwidth: input.data.UsedBandWidth,
		},
	}
}
func NewOrgResourceArrayFromResourceArray(input ResourceArray) []*RemoteResourceTable {
	result := make([]*RemoteResourceTable, input.Len())
	for i, resource := range input {
		result[i] = NewOrgResourceFromResource(resource)
	}
	return result
}
