To run the docker file first we need to build the image, we can do it using 

docker build -t [image_name]:[tag] [context_path]

```
docker build -t myimage:1.0 .
```

We use dot as the context path, because our docker file is in the current directory, means the folder path of Docker file. If your docker file has another name then

docker build -t [image_name]:[tag] -f [dockerfile] [context_path]

```
docker build -t myimage -f sampledockerfile .
```

List the images by executing the following command

```
docker images
```

To delete an image

docker rmi [image_name]
```
docker rmi myimage:latest
```

To run the image

docker run --name [container_name] [image_name]

```
docker run --name mycontainer myimage:1.0
```

To list the containers

```
docker ps -a
```

To delete the conatiner

docker rm [container_name]

```
docker rm mycontainer
```

**Docker Compose**


To run the container use the following command 

This works if your are in the same folder and the name of the file docker-compose.yaml
```
docker compose up -d
```
-d stands for detached mode

If you are having a different name use the following command
```
docker compose -f [path/to/compose/file] up -d
```
This command is used to validate the docker file
```
docker compose config
```

To view the logs
```
docker logs container_name
```

To view the couchdb container go to this url http://localhost:3000/_utils

**Volume Mapping**

To access the sample folder execute this command
```
sudo chmod -R 777 samplefolder
```

Enter this command to list out all the containers
```
docker ps -a

```
-a for showing terminated containers also.

**Note:** Copy the container ID and add it to the next command

Execute this command to open a bash shell inside the container
```
docker exec -it <container ID> bash
```

TO check the files inside the container execute this command
```
ls /hyperledger/fabric/newfolder
```

Now, create a file `sample.txt` in `samplefolder`

To view the file using the cat command
```
cat /hyperledger/fabric/newfolder/sample.txt
```

To exit from container 
```
exit
```

**Cleanup**

To bring down the containers which were up using docker-compose files use the following command
```
docker compose down
```

To remove the containers use the following commands
```
docker container prune

docker system prune

docker volume prune

docker network prune

```

To remove a particular image use this command
```
docker rmi <image_id_or_name:tag>

```



