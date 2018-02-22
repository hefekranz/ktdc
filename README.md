# ktdc

Converter // Kubernetes -> Docker-Compose v2


! Work in Progress !

So far only a few fields and file types implemented.
Maybe helpful for some Devs out there.

The Idea is to quickly translate kubernetes config files into docker-compose.
So far it only grabs a view fields from kubernetes deployment.yaml and creates a docker-compose service from this.
 
## Usage
```bash
ktdc --config=config.json --dcFile=init-compose.yaml --output=merged-compose.yaml
```

### --config
Minimum config
```json
{
  "services": {
    "myService": {
      "deploymentFile": "path/to/deployment.yaml, required!"
    }
  }
}
```


Full config
```json
{
  "services": {
    "myService": {
      "deploymentFile": "path/to/deployment.yaml, required!",
      "ports": [
        "80:9000",
        "8080:3000"
      ],
      "envs": {
        "KEY":"value",
        "AND":"so on"
      }
    },
    "globalEnvs": {
        "KEY":"value"
    }
  }
}
```
_deploymentFile_ is the only required field here.

_globalEnvs_ will be set for every service.
_service.envs_ will only be set for the specified service
 see [Config Hierarchy](#config-hierarchy)
 
### --dcFile
An initial docker-compose.yaml to add to. If you to give this flag it will just create a new file.
**Only a few field will be picked up so far!**
 - image
 - container_name
 - ports
 - environment 
 - logging.driver
 - command
 
### --output

File name of created compose file, defaults to docker-compose.yaml

#### Config Hierarchy

 kubernetes deployment < globalEnvs (config.json) < service config (config.json)
 
 deployment.yaml
 ```yaml
 apiVersion: extensions/v1beta1
 kind: Deployment
 metadata:
   name: myservice
 spec:
   template:
     spec:
       containers:
       - name: myService
         env:
           - name: DB_HOST
             value: kube 
 ```
 config.json
```json
{
  "services": {
    "myService": {
      "deploymentFile": "../..",
      "envs": {
        "DB_HOST": "service"
      }
    },
    "globalEnvs": {
        "DB_HOST": "global"
    }
  }
}
```
will result in:
docker-compose.yaml
```yaml
version: "2"
services:
    myService:
      environment:
        - DB_HOST=service
```

