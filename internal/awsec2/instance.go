package instance

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// IfInstance interface for instances methods
type IfInstance interface {
	DescribeInstance(string) (*ec2.DescribeInstancesOutput, error)
	TerminateInstance(string) (*ec2.TerminateInstancesOutput, error)
	RunInstance(string) (*ec2.Reservation, error)
}

// Instance struct for instances
type Instance struct {
	session  *session.Session
	instance string
}

// NewInstance constructor for Instance struct
func NewInstance(session *session.Session) *Instance {

	return &Instance{
		session: session,
		// instance: instance,
	}
}

// DescribeInstance method to list instance information, it need instance name as parametr
func (p *Instance) DescribeInstance(instance string) (*ec2.DescribeInstancesOutput, error) {
	svc := ec2.New(p.session)
	input := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String("instance-id"),
				Values: []*string{
					aws.String(instance),
				},
			},
		},
	}

	result, err := svc.DescribeInstances(input)

	return result, err

}

// TerminateInstance method to terminate instance, it need instance name as parametr
func (p *Instance) TerminateInstance(instance string) (*ec2.TerminateInstancesOutput, error) {
	svc := ec2.New(p.session)
	input := &ec2.TerminateInstancesInput{
		InstanceIds: []*string{
			aws.String(instance),
		},
	}

	result, err := svc.TerminateInstances(input)

	return result, err
}

// RunInstance method which is creating new instance of EC2 VMs, it need image parameter
func (p *Instance) RunInstance(image string) (*ec2.Reservation, error) {
	svc := ec2.New(p.session)
	input := &ec2.RunInstancesInput{
		// An Amazon Linux AMI ID for t2.micro instances in the us-west-2 region
		ImageId:      aws.String(image),
		InstanceType: aws.String("t2.micro"),
		MinCount:     aws.Int64(1),
		MaxCount:     aws.Int64(1),
	}

	result, err := svc.RunInstances(input)

	return result, err

}
