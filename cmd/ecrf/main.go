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

	i := iam.New(sess)
	pull, err := ecrf.CreatePullerPolicy(i, repo)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Create policy(id:%s, name: %s)", *pull.PolicyId, *pull.PolicyName)

	push, err := ecrf.CreatePusherPolicy(i, repo)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Create policy(id:%s, name: %s)", *push.PolicyId, *push.PolicyName)
}
