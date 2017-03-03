package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func main() {

	// Command line arguments, which will be compared to the EC2
	// instance tags values.
	args := os.Args[1:]

	sess, err := session.NewSession()
	if err != nil {
		panic(err)
	}

	// Create an EC2 service object in the "us-west-2" region
	// Note that you can also configure your region globally by
	// exporting the AWS_REGION environment variable
	svc := ec2.New(sess, &aws.Config{Region: aws.String("eu-west-1")})

	// Call the DescribeInstances Operation
	resp, err := svc.DescribeInstances(nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("tag_value,instance_id")
	// resp has all of the response data, pull out instance IDs:
	for idx := range resp.Reservations {

		// Instances can be iterated over from within the response
		for _, inst := range resp.Reservations[idx].Instances {
			// Only check ids match if args supplied
			if len(args) > 0 {
				// Each instance can have 0, 1 or multiple tags
				for _, tag := range inst.Tags {
					tag_value := *tag.Value
					if stringInSlice(tag_value, args) {
						fmt.Printf("%v,%v\n", tag_value, inst.InstanceId)
					}

				}
			} else {
				fmt.Printf("%v,%v\n", "", inst.InstanceId)
			}
		}
	}
}
