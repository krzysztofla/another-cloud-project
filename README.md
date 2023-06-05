# another-cloud-project
[![Continuous Integration](https://github.com/krzysztofla/another-cloud-project/actions/workflows/ci-wokflow.yml/badge.svg)](https://github.com/krzysztofla/another-cloud-project/actions/workflows/ci-wokflow.yml)


IMPORTANT:
After setting up the cluster it won't show you external ip if you using docker desktop. Docker desktop use bear metal so it means that it will have the same ip as your machine. 

Second option is tu use minikube. It will create virtual machine and you will be able to see external ip. Than you can use minikube tunnel to expose your service to the internet.

Anyway for the first option use:

```bash
kubectl describe service <service-name>
```
Than check node port and:
```bash
curl localhost:<node-port>/<endpoint>
```