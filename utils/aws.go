package utils

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/urfave/cli/v3"
)

func GetClient[T any, U any](
	ctx context.Context,
	command *cli.Command,
	fromConfig func(cfg aws.Config, optFns ...func(U)) T,
) (T, error) {
	region := command.String("region")
	profile := command.String("profile")
	cfg, errCfg := config.LoadDefaultConfig(ctx, config.WithRegion(region), config.WithSharedConfigProfile(profile))
	if errCfg != nil {
		var zero T
		return zero, errCfg
	}

	client := fromConfig(cfg)
	return client, nil
}
