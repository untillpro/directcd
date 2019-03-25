# pullcds

Continious delivery using pull requests to source-code repositories

# Repository

- `build.sh` must exist
- `version` file or `build.sh version`

# Configurations

## Direct Pull

In this scenario `cd4sources` pulls source-code by itself, builds and runs delivered application

Configuration
  - Repository url, including user name and password
  - Branch name
  - Pull interval
  - Command to build
  - Command to get version
  

## Delegated Pull

In this scenario one or few `cd4sources` are servers, they pull and build software, others are just clients

## Notification Hub

One `cd4sources` is configured as a target for `webhooks` from different code-hosting solution. It gets hooks and notifies cd4sources servers.


- `cd4sources` configuration
  - repository
  - branch
  - pull-interval
- `cd4sources` pulls sources, runs `build.sh`


## Configuring Server

- Working directory

## Quering Version

- Client sends a `watch` request for given version
- Request is blocked for a given timeout until new version is available
- If `watch` request returns new version ID client sends a `download` request

Watch Request
- `http://sourcecds/watch?<params>`
  - repo=https://github.com/mycompany/myrepo`
  - repo=https://github.com/mycompany/myrepo`


# Links

Hooks
- https://developer.github.com/v3/repos/hooks/#create-a-hook
- https://developer.github.com/v3/activity/events/types/#pushevent
- https://docs.gitea.io/en-us/webhooks

Stop process
- https://www.ctl.io/developers/blog/post/gracefully-stopping-docker-containers/
  - `docker stop ----time=30 foo`, SIGTERM
  - `docker kill ----signal=SIGQUIT nginx`, if you want to initiate a graceful shutdown of an nginx server, you should send a SIGQUIT
  - `docker kill ----signal=SIGWINCH apache`, Apache uses SIGWINCH to trigger a graceful shutdown
- https://husobee.github.io/golang/ecs/2016/05/19/ecs-graceful-go-shutdown.html