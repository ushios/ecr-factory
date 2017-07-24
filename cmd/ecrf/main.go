package main

import (
	"flag"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecr"
	"github.com/aws/aws-sdk-go/service/iam"
	ecrf "github.com/ushios/ecr-factory"
)

var (
	id     = flag.String("id", "", "aws access key id")
	secret = flag.String("secret", "", "aws secret key")
	name   = flag.String("name", "", "repository name")
)

func main() {

	flag.Parse()
	fmt.Printf("id: %s, secret: %s, name: %s\n", *id, *secret, *name)

	cre := credentials.NewStaticCredentials(*id, *secret, "")
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("ap-northeast-1"),
		Credentials: cre,
	})
	if err != nil {
		panic(err)
	}

	e := ecr.New(sess)
	repo, err := ecrf.CreateRepository(e, *name)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Created repository(name: %s)\n", *repo.RepositoryName)

	i := iam.New(sess)
	pull, err := ecrf.CreatePullerPolicy(i, repo)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Created policy(id:%s, name: %s)\n", *pull.PolicyId, *pull.PolicyName)

	push, err := ecrf.CreatePusherPolicy(i, repo)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Created policy(id:%s, name: %s)\n", *push.PolicyId, *push.PolicyName)
}
