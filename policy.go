package ecrf

import "github.com/aws/aws-sdk-go/service/iam"

// CreatePolicy create policy
func CreatePolicy(i iam.IAM) (*iam.Policy, error) {
	req := &iam.CreatePolicyInput{}

	resp, err := i.CreatePolicy(req)
	if err != nil {
		return nil, err
	}

	return resp.Policy, nil
}
