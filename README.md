# go-bitrise

(WIP)

## Usage

### Create a client with your bitrise app slug

```go
var appSlug = "your_bitrise_app_slug"
client := client.NewClient(appSlug)
```

### Trigger a build

```go
// hook_info
hookInfo := bitrise.HookInfo {
    BuildTriggerToken: your_bitrise_build_trigger_token
}

// build_params
buildParams := bitrise.BuildParams {
    Branch: "master",
    WorkflowId: "deploy",
    Tag: "v1.0.0"
}

response, err := client.TriggerBuild(hookInfo, buildParams)
```
