This project is to showcase how to deploy multiple grpc service with kubernetes and rely on kube's internal dns to make them visible to each other

# To run

1. make sure you have a running kubernetes cluster (minikube will work, see [here](https://kubernetes.io/docs/tutorials/hello-minikube/) for details on how to run minikube).

2. run server as pod/service in kube cluster

   ```
   $ cd server && kubectl apply -f server.yaml
   ```

3. run client as pod/service in kube cluster

   ```
   $ cd client && kubectl apply -f client.yaml
   ```

4. verify the client & server are working by

   ```
   $ kubectl get pods
   $ kubectl logs -f <client-pod-name-from-output-above>
   $ (in another terminal) kubectl logs -f <server-pod-name-from-output-above>
   ```

   you should see client prints out response from server every few seconds
   you should also observe server prints received message every few seconds

5 (TODO) test load balancing with grpc
