package utils

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/urfave/cli/v3"
)

func GetClient(
	ctx context.Context,
	command *cli.Command,
	fromConfig func(cfg aws.Config, optFns ...func(*ec2.Options)) *ec2.Client,
) (*ec2.Client, error) {
	region := command.String("region")
	profile := command.String("profile")
	cfg, errCfg := config.LoadDefaultConfig(ctx, config.WithRegion(region), config.WithSharedConfigProfile(profile))
	if errCfg != nil {
		return nil, errCfg
	}

	client := fromConfig(cfg)
	return client, nil
}
