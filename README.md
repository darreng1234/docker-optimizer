# Container Optimizer
We employ the Container Optimizer in our workflow to generate streamlined and high-performance images using predefined templates. This tool assists developers by providing reports on Dockerfiles, analyzing image layers associated with the node's images, and producing optimized images according to customized parameters provided by the user. This paraphrased description can be added to your GitHub repository.

## Setting up 
### Prerequisites
- Setup [Golang](https://go.dev/dl/) The application was written on Golang 1.17.3, the [Installation](https://go.dev/doc/install) can be done after the correct version has been downloaded
- [Docker Engine](https://docs.docker.com/engine/install/) installation, this is required so that the docker engine runs when the application is being run.

### Building from source
Make sure to move to the root directory of the application then execute below commands.
-  `go get` - This command installs all necessary dependencies needed which are present in the `main.go` file
- `go build` - This will compile and build the binary file required to run the application

```bash
$ go get
$ go build
```

- After the build there should a non readable binary file called `docker-optizer` which can be used to run the application.

## Usage
### Default Command Palette 
```bash
./docker-optimizer
Choose an optimization option to continue with the process

Usage:
  docker-optimizer [command]

Available Commands:
  analyse-layer    Analyze your Docker Image layers with your existing layers by passing a docker image id
  completion       Generate the autocompletion script for the specified shell
  help             Help about any command
  manifest-checker Scan your Docker manfiests by passing a Dockerfile
  slim-image       Slim your Docker image by passing build configs

Flags:
  -h, --help   help for docker-optimizer

Use "docker-optimizer [command] --help" for more information about a command.
```

### Analyze Layer
```bash
./docker-optimizer analyse-layer
Error: required flag(s) "imageId" not set
Usage:
  docker-optimizer analyse-layer [flags]

Flags:
  -h, --help             help for analyse-layer
  -i, --imageId string   The Image Id
```

### Dockerfile Checker
```bash
./docker-optimizer manifest-checker
Error: required flag(s) "file", "package" not set
Usage:
  docker-optimizer manifest-checker [flags]

Flags:
  -f, --file string      Docker manifest file path
  -h, --help             help for manifest-checker
  -p, --package string   Package manager used (pip|apt)
```

### Image Slimmer
```bash
./docker-optimizer slim-image      
Error: required flag(s) "configDir" not set
Usage:
  docker-optimizer slim-image [flags]

Flags:
  -c, --configDir string   The directory of the build configs
  -h, --help               help for slim-image
```