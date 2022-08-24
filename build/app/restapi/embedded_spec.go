// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Kubectl and tools for Direktiv.",
    "title": "kubectl",
    "version": "1.0",
    "x-direktiv-meta": {
      "categories": [
        "build"
      ],
      "container": "direktiv.azurecr.io/functions/kubectl",
      "issues": "https://github.com/direktiv-apps/kubectl/issues",
      "license": "[Apache-2.0](https://www.apache.org/licenses/LICENSE-2.0)",
      "long-description": "This function provides kubectl, kustomize and helm. There following tools are installed:\n\n- kubectl 1.25\n- helm v3.9.3\n- curl\n- wget \n\nThe required kubeconfig has to be provided as base64 encoded file.",
      "maintainer": "[direktiv.io](https://www.direktiv.io) ",
      "url": "https://github.com/direktiv-apps/kubectl"
    }
  },
  "paths": {
    "/": {
      "post": {
        "parameters": [
          {
            "type": "string",
            "default": "development",
            "description": "direktiv action id is an UUID. \nFor development it can be set to 'development'\n",
            "name": "Direktiv-ActionID",
            "in": "header"
          },
          {
            "type": "string",
            "default": "/tmp",
            "description": "direktiv temp dir is the working directory for that request\nFor development it can be set to e.g. '/tmp'\n",
            "name": "Direktiv-TempDir",
            "in": "header"
          },
          {
            "name": "body",
            "in": "body",
            "schema": {
              "type": "object",
              "required": [
                "kubectl"
              ],
              "properties": {
                "commands": {
                  "description": "Array of commands.",
                  "type": "array",
                  "default": [
                    {
                      "command": "echo Hello"
                    }
                  ],
                  "items": {
                    "type": "object",
                    "properties": {
                      "command": {
                        "description": "Command to run",
                        "type": "string"
                      },
                      "continue": {
                        "description": "Stops excecution if command fails, otherwise proceeds with next command",
                        "type": "boolean"
                      },
                      "print": {
                        "description": "If set to false the command will not print the full command with arguments to logs.",
                        "type": "boolean",
                        "default": true
                      },
                      "silent": {
                        "description": "If set to false the command will not print output to logs.",
                        "type": "boolean",
                        "default": false
                      }
                    }
                  }
                },
                "files": {
                  "description": "File to create before running commands.",
                  "type": "array",
                  "default": null,
                  "items": {
                    "$ref": "#/definitions/direktivFile"
                  }
                },
                "kubectl": {
                  "description": "kubeconfig as base64 encoded file",
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "List of executed commands.",
            "schema": {
              "type": "object",
              "properties": {
                "kubectl": {
                  "type": "array",
                  "items": {
                    "type": "object",
                    "required": [
                      "success",
                      "result"
                    ],
                    "properties": {
                      "result": {
                        "additionalProperties": false
                      },
                      "success": {
                        "type": "boolean"
                      }
                    }
                  }
                }
              }
            },
            "examples": {
              "kubectl": [
                {
                  "clientVersion": {
                    "buildDate": "2022-08-23T17:44:59Z",
                    "compiler": "gc",
                    "gitCommit": "a866cbe2e5bbaa01cfd5e969aa3e033f3282a8a2",
                    "gitTreeState": "clean",
                    "gitVersion": "v1.25.0",
                    "goVersion": "go1.19",
                    "major": "1",
                    "minor": "25",
                    "platform": "linux/amd64"
                  },
                  "kustomizeVersion": "v4.5.7",
                  "result": null
                }
              ],
              "success": true
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            },
            "headers": {
              "Direktiv-ErrorCode": {
                "type": "string"
              },
              "Direktiv-ErrorMessage": {
                "type": "string"
              }
            }
          }
        },
        "x-direktiv": {
          "cmds": [
            {
              "action": "exec",
              "exec": "bash -c 'mkdir -p ~/.kube/ \u0026\u0026 echo {{ .Kubectl }} | base64 -d \u003e ~/.kube/config'",
              "print": false,
              "silent": true
            },
            {
              "action": "foreach",
              "continue": "{{ .Item.Continue }}",
              "exec": "{{ .Item.Command }}",
              "loop": ".Commands",
              "print": "{{ .Item.Print }}",
              "silent": "{{ .Item.Silent }}"
            }
          ],
          "output": "{\n  \"kubectl\": {{ index . 1 | toJson }}\n}\n"
        },
        "x-direktiv-errors": {
          "io.direktiv.command.error": "Command execution failed",
          "io.direktiv.output.error": "Template error for output generation of the service",
          "io.direktiv.ri.error": "Can not create information object from request"
        },
        "x-direktiv-examples": [
          {
            "content": "- id: kubectl\n  type: action\n  action:\n    secrets: [\"kubeconfig\"]\n    function: kubectl\n    input: \n      kubeconfig: jq(.secrets.kubeconfig | @base64)\n      commands:\n      - command: kubectl version --client --output json",
            "title": "Version"
          },
          {
            "content": "- id: kubectl\n  type: action\n  action:\n    secrets: [\"kubeconfig\"]\n    function: kubectl\n    input: \n      kubeconfig: jq(.secrets.kubeconfig | @base64)\n      commands:\n      - command: kubectl get pods --output json\n  catch:\n  - error: \"*\"",
            "title": "Pods"
          }
        ],
        "x-direktiv-function": "functions:\n- id: kubectl\n  image: direktiv.azurecr.io/functions/kubectl:1.0\n  type: knative-workflow",
        "x-direktiv-secrets": [
          {
            "description": "Kubeconfig file as BASE64 encoded file for cluster access.",
            "name": "kubeconfig"
          }
        ]
      },
      "delete": {
        "parameters": [
          {
            "type": "string",
            "description": "On cancel Direktiv sends a DELETE request to\nthe action with id in the header\n",
            "name": "Direktiv-ActionID",
            "in": "header"
          }
        ],
        "responses": {
          "200": {
            "description": ""
          }
        },
        "x-direktiv": {
          "cancel": "echo 'cancel {{ .DirektivActionID }}'"
        }
      }
    }
  },
  "definitions": {
    "direktivFile": {
      "type": "object",
      "x-go-type": {
        "import": {
          "package": "github.com/direktiv/apps/go/pkg/apps"
        },
        "type": "DirektivFile"
      }
    },
    "error": {
      "type": "object",
      "required": [
        "errorCode",
        "errorMessage"
      ],
      "properties": {
        "errorCode": {
          "type": "string"
        },
        "errorMessage": {
          "type": "string"
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Kubectl and tools for Direktiv.",
    "title": "kubectl",
    "version": "1.0",
    "x-direktiv-meta": {
      "categories": [
        "build"
      ],
      "container": "direktiv.azurecr.io/functions/kubectl",
      "issues": "https://github.com/direktiv-apps/kubectl/issues",
      "license": "[Apache-2.0](https://www.apache.org/licenses/LICENSE-2.0)",
      "long-description": "This function provides kubectl, kustomize and helm. There following tools are installed:\n\n- kubectl 1.25\n- helm v3.9.3\n- curl\n- wget \n\nThe required kubeconfig has to be provided as base64 encoded file.",
      "maintainer": "[direktiv.io](https://www.direktiv.io) ",
      "url": "https://github.com/direktiv-apps/kubectl"
    }
  },
  "paths": {
    "/": {
      "post": {
        "parameters": [
          {
            "type": "string",
            "default": "development",
            "description": "direktiv action id is an UUID. \nFor development it can be set to 'development'\n",
            "name": "Direktiv-ActionID",
            "in": "header"
          },
          {
            "type": "string",
            "default": "/tmp",
            "description": "direktiv temp dir is the working directory for that request\nFor development it can be set to e.g. '/tmp'\n",
            "name": "Direktiv-TempDir",
            "in": "header"
          },
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/postParamsBody"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "List of executed commands.",
            "schema": {
              "$ref": "#/definitions/postOKBody"
            },
            "examples": {
              "kubectl": [
                {
                  "clientVersion": {
                    "buildDate": "2022-08-23T17:44:59Z",
                    "compiler": "gc",
                    "gitCommit": "a866cbe2e5bbaa01cfd5e969aa3e033f3282a8a2",
                    "gitTreeState": "clean",
                    "gitVersion": "v1.25.0",
                    "goVersion": "go1.19",
                    "major": "1",
                    "minor": "25",
                    "platform": "linux/amd64"
                  },
                  "kustomizeVersion": "v4.5.7",
                  "result": null
                }
              ],
              "success": true
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            },
            "headers": {
              "Direktiv-ErrorCode": {
                "type": "string"
              },
              "Direktiv-ErrorMessage": {
                "type": "string"
              }
            }
          }
        },
        "x-direktiv": {
          "cmds": [
            {
              "action": "exec",
              "exec": "bash -c 'mkdir -p ~/.kube/ \u0026\u0026 echo {{ .Kubectl }} | base64 -d \u003e ~/.kube/config'",
              "print": false,
              "silent": true
            },
            {
              "action": "foreach",
              "continue": "{{ .Item.Continue }}",
              "exec": "{{ .Item.Command }}",
              "loop": ".Commands",
              "print": "{{ .Item.Print }}",
              "silent": "{{ .Item.Silent }}"
            }
          ],
          "output": "{\n  \"kubectl\": {{ index . 1 | toJson }}\n}\n"
        },
        "x-direktiv-errors": {
          "io.direktiv.command.error": "Command execution failed",
          "io.direktiv.output.error": "Template error for output generation of the service",
          "io.direktiv.ri.error": "Can not create information object from request"
        },
        "x-direktiv-examples": [
          {
            "content": "- id: kubectl\n  type: action\n  action:\n    secrets: [\"kubeconfig\"]\n    function: kubectl\n    input: \n      kubeconfig: jq(.secrets.kubeconfig | @base64)\n      commands:\n      - command: kubectl version --client --output json",
            "title": "Version"
          },
          {
            "content": "- id: kubectl\n  type: action\n  action:\n    secrets: [\"kubeconfig\"]\n    function: kubectl\n    input: \n      kubeconfig: jq(.secrets.kubeconfig | @base64)\n      commands:\n      - command: kubectl get pods --output json\n  catch:\n  - error: \"*\"",
            "title": "Pods"
          }
        ],
        "x-direktiv-function": "functions:\n- id: kubectl\n  image: direktiv.azurecr.io/functions/kubectl:1.0\n  type: knative-workflow",
        "x-direktiv-secrets": [
          {
            "description": "Kubeconfig file as BASE64 encoded file for cluster access.",
            "name": "kubeconfig"
          }
        ]
      },
      "delete": {
        "parameters": [
          {
            "type": "string",
            "description": "On cancel Direktiv sends a DELETE request to\nthe action with id in the header\n",
            "name": "Direktiv-ActionID",
            "in": "header"
          }
        ],
        "responses": {
          "200": {
            "description": ""
          }
        },
        "x-direktiv": {
          "cancel": "echo 'cancel {{ .DirektivActionID }}'"
        }
      }
    }
  },
  "definitions": {
    "direktivFile": {
      "type": "object",
      "x-go-type": {
        "import": {
          "package": "github.com/direktiv/apps/go/pkg/apps"
        },
        "type": "DirektivFile"
      }
    },
    "error": {
      "type": "object",
      "required": [
        "errorCode",
        "errorMessage"
      ],
      "properties": {
        "errorCode": {
          "type": "string"
        },
        "errorMessage": {
          "type": "string"
        }
      }
    },
    "postOKBody": {
      "type": "object",
      "properties": {
        "kubectl": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/postOKBodyKubectlItems"
          }
        }
      },
      "x-go-gen-location": "operations"
    },
    "postOKBodyKubectlItems": {
      "type": "object",
      "required": [
        "success",
        "result"
      ],
      "properties": {
        "result": {
          "additionalProperties": false
        },
        "success": {
          "type": "boolean"
        }
      },
      "x-go-gen-location": "operations"
    },
    "postParamsBody": {
      "type": "object",
      "required": [
        "kubectl"
      ],
      "properties": {
        "commands": {
          "description": "Array of commands.",
          "type": "array",
          "default": [
            {
              "command": "echo Hello"
            }
          ],
          "items": {
            "$ref": "#/definitions/postParamsBodyCommandsItems"
          }
        },
        "files": {
          "description": "File to create before running commands.",
          "type": "array",
          "default": [],
          "items": {
            "$ref": "#/definitions/direktivFile"
          }
        },
        "kubectl": {
          "description": "kubeconfig as base64 encoded file",
          "type": "string"
        }
      },
      "x-go-gen-location": "operations"
    },
    "postParamsBodyCommandsItems": {
      "type": "object",
      "properties": {
        "command": {
          "description": "Command to run",
          "type": "string"
        },
        "continue": {
          "description": "Stops excecution if command fails, otherwise proceeds with next command",
          "type": "boolean"
        },
        "print": {
          "description": "If set to false the command will not print the full command with arguments to logs.",
          "type": "boolean",
          "default": true
        },
        "silent": {
          "description": "If set to false the command will not print output to logs.",
          "type": "boolean",
          "default": false
        }
      },
      "x-go-gen-location": "operations"
    }
  }
}`))
}
