## Go Packages for interacting with AWS services

This script is designed to print EC2 instance ids to stdout in the format:

tag_value,instance_id
'tag_value','instance_id_1'
'tag_value','instance_id_2'

for all instance IDs assosciated with the AWS account keys stored in ~/.aws/credentials.

This can be used as a splunk lookup to identify specific EC2 instances of interest in log files.

#### Getting Started

Create a config file with your AWS access keys as per [these instructions](http://docs.aws.amazon.com/cli/latest/userguide/cli-chap-getting-started.html#cli-multiple-profiles).

Run the service `go run ec2ids.go` from the source code.

#### License

Copyright ©‎ 2017, Office for National Statistics (https://www.ons.gov.uk)

Released under MIT license, see [LICENSE](LICENSE.md) for details.
