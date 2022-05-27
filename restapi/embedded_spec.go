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
    "description": "Run kubectl commands in Direktiv.",
    "title": "kubectl",
    "version": "1.0",
    "x-direktiv-meta": {
      "categories": [
        "cloud",
        "tools"
      ],
      "container": "direktiv/kubectl",
      "issues": null,
      "license": "[Apache-2.0](https://www.apache.org/licenses/LICENSE-2.0)",
      "long-description": "This function asllows to run kubectl commands in Direktiv. It uses ` + "`" + `kubectl.yaml` + "`" + ` in the working directory for authentication.\nThis file can be provided in the payload as base64 string or in ` + "`" + `files` + "`" + ` as well as via Direktiv files for actions.",
      "maintainer": null,
      "url": null
    }
  },
  "paths": {
    "/": {
      "post": {
        "parameters": [
          {
            "type": "string",
            "description": "direktiv action id is an UUID. \nFor development it can be set to 'development'\n",
            "name": "Direktiv-ActionID",
            "in": "header"
          },
          {
            "type": "string",
            "description": "direktiv temp dir is the working directory for that request\nFor development it can be set to e.g. '/tmp'\n",
            "name": "Direktiv-TempDir",
            "in": "header"
          },
          {
            "name": "body",
            "in": "body",
            "schema": {
              "type": "object",
              "properties": {
                "commands": {
                  "description": "Array of commands.",
                  "type": "array",
                  "items": {
                    "type": "object",
                    "properties": {
                      "command": {
                        "description": "Command to run",
                        "type": "string",
                        "example": "kubectl version --client=true -o json"
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
                  "description": "File to create before running commands. This can include a ` + "`" + `kubectl.yaml` + "`" + ` file from secrets.",
                  "type": "array",
                  "items": {
                    "$ref": "#/definitions/direktivFile"
                  }
                },
                "kubeconfig": {
                  "description": "Base64 kubectl.yaml file. If not set ` + "`" + `kubectl.yaml` + "`" + ` will be used. This can be provided via Direktiv files.",
                  "type": "string",
                  "example": "kubeconfig.yaml"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Return values from commands in the command array.",
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
              "kubectl": {
                "kubectl": [
                  {
                    "result": {
                      "clientVersion": {
                        "buildDate": "2022-05-24T12:26:19Z",
                        "compiler": "gc",
                        "gitCommit": "3ddd0f45aa91e2f30c70734b175631bec5b5825a",
                        "gitTreeState": "clean",
                        "gitVersion": "v1.24.1",
                        "goVersion": "go1.18.2",
                        "major": "1",
                        "minor": "24",
                        "platform": "linux/amd64"
                      },
                      "kustomizeVersion": "v4.5.4"
                    },
                    "success": true
                  }
                ]
              }
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
              "exec": "{{- if empty .Kubeconfig }}\necho \"no kubeconfig in payload\"\n{{- else }}\nbash -c \"echo {{ .Kubeconfig }} | base64 -d \u003e kubectl.yaml\"\n{{- end }}",
              "print": false,
              "silent": true
            },
            {
              "action": "foreach",
              "continue": "{{ .Item.Continue }}",
              "env": [
                "KUBECONFIG={{- .DirektivDir }}/kubectl.yaml"
              ],
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
            "content": "- id: req\n     type: action\n     action:\n       function: kubectl",
            "title": "Basic"
          }
        ],
        "x-direktiv-function": "functions:\n  - id: kubectl\n    image: direktiv/kubectl:1.0\n    type: knative-workflow"
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
    "description": "Run kubectl commands in Direktiv.",
    "title": "kubectl",
    "version": "1.0",
    "x-direktiv-meta": {
      "categories": [
        "cloud",
        "tools"
      ],
      "container": "direktiv/kubectl",
      "issues": null,
      "license": "[Apache-2.0](https://www.apache.org/licenses/LICENSE-2.0)",
      "long-description": "This function asllows to run kubectl commands in Direktiv. It uses ` + "`" + `kubectl.yaml` + "`" + ` in the working directory for authentication.\nThis file can be provided in the payload as base64 string or in ` + "`" + `files` + "`" + ` as well as via Direktiv files for actions.",
      "maintainer": null,
      "url": null
    }
  },
  "paths": {
    "/": {
      "post": {
        "parameters": [
          {
            "type": "string",
            "description": "direktiv action id is an UUID. \nFor development it can be set to 'development'\n",
            "name": "Direktiv-ActionID",
            "in": "header"
          },
          {
            "type": "string",
            "description": "direktiv temp dir is the working directory for that request\nFor development it can be set to e.g. '/tmp'\n",
            "name": "Direktiv-TempDir",
            "in": "header"
          },
          {
            "name": "body",
            "in": "body",
            "schema": {
              "type": "object",
              "properties": {
                "commands": {
                  "description": "Array of commands.",
                  "type": "array",
                  "items": {
                    "$ref": "#/definitions/CommandsItems0"
                  }
                },
                "files": {
                  "description": "File to create before running commands. This can include a ` + "`" + `kubectl.yaml` + "`" + ` file from secrets.",
                  "type": "array",
                  "items": {
                    "$ref": "#/definitions/direktivFile"
                  }
                },
                "kubeconfig": {
                  "description": "Base64 kubectl.yaml file. If not set ` + "`" + `kubectl.yaml` + "`" + ` will be used. This can be provided via Direktiv files.",
                  "type": "string",
                  "example": "kubeconfig.yaml"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Return values from commands in the command array.",
            "schema": {
              "type": "object",
              "properties": {
                "kubectl": {
                  "type": "array",
                  "items": {
                    "$ref": "#/definitions/KubectlItems0"
                  }
                }
              }
            },
            "examples": {
              "kubectl": {
                "kubectl": [
                  {
                    "result": {
                      "clientVersion": {
                        "buildDate": "2022-05-24T12:26:19Z",
                        "compiler": "gc",
                        "gitCommit": "3ddd0f45aa91e2f30c70734b175631bec5b5825a",
                        "gitTreeState": "clean",
                        "gitVersion": "v1.24.1",
                        "goVersion": "go1.18.2",
                        "major": "1",
                        "minor": "24",
                        "platform": "linux/amd64"
                      },
                      "kustomizeVersion": "v4.5.4"
                    },
                    "success": true
                  }
                ]
              }
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
              "exec": "{{- if empty .Kubeconfig }}\necho \"no kubeconfig in payload\"\n{{- else }}\nbash -c \"echo {{ .Kubeconfig }} | base64 -d \u003e kubectl.yaml\"\n{{- end }}",
              "print": false,
              "silent": true
            },
            {
              "action": "foreach",
              "continue": "{{ .Item.Continue }}",
              "env": [
                "KUBECONFIG={{- .DirektivDir }}/kubectl.yaml"
              ],
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
            "content": "- id: req\n     type: action\n     action:\n       function: kubectl",
            "title": "Basic"
          }
        ],
        "x-direktiv-function": "functions:\n  - id: kubectl\n    image: direktiv/kubectl:1.0\n    type: knative-workflow"
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
    "CommandsItems0": {
      "type": "object",
      "properties": {
        "command": {
          "description": "Command to run",
          "type": "string",
          "example": "kubectl version --client=true -o json"
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
    },
    "KubectlItems0": {
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
    },
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
}
