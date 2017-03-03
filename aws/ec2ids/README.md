## Go Packages for interacting with AWS services

This script is designed to print EC2 instance ids to stdout in the format:

tag_value,instance_id
'tag_value','instance_id_1'
'tag_value','instance_id_2'

for all instance IDs assosciated with the AWS account keys stored in ~/.aws/credentials.

This can be used as a splunk lookup to identify specific EC2 instances of interest in log files.

#### Getting Started

Follow [these instructions](https://developers.google.com/identity/protocols/application-default-credentials#howtheywork)
to establish a Google service account and set up the `GOOGLE_APPLICATION_CREDENTIALS`
environment variable with a route to the JSON you created.

Share any services (calendars, analytics, etc...) you want this service to access with the service account address.

Optional: Set up whatever monitoring you need to read the stdout of the machine
you are running from.

Run the service `go run main.go` from the source code (or eventually just from a release, details to follow).

#### License

Copyright ©‎ 2016, Office for National Statistics (https://www.ons.gov.uk)

Released under MIT license, see [LICENSE](LICENSE.md) for details.
