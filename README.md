# k8s-website-controller
A barely working example of a Kubernetes controller, which watches the Kubernetes API server for `website` objects and runs an Nginx webserver for each of them. 

NOTE: This is not the correct way to create a Kubernetes controller, as it simply performs a watch request in a loop. This will never work properly, because some watch events will be lost. 

## How to build the application?

***Requirements:*** Docker already installed and started on your environment.

For linux OS with Go already installed, you can build the application with the `make` command. You can also build the application and create the docker image with `luksa/website-controller` as name and `latest` as tag directly using the `make image` command. Finally, you can build the application, create the docker image and push it directly in a docker registry with the command `make push`.


For all other OS and Linux without Go installed, you can build the application and create the docker image with `luksa/website-controller` as name and `latest` as tag, using the command `docker build . -f Dockerfile.Build -t luksa/website-controller`. In this case, the application is built while creating the docker image.
