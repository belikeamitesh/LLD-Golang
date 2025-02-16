# Initializing rabbitmq on local 
### Install Docker 
### Use command -  docker run -d --hostname rmq --name rabbit-server -p 8080:15672 -p 5672:5672 rabbitmq:3-management
### Access the rabbitmq server on http://localhost:8080

### Meaning of the docker command - 
docker run	- Runs a new container from an image.
-d -	Runs the container in detached mode (in the background).
--hostname rmq	- Sets the container's hostname to rmq (useful for internal networking).
--name rabbit-server	- Assigns a custom name (rabbit-server) to the container instead of a random one.
-p 8080:15672 - 	Maps port 15672 (RabbitMQ Management UI inside the container) to 8080 on the host machine.
-p 5672:5672 - 	Maps RabbitMQ's AMQP protocol port 5672 from the container to the same port on the host.
rabbitmq:3-management - 	Specifies the image name (rabbitmq) and its tag (3-management), which includes the RabbitMQ management plugin.
