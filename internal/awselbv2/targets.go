package targets

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elbv2"
)

// IfTarget Interface for load balancer node group methods
type IfTarget interface {
	RegisterTarget(string) (*elbv2.RegisterTargetsOutput, error)
	DeregisterTarget(string) (*elbv2.DeregisterTargetsOutput, error)
	DescribeTargets() (*elbv2.DescribeTargetHealthOutput, error)
	TakeIds(result *elbv2.DescribeTargetHealthOutput) []string
}

// Target Struct for load balancer targets
type Target struct {
	session *session.Session
	arn     string
}

// NewTarget Constructor of load balancer targets
func NewTarget(session *session.Session, arn string) *Target {

	return &Target{
		session: session,
		arn:     arn,
	}
}

// RegisterTarget Add target to target group of load balancer
func (p *Target) RegisterTarget(instance string) (*elbv2.RegisterTargetsOutput, error) {

	svc := elbv2.New(p.session)
	input := &elbv2.RegisterTargetsInput{
		TargetGroupArn: aws.String(p.arn),
		Targets: []*elbv2.TargetDescription{
			{
				Id: aws.String(instance),
			},
		},
	}

	result, err := svc.RegisterTargets(input)

	return result, err
}

// DeregisterTarget Remove target from target group of load balancer
func (p *Target) DeregisterTarget(instance string) (*elbv2.DeregisterTargetsOutput, error) {
	svc := elbv2.New(p.session)
	input := &elbv2.DeregisterTargetsInput{
		TargetGroupArn: aws.String(p.arn),
		Targets: []*elbv2.TargetDescription{
			{
				Id: aws.String(instance),
			},
		},
	}

	result, err := svc.DeregisterTargets(input)

	return result, err
}

// DescribeTargets List informations about available targets in target group
func (p *Target) DescribeTargets() (*elbv2.DescribeTargetHealthOutput, error) {
	svc := elbv2.New(p.session)
	input := &elbv2.DescribeTargetHealthInput{
		TargetGroupArn: aws.String(p.arn),
	}

	result, err := svc.DescribeTargetHealth(input)

	return result, err
}

// TakeIds grab target instances ids to array of strings
func (p *Target) TakeIds(result *elbv2.DescribeTargetHealthOutput) []string {

	var ids []string
	for _, id := range result.TargetHealthDescriptions {

		ids = append(ids, *id.Target.Id)
	}

	return ids
}
