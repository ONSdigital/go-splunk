package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

var (
	Error *log.Logger
)

func log_handler(errorHandle io.Writer) {

	Error = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func main() {

	log_handler(os.Stdout)

	// Command line arguments, which will be compared to the EC2
	// instance tags values.
	args := os.Args[1:]

	if len(args) == 0 {
		Error.Println("No tag value supplied")
		os.Exit(1)
	}

	// If error, write to STDOUT and exit(1)
	sess, err := session.NewSession()
	if err != nil {
		Error.Println("Error initialising new AWS session")
		os.Exit(1)
	}

	// Create an EC2 service object in the "eu-west-1" region
	// Note that you can also configure your region globally by
	// exporting the AWS_REGION environment variable
	svc := ec2.New(sess)

	// Call the DescribeInstances Operation
	resp, err := svc.DescribeInstances(nil)
	if err != nil {
		os.Exit(1)
	}

	fmt.Println("tag_value,instance_id")

	// resp has all of the response data, pull out instance IDs:
	for idx := range resp.Reservations {
		// Instances can be iterated over from within the response
		for _, inst := range resp.Reservations[idx].Instances {
			// Each instance can have 0, 1 or multiple tags
			for _, tag := range inst.Tags {
				tag_value := *tag.Value
				if stringInSlice(tag_value, args) {
					fmt.Printf("%v,%v\n", tag_value, inst.InstanceId)
				}
			}
		}
	}
}
