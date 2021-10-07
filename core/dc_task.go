package core

import (
	"errors"
	"fmt"
	"github.com/RosettaFlow/Carrier-Go/common/timeutils"
	"github.com/RosettaFlow/Carrier-Go/core/rawdb"
	"github.com/RosettaFlow/Carrier-Go/lib/center/api"
	libtypes "github.com/RosettaFlow/Carrier-Go/lib/types"
	"github.com/RosettaFlow/Carrier-Go/types"
)

// about task on local
// local task
func (dc *DataCenter) StoreLocalTask(task *types.Task) error {
	if task == nil {
		return errors.New("invalid params for task")
	}
	dc.mu.Lock()
	defer dc.mu.Unlock()
	rawdb.WriteLocalTask(dc.db, task)
	return nil
}

func (dc *DataCenter) RemoveLocalTask(taskId string) error {
	if taskId == "" {
		return errors.New("invalid params for taskId to DelLocalTask")
	}
	dc.mu.Lock()
	defer dc.mu.Unlock()
	rawdb.DeleteLocalTask(dc.db, taskId)
	return nil
}

func (dc *DataCenter) QueryLocalTask(taskId string) (*types.Task, error) {
	if taskId == "" {
		return nil, errors.New("invalid params taskId for QueryLocalTask")
	}
	dc.mu.Lock()
	defer dc.mu.Unlock()
	//log.Debugf("QueryLocalTask, taskId: {%s}", taskId)
	return rawdb.ReadLocalTask(dc.db, taskId)
}

func (dc *DataCenter) QueryLocalTaskListByIds(taskIds []string) (types.TaskDataArray, error) {
	dc.mu.RLock()
	defer dc.mu.RUnlock()
	return rawdb.ReadLocalTaskByIds(dc.db, taskIds)
}

func (dc *DataCenter) QueryLocalTaskList() (types.TaskDataArray, error) {
	dc.mu.RLock()
	defer dc.mu.RUnlock()
	return rawdb.ReadAllLocalTasks(dc.db)
}

func (dc *DataCenter) QueryLocalTaskAndEvents(taskId string) (*types.Task, error) {
	if taskId == "" {
		return nil, errors.New("invalid params taskId for QueryLocalTask")
	}
	dc.mu.Lock()
	defer dc.mu.Unlock()
	task, err := rawdb.ReadLocalTask(dc.db, taskId)
	if nil != err {
		return nil, err
	}
	list, err := rawdb.ReadTaskEvent(dc.db, taskId)
	if nil != err {
		return nil, err
	}
	task.SetEventList(list)
	return task, nil
}

func (dc *DataCenter) QueryLocalTaskAndEventsListByIds(taskIds []string) (types.TaskDataArray, error) {
	dc.mu.RLock()
	defer dc.mu.RUnlock()
	tasks, err := rawdb.ReadLocalTaskByIds(dc.db, taskIds)
	if nil != err {
		return nil, err
	}
	for i, task := range tasks {
		list, err := rawdb.ReadTaskEvent(dc.db, task.GetTaskId())
		if nil != err {
			return nil, err
		}
		task.SetEventList(list)
		tasks[i] = task
	}
	return tasks, nil
}

func (dc *DataCenter) QueryLocalTaskAndEventsList() (types.TaskDataArray, error) {
	dc.mu.RLock()
	defer dc.mu.RUnlock()

	tasks, err := rawdb.ReadAllLocalTasks(dc.db)
	if nil != err {
		return nil, err
	}
	for i, task := range tasks {
		list, err := rawdb.ReadTaskEvent(dc.db, task.GetTaskId())
		if nil != err {
			return nil, err
		}
		task.SetEventList(list)
		tasks[i] = task
	}
	return tasks, nil
}

func (dc *DataCenter) StoreJobNodeRunningTaskId(jobNodeId, taskId string) error {
	dc.mu.Lock()
	defer dc.mu.Unlock()
	return rawdb.StoreResourceTaskId(dc.db, jobNodeId, taskId)
}

func (dc *DataCenter) RemoveJobNodeRunningTaskId(jobNodeId, taskId string) error {
	dc.mu.Lock()
	defer dc.mu.Unlock()
	return rawdb.RemoveResourceTaskId(dc.db, jobNodeId, taskId)
}

func (dc *DataCenter) QueryRunningTaskCountOnJobNode(jobNodeId string) (uint32, error) {
	dc.mu.RLock()
	defer dc.mu.RUnlock()
	taskIds, err := rawdb.QueryResourceTaskIds(dc.db, jobNodeId)
	if rawdb.IsNoDBNotFoundErr(err) {
		return 0, err
	} else if rawdb.IsDBNotFoundErr(err) {
		return 0, nil
	}
	return uint32(len(taskIds)), nil
}

func (dc *DataCenter) QueryJobNodeRunningTaskIdList(jobNodeId string) ([]string, error) {
	dc.mu.RLock()
	defer dc.mu.RUnlock()
	ids, err := rawdb.QueryResourceTaskIds(dc.db, jobNodeId)
	if rawdb.IsNoDBNotFoundErr(err) {
		return nil, err
	} else if rawdb.IsDBNotFoundErr(err) {
		return nil, nil
	}
	return ids, nil
}

func (dc *DataCenter) IncreaseResourceTaskPartyIdCount (jobNodeId, taskId string) error {
	dc.mu.Lock()
	defer dc.mu.Unlock()
	return rawdb.IncreaseResourceTaskPartyIdCount(dc.db, jobNodeId, taskId)
}

func (dc *DataCenter) DecreaseResourceTaskPartyIdCount (jobNodeId, taskId string) error {
	dc.mu.Lock()
	defer dc.mu.Unlock()
	return rawdb.DecreaseResourceTaskPartyIdCount(dc.db, jobNodeId, taskId)
}

func (dc *DataCenter) QueryResourceTaskPartyIdCount (jobNodeId, taskId string) (uint32, error) {
	dc.mu.Lock()
	defer dc.mu.Unlock()
	return rawdb.QueryResourceTaskPartyIdCount(dc.db, jobNodeId, taskId)
}

func  (dc *DataCenter) IncreaseResourceTaskTotalCount (jobNodeId string) error {
	dc.mu.Lock()
	defer dc.mu.Unlock()
	return rawdb.IncreaseResourceTaskTotalCount(dc.db, jobNodeId)
}

func  (dc *DataCenter) DecreaseResourceTaskTotalCount (jobNodeId string) error {
	dc.mu.Lock()
	defer dc.mu.Unlock()
	return rawdb.DecreaseResourceTaskTotalCount(dc.db, jobNodeId)
}

func  (dc *DataCenter) RemoveResourceTaskTotalCount (jobNodeId string) error {
	dc.mu.Lock()
	defer dc.mu.Unlock()
	return rawdb.RemoveResourceTaskTotalCount(dc.db, jobNodeId)
}

func  (dc *DataCenter) QueryResourceTaskTotalCount (jobNodeId string) (uint32, error) {
	dc.mu.RLock()
	defer dc.mu.RUnlock()
	return rawdb.QueryResourceTaskTotalCount(dc.db, jobNodeId)
}


// about task on datacenter
func (dc *DataCenter) InsertTask(task *types.Task) error {
	dc.serviceMu.Lock()
	defer dc.serviceMu.Unlock()
	response, err := dc.client.SaveTask(dc.ctx, types.NewSaveTaskRequest(task))
	if err != nil {
		log.WithError(err).WithField("taskId", task.GetTaskId()).Errorf("InsertTask failed")
		return err
	}
	if response.Status != 0 {
		return fmt.Errorf("insert task, taskId: {%s},  error: %s", task.GetTaskId(), response.Msg)
	}
	return nil
}

func (dc *DataCenter) QueryTaskListByIdentityId(identityId string) (types.TaskDataArray, error) {
	dc.serviceMu.Lock()
	defer dc.serviceMu.Unlock()
	//taskListResponse, err := dc.client.ListTask(dc.ctx, &api.ListTaskRequest{LastUpdated: uint64(timeutils.UnixMsec())})
	taskListResponse, err := dc.client.ListTaskByIdentity(dc.ctx, &api.ListTaskByIdentityRequest{
		LastUpdated: timeutils.UnixMsecUint64(),
		IdentityId: identityId,
	})
	return types.NewTaskArrayFromResponse(taskListResponse), err
}

func (dc *DataCenter) QueryRunningTaskCountOnOrg() uint32 {
	dc.mu.RLock()
	defer dc.mu.RUnlock()
	taskList, err := rawdb.ReadAllLocalTasks(dc.db)
	if nil != err {
		return 0
	}
	if taskList != nil {
		return uint32(taskList.Len())
	}
	return 0
}

func (dc *DataCenter) QueryTaskEventListByTaskId(taskId string) ([]*libtypes.TaskEvent, error) {
	dc.serviceMu.Lock()
	defer dc.serviceMu.Unlock()
	taskEventResponse, err := dc.client.ListTaskEvent(dc.ctx, &api.ListTaskEventRequest{
		TaskId: taskId,
	})
	return taskEventResponse.TaskEvents, err
}

func (dc *DataCenter) QueryTaskEventListByTaskIds(taskIds []string) ([]*libtypes.TaskEvent, error) {
	dc.serviceMu.Lock()
	defer dc.serviceMu.Unlock()

	eventList := make([]*libtypes.TaskEvent, 0)
	for _, taskId := range taskIds {
		taskEventResponse, err := dc.client.ListTaskEvent(dc.ctx, &api.ListTaskEventRequest{
			TaskId: taskId,
		})
		if nil != err {
			return nil, err
		}
		eventList = append(eventList, taskEventResponse.TaskEvents...)
	}
	return eventList, nil
}


//func (dc *DataCenter) UpdateLocalTaskState(taskId, state string) error {
//	if taskId == "" || state == "" {
//		return errors.New("invalid params taskId or state for UpdateLocalTaskState")
//	}
//	log.Debugf("Start to update local task state, taskId: {%s}, need update state: {%s}", taskId, state)
//
//	dc.mu.Lock()
//	defer dc.mu.Unlock()
//	task, err := rawdb.ReadLocalTask(dc.db, taskId)
//	if nil != err {
//		return err
//	}
//	task.TaskPB().GetState = state
//	rawdb.DeleteLocalTask(dc.db, taskId)
//	rawdb.WriteLocalTask(dc.db, task)
//	return nil
//}
