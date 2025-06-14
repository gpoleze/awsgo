package utils

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/samber/lo"
	"github.com/urfave/cli/v3"
	"io"
	"os"
	"os/exec"
	"strings"
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

func SelectWithFzf[T any](items []T, function func(T, int) string) (string, error) {
	mappedItems := lo.Map(items, function)
	data := strings.NewReader(strings.Join(mappedItems, "\n"))
	if selected, err := Fzf(data); err == nil {
		return selected, nil
	} else {
		return "", err
	}
}

func Fzf(data io.Reader) (string, error) {
	var result strings.Builder
	cmd := exec.Command("fzf")
	cmd.Stdout = &result
	cmd.Stderr = os.Stderr

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return "", err
	}
	_, err = io.Copy(stdin, data)
	//_, err = data.WriteTo(stdin)
	if err != nil {
		return "", err
	}
	err = stdin.Close()
	if err != nil {
		return "", err
	}

	err = cmd.Start()
	if err != nil {
		return "", err
	}

	err = cmd.Wait()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(result.String()), nil

}
