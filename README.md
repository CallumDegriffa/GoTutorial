This project has been set up to learn some of the core go language concepts
it will host multiple REST endpoints each performing a different function within the application.

Dependencies:

- java based microservice providing some data processing service
- cassandra nosql DB for data storage
- kafka for event driven processing
- docker to containerize the application
- docker-compose to spin up all the dependent services. 
- kubernetes to manage the deployment and scaling of the system. 
- react and node.js for an admin panel
- ELK stack for logging and monitoring

# TODO
- configuration file for injecting env vars
- DB integration with Cassandra
- dockerfile
- concurrent async operations