package logic

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/elbv2"
	instance "github.com/holoGDM/awstool/internal/awsec2"
	targets "github.com/holoGDM/awstool/internal/awselbv2"
)

type mockTargets struct {
	targets.IfTarget
}

type mockInstance struct {
	instance.IfInstance
}

func (p *mockTargets) RegisterTarget(string) (*elbv2.RegisterTargetsOutput, error) {
	return &elbv2.RegisterTargetsOutput{}, nil
}

func (p *mockTargets) DeregisterTarget(string) (*elbv2.DeregisterTargetsOutput, error) {
	return &elbv2.DeregisterTargetsOutput{}, nil
}

func (p *mockTargets) TakeIds(result *elbv2.DescribeTargetHealthOutput) []string {

	var ids []string
	for _, id := range result.TargetHealthDescriptions {
		fmt.Println("jestem w for")
		fmt.Println(id)
		ids = append(ids, *id.Target.Id)
	}

	return ids
}

func (p *mockTargets) DescribeTargets() (*elbv2.DescribeTargetHealthOutput, error) {

	port := "80"
	var portInt64 int64 = 80
	zone := "eu-central-1b"
	instance1 := "i-123412349895858"
	instance2 := "i-747474747474747"

	description1 := "Test target health for instance1"
	description2 := "Test target health for instance2"
	reason := "Everything is fine"
	state := "Healthy"

	describeTargetsHealthOutput := &elbv2.DescribeTargetHealthOutput{
		TargetHealthDescriptions: []*elbv2.TargetHealthDescription{
			&elbv2.TargetHealthDescription{
				HealthCheckPort: &port,
				Target: &elbv2.TargetDescription{
					AvailabilityZone: &zone,
					Id:               &instance1,
					Port:             &portInt64,
				},
				TargetHealth: &elbv2.TargetHealth{
					Description: &description1,
					Reason:      &reason,
					State:       &state,
				},
			},
			&elbv2.TargetHealthDescription{
				HealthCheckPort: &port,
				Target: &elbv2.TargetDescription{
					AvailabilityZone: &zone,
					Id:               &instance2,
					Port:             &portInt64,
				},
				TargetHealth: &elbv2.TargetHealth{
					Description: &description2,
					Reason:      &reason,
					State:       &state,
				},
			},
		},
	}

	fmt.Println("jestem w outpcie\n", describeTargetsHealthOutput)

	return describeTargetsHealthOutput, nil
}

func (p *mockInstance) DescribeInstance(string) (*ec2.DescribeInstancesOutput, error) {

	nextToken := "l1kj234klj21h34klj12kl3j41klj234klj1h"
	groupID := "g-12341234"
	groupName := "test-group"
	imageID := "ami-98127349182734"
	var amiLauncherIndex int64 = 0
	architecture := "x86_64"
	instanceID := "i-87329182734981273"

	running := "running"
	var code int64 = 124
	state := &ec2.InstanceState{
		Name: &running,
		Code: &code,
	}

	groups := []*ec2.GroupIdentifier{
		&ec2.GroupIdentifier{
			GroupId:   &groupID,
			GroupName: &groupName,
		}}

	instances := []*ec2.Instance{
		&ec2.Instance{
			AmiLaunchIndex: &amiLauncherIndex,
			Architecture:   &architecture,
			ImageId:        &imageID,
			InstanceId:     &instanceID,
			State:          state,
		},
	}

	reservations := []*ec2.Reservation{
		&ec2.Reservation{
			Groups:    groups,
			Instances: instances,
		},
	}

	return &ec2.DescribeInstancesOutput{
		NextToken:    &nextToken,
		Reservations: reservations,
	}, nil

}

func (p *mockInstance) TerminateInstance(instance string) (*ec2.TerminateInstancesOutput, error) {

	terminating := "Terminating"
	running := "Running"

	var codeInt64 int64 = 123

	return &ec2.TerminateInstancesOutput{
		TerminatingInstances: []*ec2.InstanceStateChange{
			&ec2.InstanceStateChange{
				CurrentState: &ec2.InstanceState{
					Code: &codeInt64,
					Name: &terminating,
				},
				PreviousState: &ec2.InstanceState{
					Code: &codeInt64,
					Name: &running,
				},
				InstanceId: &instance,
			},
		},
	}, nil
}

func (p *mockInstance) RunInstance(string) (*ec2.Reservation, error) {

	groupID := "g-12341234"
	groupName := "test-group"
	imageID := "ami-98127349182734"
	var amiLauncherIndex int64 = 0
	architecture := "x86_64"
	instanceID := "i-87329182734981273"

	running := "running"
	var code int64 = 124
	state := &ec2.InstanceState{
		Name: &running,
		Code: &code,
	}

	groups := []*ec2.GroupIdentifier{
		&ec2.GroupIdentifier{
			GroupId:   &groupID,
			GroupName: &groupName,
		}}

	instances := []*ec2.Instance{
		&ec2.Instance{
			AmiLaunchIndex: &amiLauncherIndex,
			Architecture:   &architecture,
			ImageId:        &imageID,
			InstanceId:     &instanceID,
			State:          state,
		},
	}

	reservation := &ec2.Reservation{
		Groups:    groups,
		Instances: instances,
	}

	return reservation, nil

}

func TestChangeImagesForInstances(t *testing.T) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("eu-central-1"),
	})
	if err != nil {
		t.Fatal(err)
	}

	first := "ami-281348182348"
	second := "ami-29349218342"

	trg := &mockTargets{}
	instance := &mockInstance{}

	err = ChangeImagesForInstances(sess, trg, instance, &first, &second)
	if err != nil {
		t.Fatal(err)
	}
}
