package ecr

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/service/ecr/types"
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
	tags := ""
	line := 0
	maxLine := 50
	for _, i := range imageDetail.ImageTags {
		if tags == "" {
			tags = i
			line = len(i)
		} else if line+len(i) <= maxLine {
			tags = tags + ", " + i
			line = line + len(i)
		} else {
			tags = tags + ",\n" + i
			line = len(i)
		}
	}
	image := Image{
		Tags:     tags,
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
