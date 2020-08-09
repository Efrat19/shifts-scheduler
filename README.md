# shifts-scheduler

## What is it

This is a simple task manager to distribute information of any kind from kubernetes configmaps via built-in Slack support
you have a configmap where you schedule shifts, for example:
```yaml

```
than you register a slack app with slash-command & webhook permissions, and you point it 

## Installation

The easiset way to deploy teleskope is with helm:
```shell
git clone git@github.com:Efrat19/shifts-scheduler.git
cd chart
kubectl create ns shifts-scheduler
helm install --name shifts-scheduler -f example.yaml
```

## How 

## Built With

* [go](https://golang.org/) - Programing language
* [docker](https://www.docker.com/) - Containerized with docker
* [helm](https://www.helm.sh/) - Packaged with helm

## Author

* [Efrat Lavitan](https://github.com/efrat19) - *Initial work*
