package rds

import (
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"time"
)

type DBInstance struct {
	DBInstanceIdentifier string
	DBInstanceStatus     string
	VpcId                string
	DBName               string
	DBInstanceClass      string
	Engine               string
	EngineVersion        string
	LatestRestorableTime time.Time
	InstanceCreateTime   time.Time
	EndpointAddress      string
}

func NewDBInstance(dbInstance types.DBInstance) DBInstance {
	var instance = DBInstance{
		DBInstanceIdentifier: *dbInstance.DBInstanceIdentifier,
		DBInstanceStatus:     *dbInstance.DBInstanceStatus,
		VpcId:                *dbInstance.DBSubnetGroup.VpcId,
		DBInstanceClass:      *dbInstance.DBInstanceClass,
		Engine:               *dbInstance.Engine,
		EngineVersion:        *dbInstance.EngineVersion,
		EndpointAddress:      *dbInstance.Endpoint.Address,
	}

	if dbInstance.DBName != nil {
		dbInstance.DBName = dbInstance.DBName
	}
	if dbInstance.InstanceCreateTime != nil {
		dbInstance.InstanceCreateTime = dbInstance.InstanceCreateTime
	}
	if dbInstance.LatestRestorableTime != nil {
		dbInstance.LatestRestorableTime = dbInstance.LatestRestorableTime
	}

	return instance
}
