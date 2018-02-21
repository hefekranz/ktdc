# ktdc

Converter // Kubernetes -> Docker-Compose

So far only a few fields and file types implemented.
Maybe helpful for some Devs out there.


Usage:

### Single Deployment File
```
ktdc --kdeploymentFile=deployment.yaml
```
Will create a new docker-compose file with one service inside.

### Predefined Compose File
```
ktdc \
  --kdFile=deployment.yaml
  --dcFile=docker-compose-init.yaml
```
Will add the service defined in kdFile to a predefined docker-compose.yaml

### With Config
```
ktdc \
  --kdFile=deployment.yaml
  --dcFile=docker-compose-init.yaml
  --config=config.json
```
Will do all the above and apply config to it
(Overwrite Envs, mostly)

