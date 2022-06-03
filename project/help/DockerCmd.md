# Docker

- `docker login`
- `docker build -f filename.dockerfile -t ilingu/xxx-service:1.0.0 .` (be in the right dir)
- `docker push ilingu/xxx-service:1.0.0`

# Updating

- `Generate New Binary of the update`
- `docker build -f filename.dockerfile -t ilingu/xxx-service:1.0.1 .`
- `docker push ilingu/xxx-service:1.0.1`
- Change the `Swarm.yml` version of this service to the new one

# Swarm

- `docker swarm init`
- `docker swarm worker` --> Return cmd to exec to add a worker to the current swarm
- `docker swarm manager` --> Return cmd to exec to add a manager to the current swarm
- `docker stack deploy -c swarm.yml myapp` --> Deploy
- `docker service ls` --> See all the active services
- `docker service scale myapp_xxx-service=3` --> Change number of services
- `docker service scale myapp_xxx-service=0` --> Stop de service
- `docker service update --image ilingu/xxx-service:1.0.1 myapp_xxx-service` --> Upgrade to New Version or Downgrade
- `docker stack rm myapp` --> Remove app from swarm
- `docker swarm leave --force` --> ⚡ Remove/Delete swarm server entirely
