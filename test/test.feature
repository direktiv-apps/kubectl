Feature: greeting end-point

Background:
* url demoBaseUrl

Scenario: version

    Given path '/'
    Given header Direktiv-ActionID = 'development'
    Given header Direktiv-Tempdir = '/tmp'
    And request 
    """
    { 
        "commands": [
            {
                "command": "kubectl version --client=true -o json"
            }
        ]
    }
    """
    When method post
    Then status 200