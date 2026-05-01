package route53

import awsRoute53 "github.com/aws/aws-sdk-go-v2/service/route53/types"

type HostedZone struct {
	Id     string `json:"Id"`
	Name   string `json:"Name"`
	Config struct {
		Comment     string `json:"Comment"`
		PrivateZone bool   `json:"PrivateZone"`
	} `json:"Config"`
	ResourceRecordSetCount int64 `json:"ResourceRecordSetCount"`
}

func NewHostedZone(awsRoute53HostedZone awsRoute53.HostedZone) HostedZone {

	hz := HostedZone{
		Id:                     *awsRoute53HostedZone.Id,
		Name:                   *awsRoute53HostedZone.Name,
		ResourceRecordSetCount: *awsRoute53HostedZone.ResourceRecordSetCount,
	}
	if awsRoute53HostedZone.Config.Comment != nil {
		hz.Config.Comment = *awsRoute53HostedZone.Config.Comment
	}
	hz.Config.PrivateZone = awsRoute53HostedZone.Config.PrivateZone
	return hz
}
