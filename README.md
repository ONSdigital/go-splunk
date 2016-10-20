## Google Calendar - JSON Event Builder

This script is designed to authorize with Google to read Calendar events and
convert the current day's events into JSON. The intention behind it is to
generate a log file which can be used in a log aggregator or event monitoring
service.

A specific use for this may be to overlay maintenance windows or a support
schedule into a monitoring service such as Splunk.

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
