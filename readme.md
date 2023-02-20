
# kubetcl 1.0

ubectl and tools for Direktiv.

---
- #### Categories: unknown
- #### Image: gcr.io/direktiv/functions/kubetcl 
- #### License: [Apache-2.0](https://www.apache.org/licenses/LICENSE-2.0)
- #### Issue Tracking: https://github.com/direktiv-apps/kubetcl/issues
- #### URL: https://github.com/direktiv-apps/kubetcl
- #### Maintainer: [direktiv.io](https://www.direktiv.io) 
---

## About kubetcl

This function provides kubectl, kustomize and helm. There following tools are installed:

- kubectl v1.26.1
- helm v3.11.1
- curl
- wget 
- git

The required kubeconfig has to be provided as base64 encoded file.

### Example(s)
  #### Function Configuration
```yaml
functions:
- id: kubetcl
  image: gcr.io/direktiv/functions/kubetcl:1.0
  type: knative-workflow
```
   #### Basic
```yaml
- id: kubectl
  type: action
  action:
    secrets: ["kubeconfig"]
    function: kubectl
    input: 
      kubeconfig: jq(.secrets.kubeconfig | @base64)
      commands:
      - command: kubectl version --client --output json
```
   #### Advanced
```yaml
- id: kubectl
  type: action
  action:
    secrets: ["kubeconfig"]
    function: kubectl
    input: 
      kubeconfig: jq(.secrets.kubeconfig | @base64)
      commands:
      - command: kubectl get pods --output json
  catch:
  - error: "*"
```

   ### Secrets


- **kubeconfig**: Kubeconfig file as BASE64 encoded file for cluster access.






### Request



#### Request Attributes
[PostParamsBody](#post-params-body)

### Response
  List of executed commands.
#### Reponse Types
    
  

[PostOKBody](#post-o-k-body)
#### Example Reponses
    
```json
[
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
]
```
```json
true
```

### Errors
| Type | Description
|------|---------|
| io.direktiv.command.error | Command execution failed |
| io.direktiv.output.error | Template error for output generation of the service |
| io.direktiv.ri.error | Can not create information object from request |


### Types
#### <span id="post-o-k-body"></span> postOKBody

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| kubectl | [][PostOKBodyKubectlItems](#post-o-k-body-kubectl-items)| `[]*PostOKBodyKubectlItems` |  | |  |  |


#### <span id="post-o-k-body-kubectl-items"></span> postOKBodyKubectlItems

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| result | [interface{}](#interface)| `interface{}` | ✓ | |  |  |
| success | boolean| `bool` | ✓ | |  |  |


#### <span id="post-params-body"></span> postParamsBody

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| commands | [][PostParamsBodyCommandsItems](#post-params-body-commands-items)| `[]*PostParamsBodyCommandsItems` |  | `[{"command":"kubectl version"}]`| Array of commands. |  |
| files | [][DirektivFile](#direktiv-file)| `[]apps.DirektivFile` |  | | File to create before running commands. |  |
| kubeconfig | string| `string` |  | | kubeconfig as base64 encoded file |  |


#### <span id="post-params-body-commands-items"></span> postParamsBodyCommandsItems

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| command | string| `string` |  | | Command to run |  |
| continue | boolean| `bool` |  | | Stops excecution if command fails, otherwise proceeds with next command |  |
| print | boolean| `bool` |  | `true`| If set to false the command will not print the full command with arguments to logs. |  |
| silent | boolean| `bool` |  | | If set to false the command will not print output to logs. |  |

 
