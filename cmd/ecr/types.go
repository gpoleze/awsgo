package ecr

import "github.com/aws/aws-sdk-go-v2/service/ecr/types"

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
