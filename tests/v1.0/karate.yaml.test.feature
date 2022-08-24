
Feature: Basic

# The secrects can be used in the payload with the following syntax #(mysecretname)
Background:
* def kubeconfig = karate.properties['kubeconfig']


Scenario:  version

	Given url karate.properties['testURL']

	And path '/'
	And header Direktiv-ActionID = 'development'
	And header Direktiv-TempDir = '/tmp'
	And request
	"""
	{
		"kubectl": #(kubeconfig),
		"commands": [
		{
			"command": "kubectl version --client --output json",
			"silent": true,
			"print": false,
		}
		]
	}
	"""
	When method POST
	Then status 200
	And match $ ==
	"""
	{
	"kubectl": [
	{
		"result": {
			"clientVersion": #notnull,
			"kustomizeVersion": #notnull
		},
		"success": true
	}
	]
	}
	"""
	
Scenario: pods

	Given url karate.properties['testURL']

	And path '/'
	And header Direktiv-ActionID = 'development'
	And header Direktiv-TempDir = '/tmp'
	And request
	"""
	{
		"kubectl": #(kubeconfig),
		"commands": [
		{
			"command": "kubectl get pods --output json",
			"silent": true,
			"print": false,
		}
		]
	}
	"""
	When method POST
	Then status 500
	