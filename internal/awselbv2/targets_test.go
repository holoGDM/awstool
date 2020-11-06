package targets

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elbv2"
)

func TestTakeIds(t *testing.T) {
	// Setup Test

	expect := []string{"i-92734747378274", "i-12341234123556"}

	region := "eu-central-1"
	id1 := "i-92734747378274"
	id2 := "i-12341234123556"
	var port int64 = 80

	testData := &elbv2.DescribeTargetHealthOutput{
		TargetHealthDescriptions: []*elbv2.TargetHealthDescription{
			&elbv2.TargetHealthDescription{
				Target: &elbv2.TargetDescription{
					AvailabilityZone: &region,
					Id:               &id1,
					Port:             &port,
				},
			},
			&elbv2.TargetHealthDescription{
				Target: &elbv2.TargetDescription{
					AvailabilityZone: &region,
					Id:               &id2,
					Port:             &port,
				},
			},
		},
	}

	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String("eu-central-1"),
	})

	target := NewTarget(sess, "arn:aws:elasticloadbalancing:eu-central-1:785702789599:targetgroup/test/c7a44c37a781fe94")

	array := target.TakeIds(testData)

	fmt.Println(array)
	if expect[0] != array[0] {
		t.Fatal("test not passed")
	}

	if expect[1] != array[1] {
		t.Fatal("test not passed")
	}

}
