<p align="center">
  <img src="./docs/logo.svg" height="200">
  <h1 align="center">Baelfire</h1>
  <p align="center">Baelfire helps keep track of the version of different 3rd party softwares deploy on your stack (DBs, monitoring tools, ...)<p>
  <p align="center">
    <a href="https://github.com/Impsy/baelfire/blob/master/LICENSE">
      <img src="https://img.shields.io/badge/License-Apache%202.0-g.svg" />
    </a>
  </p>
</p>

## Install

Currently Baelfire is only available on [dockerhub](https://hub.docker.com/r/impsy/baelfire)

You can launch a Baelfire container for trying it out with
```
$ docker run --name baelfire -d -p 127.0.0.1:1323:1323 impsy/baelfire:v0.1.0
```
Baelfire will now be reachable at http://localhost:1323/.


## Projects supported (aka targets)
- Alertmanager
- Baelfire (self)
- Grafana
- Metabase
- Prometheus


## API
| URI                           | method  | Description |
|:----                          |:-------:|:------------|
| /api/v1/version               | GET     | Get project version |
| /api/v1/targets               | GET     | List all targets name |
| /api/v1/targets               | POST    | Create new target |
| /api/v1/targets/:name         | GET     | Get target detail |
| /api/v1/targets/:name         | DELETE  | Delete target |
| /api/v1/targets/:name/version | GET     | Get target version |

### Target Parameters
```json
{
  "name": "the name of the target",
  "project": "the project of the target (in lowercase)",
  "host": "the scheme, ip/dns and port of the project"
}
```