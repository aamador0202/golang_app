
# Golang Http Server Application

This repository was created to display a golang assignment. Contents of the assignment can be found in Golang_Assignment.pdf.




## Content

 - main.go
 - Dockerfile
 - go-app-deployment.yaml
 - go-app-service.yaml
 


## Go App Local Deployment 

You can run this application locally having already installed golang on your local machine. Please reference https://go.dev/doc/install

You may have to initialize a go module before hand. Please reference https://go.dev/doc/install. In this instance a go.mod file is provided.

Once the golang installation is complete you can now run the main.go file provided.

```bash
  go run main.go
```
The http server should be running now on your local machine. To view the application running use your browser and type in the follow URL example.

```bash
  localhost:8080/?limit=10&sortKey=views
```
The application is set to listen on port 8080.

The `limit` value will display a certain amount of information from the queried URLs provided in the golang assignment file. In the example above the `limit` value is set to `10` so it will display 10 counts of information found in the URLs that are being queried.

The sortKey value has 2 valid values which are `views` and `relevanceScore`. Using any other value will display a message saying 

```bash
 Invalid sortKey value, please try using views or relevanceScore
```
Once either `views` or `relevanceScore` is set it will display data obtained from the URLs in accending order of views or relevancescore depending on which one was provided.

## Go App Minikube Deployment 

To run the go application on a Minikube cluster make sure you have installed it correctly and with a container runtime. In my case I used docker.
You can find more information on how to install minikube here: https://minikube.sigs.k8s.io/docs/start/

Once the minikube cluser is up and running you can use my image found in dockerhub 

```bash 
anthony0202/golang_http_server
```
If you would like to create your own image tag for your docker hub you can find the Dockerfile in this repo as well.

To create the image run:

```bash
docker build -t <your_repo>/<image-name> .
```
Make sure that your Dockerfile is located in your current directy in order for it to get built.

Once the image is ready you can now run: 
```bash 
kubectl create -f go-app-deployment.yaml
```
If you are using a custom image you will need to edit the go-app-deployment.yaml file.

This deployment has 3 replicas, if you want less or more pods you will need to edit the go-app-deployment.yaml as well.

Now that the deployment is up and running we need to expose the application, we do this by creating a NodePort service. The manifest for the service is also provided in this repo. 

To create the service run:

```bash
kubectl create -f go-app-service.yaml
```
This will connect port `8080` of the pod to port `8080` of the service and expose it to the node on port `300007`

Since the minikube cluster is a single node you can find the node ip by running:

```bash
minikube ip
```

Once you have the minikube ip you can access the application by running:

```bash
<minkube_ip>:30007/?limit=10&sortKey=views
```

## Disclaimer

The main.go code is missing some error handlings when an invalid limit value is provided, such a negative integer or a string that cannot be converted to integer.

