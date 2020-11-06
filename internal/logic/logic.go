package logic

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"

	instance "github.com/holoGDM/awstool/internal/awsec2"
	targets "github.com/holoGDM/awstool/internal/awselbv2"
)

//ChangeImagesForInstances creating target object to make orders for target groups
func ChangeImagesForInstances(sess *session.Session, target targets.IfTarget, machineInst instance.IfInstance, firstImage *string, secondImage *string) error {

	fmt.Println("starting DescribeTargets")
	result, err := target.DescribeTargets()
	if err != nil {
		return err
	}
	instances := target.TakeIds(result)

	// Loop for removing and creating instances with new image
	for _, inst := range instances {

		info, err := machineInst.DescribeInstance(inst)
		if err != nil {
			return err
		}

		presentImage := info.Reservations[0].Instances[0].ImageId

		// Check if image need to be changed
		if *presentImage == *firstImage {
			fmt.Println("Starting ami change for instance: ", inst)

			fmt.Println("Deregistring target instance from target group")

			// Remove instance from target group before terminating instance
			target.DeregisterTarget(inst)

			// Terminate old image instance
			machineInst.TerminateInstance(inst)

			// Create new instance with new image
			newMachine, err := machineInst.RunInstance(*secondImage)
			if err != nil {
				return err
			}

			// Take instance id for future usage
			inst = *newMachine.Instances[0].InstanceId
			info, err := machineInst.DescribeInstance(inst)
			if err != nil {
				return err
			}

			// Check state of newly create instance
			state := info.Reservations[0].Instances[0].State.Name

			// Waiting loop for instance be fully available
			fmt.Println("Waiting for instance to be running")
			for *state != "running" {

				info, err := machineInst.DescribeInstance(inst)
				if err != nil {
					return err
				}

				state = info.Reservations[0].Instances[0].State.Name
				fmt.Println(*state)
				time.Sleep(1 * time.Second)
			}
			fmt.Println("Registring new target instance in target group")
			target.RegisterTarget(inst)

			// before we proceed to another instance, here we should add helth status checking
			// for target in target group
			// but i do not have full accesses to AWS and time to create fully working test environment

		}
	}

	return nil
}
