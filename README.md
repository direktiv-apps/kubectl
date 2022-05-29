
# kubectl 1.0

Run kubectl commands in Direktiv.

---
- #### Categories: cloud, tools
- #### Image: gcr.io/direktiv/apps/kubectl 
- #### License: [Apache-2.0](https://www.apache.org/licenses/LICENSE-2.0)
- #### Issue Tracking: https://github.com/direktiv-apps/kubectl/issues
- #### URL: https://github.com/direktiv-apps/kubectl
- #### Maintainer: [direktiv.io](https://www.direktiv.io)
---

## About kubectl

This function allows to run kubectl commands in Direktiv. It uses `kubectl.yaml` in the working directory for authentication.
This file can be provided in the payload as base64 string or in `files` as well as via Direktiv files for actions.

### Example(s)
  #### Function Configuration
  ```yaml
  functions:
  - id: kubectl
    image: gcr.io/direktiv/apps/kubectl:1.0
    type: knative-workflow
  ```
   #### Kubectl with Secrets
   ```yaml
   - id: kubectl 
      type: action
      action:
        secrets: ["kubectl"]
        function: get
        input: 
          kubeconfig: jq(.secrets.kubectl | @base64)
          commands:
          - command: kubectl --insecure-skip-tls-verify --server=https://myserver:6443/ -o json get pods
   ```
   #### Kubectl with Variable
   ```yaml
   - id: get-kubeconfig
      type: getter
      variables:
      - key: k3s.yaml
        scope: workflow
      transition: kubectl
    - id: kubectl 
      type: action
      action:
        secrets: ["k3s"]
        function: get
        input: 
          kubeconfig: jq(.var."k3s.yaml")
          commands:
          - command: kubectl -o json get pods
   ```
   #### Kubectl with Direktiv File
   ```yaml
   - id: kubectl 
      type: action
      action:
        secrets: ["k3s"]
        function: get
        files:
        - key: k3s.yaml
          scope: workflow
          as: kubectl.yaml
        input: 
          kubeconfig: jq(.secrets.k3s | @base64)
          commands:
          - command: kubectl --server=https://myserver:6443/ -o json get pods
   ```

### Request



#### Request Attributes
[PostParamsBody](#post-params-body)

### Response
  Return values from commands in the command array.
#### Reponse Types
    
  

[PostOKBody](#post-o-k-body)
#### Example Reponses
    
```json
{
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
| commands | [][PostParamsBodyCommandsItems](#post-params-body-commands-items)| `[]*PostParamsBodyCommandsItems` |  | | Array of commands. |  |
| files | [][DirektivFile](#direktiv-file)| `[]apps.DirektivFile` |  | | File to create before running commands. This can include a `kubectl.yaml` file from secrets. |  |
| kubeconfig | string| `string` |  | | Base64 kubectl.yaml file. If not set `kubectl.yaml` will be used. This can be provided via Direktiv files. | `tLS0tCk1IY0NBUUVFSUlQN...Fa0luUW1ZbGovY0lIbjQwakZ1eUUxe` |


#### <span id="post-params-body-commands-items"></span> postParamsBodyCommandsItems

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| command | string| `string` |  | | Command to run | `kubectl version --client=true -o json` |
| continue | boolean| `bool` |  | | Stops excecution if command fails, otherwise proceeds with next command |  |
| print | boolean| `bool` |  | `true`| If set to false the command will not print the full command with arguments to logs. |  |
| silent | boolean| `bool` |  | | If set to false the command will not print output to logs. |  |

 
