package ecr

import (
	"github.com/aws/aws-sdk-go-v2/service/ecr/types"
	"strings"
	"time"
)

type Repository struct {
	RepositoryName string
	RepositoryUri  string
	CreatedAt      string
}

func NewRepository(repository types.Repository) Repository {
	return Repository{
		RepositoryName: *repository.RepositoryName,
		RepositoryUri:  *repository.RepositoryUri,
		CreatedAt:      repository.CreatedAt.String(),
	}
}

type Image struct {
	Tags      string
	PushedAt  time.Time
	MediaType string
	SizeMB    int64
	Digest    string
}

func NewImage(imageDetail types.ImageDetail) Image {
	image := Image{
		Tags:     strings.Join(imageDetail.ImageTags, ", "),
		PushedAt: *imageDetail.ImagePushedAt,
		SizeMB:   *imageDetail.ImageSizeInBytes / 1024 / 1024,
		Digest:   *imageDetail.ImageDigest,
	}
	if imageDetail.ArtifactMediaType != nil && *imageDetail.ArtifactMediaType == "application/vnd.docker.container.image.v1+json" {
		image.MediaType = "image"
	} else {
		image.MediaType = "Image Index"
	}
	return image
}
