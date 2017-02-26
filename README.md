# k8s-website-controller
A barely working example of a Kubernetes controller, which watches the Kubernetes API server for `website` objects and runs an Nginx webserver for each of them. 

NOTE: This is not the correct way to create a Kubernetes controller, as it simply performs a watch request in a loop. This will never work properly, because some watch events will be lost. 

