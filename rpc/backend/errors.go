package backend

type RpcBizErr struct {
	Code int32
	Msg  string
}

func NewRpcBizErr(code int32, msg string) *RpcBizErr { return &RpcBizErr{Code: code, Msg: msg} }

func (e *RpcBizErr) Error() string {
	return e.Msg
}

func (e *RpcBizErr) ErrCode() int32 {
	return e.Code
}

// common
var (
	//OK               = &RpcBizErr{Code: 0, Msg: "OK"}
	ErrSystemPanic      = &RpcBizErr{Code: 00001, Msg: "System error"}
	ErrRequireParams    = &RpcBizErr{Code: 00002, Msg: "Require params"}
	ErrRequirePrivateIP = &RpcBizErr{Code: 00003, Msg: "Require private ip"}
)

// auth
var (
	ErrApplyIdentityMsg         = &RpcBizErr{Code: 10001, Msg: "ApplyIdentity failed"}
	ErrRevokeIdentityMsg        = &RpcBizErr{Code: 10002, Msg: "RevokeIdentity failed"}
	ErrQueryNodeIdentity        = &RpcBizErr{Code: 10003, Msg: "Query local identity failed"}
	ErrQueryIdentityList        = &RpcBizErr{Code: 10004, Msg: "Query all identity list failed"}
	ErrQueryAuthorityList       = &RpcBizErr{Code: 10005, Msg: "Query all authority list failed"}
	ErrApplyMetadataAuthority   = &RpcBizErr{Code: 10006, Msg: "ApplyMetadataAuthority failed"}
	ErrRevokeMetadataAuthority  = &RpcBizErr{Code: 10007, Msg: "RevokeMetadataAuthority failed"}
	ErrAuditMetadataAuth        = &RpcBizErr{Code: 10008, Msg: "AuditMetadataAuth failed"}
	ErrUpdateIdentityCredential = &RpcBizErr{Code: 10009, Msg: "UpdateIdentityCredential params error"}
)

// metadata
var (
	ErrQueryMetadataDetail                 = &RpcBizErr{Code: 11001, Msg: "Query metadata detail failed"}
	ErrQueryMetadataDetailList             = &RpcBizErr{Code: 11002, Msg: "Query metadata detail list failed"}
	ErrPublishMetadataMsg                  = &RpcBizErr{Code: 11003, Msg: "PublishMetaData failed"}
	ErrRevokeMetadataMsg                   = &RpcBizErr{Code: 11004, Msg: "RevokemetaData failed"}
	ErrQueryMetadataUsedTaskIdList         = &RpcBizErr{Code: 11005, Msg: "Query metadata used taskId list failed"}
	ErrQueryMetadataDetailById             = &RpcBizErr{Code: 11007, Msg: "Query metadata detail failed by metadataId"}
	ErrIdentityNotEqualLocalNodeIdentityId = &RpcBizErr{Code: 11008, Msg: "The id of the organization to which the metadata belongs does not match the id of the local organization, and the metadata information cannot be updated"}
)

// power
var (
	ErrRevokePowerMsg       = &RpcBizErr{Code: 12001, Msg: "RevokePower failed"}
	ErrQueryGlobalPowerList = &RpcBizErr{Code: 12002, Msg: "Query global power list failed"}
	ErrQueryLocalPowerList  = &RpcBizErr{Code: 12003, Msg: "Query local power list failed"}
	ErrPublishPowerMsg      = &RpcBizErr{Code: 12004, Msg: "PublishPower failed"}
)

// task
var (
	ErrQueryLocalTaskList  = &RpcBizErr{Code: 13001, Msg: "Query all tasks of current node failed"}
	ErrQueryGlobalTaskList = &RpcBizErr{Code: 13002, Msg: "Query all tasks of global node failed"}
	ErrQueryTaskListByIds  = &RpcBizErr{Code: 13003, Msg: "Query all tasks by taskIds failed"}
	ErrQueryTaskEventList  = &RpcBizErr{Code: 13004, Msg: "Query all events of task by taskId failed"}
	ErrPublishTaskMsg      = &RpcBizErr{Code: 13005, Msg: "PublishTask failed"}
	ErrTerminateTaskMsg    = &RpcBizErr{Code: 13006, Msg: "TerminateTask failed, send task terminate msg failed"}
	ErrEstimateTaskGas     = &RpcBizErr{Code: 13007, Msg: "Estimate task gas failed"}
)

// yarn
var (
	ErrSetSeedNode                    = &RpcBizErr{Code: 14001, Msg: "Set seed node failed"}
	ErrDeleteSeedNode                 = &RpcBizErr{Code: 14002, Msg: "Delete seed node failed"}
	ErrQuerySeedNodeList              = &RpcBizErr{Code: 14003, Msg: "Query seed nodes failed"}
	ErrSetDataNode                    = &RpcBizErr{Code: 14011, Msg: "Set data node failed"}
	ErrDeleteDataNode                 = &RpcBizErr{Code: 14012, Msg: "Delete data node failed"}
	ErrQueryDataNodeList              = &RpcBizErr{Code: 14013, Msg: "Query data nodes failed"}
	ErrSetJobNode                     = &RpcBizErr{Code: 14021, Msg: "Set job node failed"}
	ErrDeleteJobNode                  = &RpcBizErr{Code: 14022, Msg: "Delete job node failed"}
	ErrQueryJobNodeList               = &RpcBizErr{Code: 14023, Msg: "Query job nodes failed"}
	ErrReportTaskEvent                = &RpcBizErr{Code: 14031, Msg: "Report taskEvent failed"}
	ErrReportUpDataSummary            = &RpcBizErr{Code: 14032, Msg: "Report upDataSummary failed"}
	ErrQueryTaskResultDataSummary     = &RpcBizErr{Code: 14033, Msg: "Query taskResultDataSummary failed"}
	ErrQueryNodeInfo                  = &RpcBizErr{Code: 14034, Msg: "Query yarn node information failed"}
	ErrReportTaskResultDataSummary    = &RpcBizErr{Code: 14035, Msg: "Report taskResultDataSummary failed"}
	ErrQueryAvailableDataNode         = &RpcBizErr{Code: 14036, Msg: "Query availableDataNode failed"}
	ErrQueryFilePosition              = &RpcBizErr{Code: 14037, Msg: "Query filePosition failed"}
	ErrQueryTaskResultDataSummaryList = &RpcBizErr{Code: 14038, Msg: "Query taskResultDataSummary list failed"}
	ErrReportTaskResourceExpense      = &RpcBizErr{Code: 14039, Msg: "Report taskResourceExpense failed"}
	ErrGenerateObserverProxyWallet    = &RpcBizErr{Code: 14040, Msg: "Generate observer proxy wallet failed"}
)

// did
var (
	ErrCreateDID        = &RpcBizErr{Code: 15001, Msg: "Create DID failed"}
	ErrDIDExists        = &RpcBizErr{Code: 15002, Msg: "DID exists"}
	ErrApplyVC          = &RpcBizErr{Code: 15003, Msg: "Apply VC failed"}
	ErrCreateVC         = &RpcBizErr{Code: 15004, Msg: "Create VC failed"}
	ErrDownloadVC       = &RpcBizErr{Code: 15005, Msg: "Download VC failed"}
	ErrSubmitProposal   = &RpcBizErr{Code: 15006, Msg: "Submit proposal failed"}
	ErrWithdrawProposal = &RpcBizErr{Code: 15007, Msg: "Withdraw proposal failed"}
	ErrVoteProposal     = &RpcBizErr{Code: 15008, Msg: "Vote proposal failed"}
	ErrEffectProposal   = &RpcBizErr{Code: 15009, Msg: "Effect proposal failed"}
)
