package core

import (
	"fmt"
	"github.com/RosettaFlow/Carrier-Go/common/timeutils"
	"github.com/RosettaFlow/Carrier-Go/core/rawdb"
	"github.com/RosettaFlow/Carrier-Go/lib/center/api"
	"github.com/RosettaFlow/Carrier-Go/types"
)

// about metaData

// on local
func (dc *DataCenter) StoreLocalMetadata(metadata *types.Metadata) error {
	dc.mu.Lock()
	defer dc.mu.Unlock()
	rawdb.WriteLocalMetadata(dc.db, metadata)
	return nil
}

func (dc *DataCenter) GetLocalMetadataByDataId(metadataId string) (*types.Metadata, error) {
	dc.mu.RLock()
	defer dc.mu.RUnlock()
	metadata, err := rawdb.ReadLocalMetadata(dc.db, metadataId)
	return metadata, err
}

func (dc *DataCenter) GetLocalMetadataList() (types.MetadataArray, error) {
	dc.mu.RLock()
	defer dc.mu.RUnlock()
	return rawdb.ReadAllLocalMetadata(dc.db)
}

// on datecenter
func (dc *DataCenter) InsertMetadata(metadata *types.Metadata) error {
	dc.serviceMu.Lock()
	defer dc.serviceMu.Unlock()
	response, err := dc.client.SaveMetadata(dc.ctx, types.NewMetadataSaveRequest(metadata))
	if err != nil {
		log.WithError(err).WithField("hash", metadata.Hash()).Errorf("InsertMetadata failed")
		return err
	}
	if response.Status != 0 {
		return fmt.Errorf("insert metadata error: %s", response.Msg)
	}
	return nil
}

func (dc *DataCenter) RevokeMetadata(metadata *types.Metadata) error {
	dc.serviceMu.Lock()
	defer dc.serviceMu.Unlock()
	response, err := dc.client.RevokeMetadata(dc.ctx, types.NewMetadataRevokeRequest(metadata))
	if err != nil {
		log.WithError(err).WithField("hash", metadata.Hash()).Errorf("RevokeMetadata failed")
		return err
	}
	if response.Status != 0 {
		return fmt.Errorf("revoke metadata error: %s", response.Msg)
	}
	return nil
}

func (dc *DataCenter) GetMetadataByDataId(dataId string) (*types.Metadata, error) {
	dc.serviceMu.Lock()
	defer dc.serviceMu.Unlock()
	metadataByIdResponse, err := dc.client.GetMetadataById(dc.ctx, &api.FindMetadataByIdRequest{
		MetadataId: dataId,
	})
	return types.NewMetadataFromResponse(metadataByIdResponse), err
}

func (dc *DataCenter) GetMetadataList() (types.MetadataArray, error) {
	dc.serviceMu.Lock()
	defer dc.serviceMu.Unlock()
	metaDataListResponse, err := dc.client.GetMetadataList(dc.ctx, &api.ListMetadataRequest{
		LastUpdated:      uint64( timeutils.Now().Unix()),
	})
	return types.NewMetadataArrayFromDetailListResponse(metaDataListResponse), err
}

