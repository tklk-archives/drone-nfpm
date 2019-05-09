# drone-nfpm
nfpm (deb and rpm packager) plugin for drone

## Usage

Execute from the working directory:

```
docker run --rm \
  -e PLUGIN_CONFIG=<path_to_config> \
  -e PLUGIN_TARGET=<target> \
  -v $(pwd):$(pwd) \
  -w $(pwd) \
  techknowlogick/drone-nfpm
```
