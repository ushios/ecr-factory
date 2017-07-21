package ecrf

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ecr"
)

// CreateRepository create name repository
func CreateRepository(c ecr.ECR, name string) (*ecr.Repository, error) {
	i := &ecr.CreateRepositoryInput{
		RepositoryName: aws.String(name),
	}

	resp, err := c.CreateRepository(i)
	if err != nil {
		return nil, err
	}

	return resp.Repository, nil
}
