# Game Simulation (soccer-like game)

## Dev Environment

Run container
```sh
./scripts/docker.sh
```

This will run the golang image and mount game source code and lib dir into the container - the lib dir is gitignored as it is the place where all the dependencies will be stored locally - think of it as a node_modules dir; it doesn't get committed and you do not have to install all dependencies everytime you run the container.

When the container is running and the game app is running it exposes the http server to http://localhost:8877

Install dependencies (from within the container at game/)
```sh
go get -v ./...
```

Build (from within container)
```sh
go install game
```

Execute (from within container at game/)
```sh
./scripts/run.sh
```

Enter running container if needed
```sh
docker exec -it kickoff-game /bin/bash
```


## Soccer rules

- Out of bounds
- Goal scored
- Ball stuck


## Flow

- execute
  - wait for connections


## Interface

- ping
  - pong

- setup (id, size, players, duration, start)
  - return id

- connect(id)
  - return event emitter

- asynch
  - publish game state via event emitter

- kick
  - update game state
