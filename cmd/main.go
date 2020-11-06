package main

import (
	"flag"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"

	instance "github.com/holoGDM/awstool/internal/awsec2"
	targets "github.com/holoGDM/awstool/internal/awselbv2"
	"github.com/holoGDM/awstool/internal/logic"
)

// Const variables in which we are setting our arn of our target group and region which we want to use
const (
	arnTargetgroup = "arn:aws:elasticloadbalancing:eu-central-1:785702789599:targetgroup/test/c7a44c37a781fe94"
	region         = "eu-central-1"
)

// Error checking function
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// Our program arguments/flags we can add more of them bellow
	firstImage := flag.String("firstImage", "ami-0c960b947cbb2dd16", "image which we want to change")
	secondImage := flag.String("secondImage", "ami-0ed0be684d3f014bf", "image to which we want to change")

	flag.Parse()

	//Session creation for AWS sdk
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	check(err)

	trg := targets.NewTarget(sess, arnTargetgroup)
	ins := instance.NewInstance(sess)
	err = logic.ChangeImagesForInstances(sess, trg, ins, firstImage, secondImage)
	check(err)

}
