package aws

import (
	"log"
	"os"
)

var customEndpoints map[string]string

// SetAwsEndpointsOverrides function to override the aws endpoints globally
// this was configured to only happen one time, but why?  Other than being "confusing"
// I don't think that is an actual problem
func SetAwsEndpointsOverrides(endpoints map[string]string) {
	if HasAwsCustomEndPoints() {
		log.Println("SetAwsEndpointsOverrides() has already been called once before.")
		os.Exit(0)
	}
	customEndpoints = endpoints
}

// HasAwsCustomEndPoints Function to verify if already has been setup the custom endpoints
func HasAwsCustomEndPoints() bool {
	return len(customEndpoints) > 0
}

// EndpointShouldBeOverride fucntion check if the current service evaulate has custom endpoint to be override
func EndpointShouldBeOverride(service string) (string, bool) {
	customServiceEndpoint, endpointShouldBeOverride := customEndpoints[service]

	return customServiceEndpoint, endpointShouldBeOverride
}
