# Game Simulation (soccer-like game)

## Dev Environment

Run container
```sh
docker run -it --rm --name kickoff-game --volume "$PWD/src:/go/src/game" --entrypoint /bin/bash golang
```

Build (from within container)
```sh
go install game
```

Execute (from within container)
```sh
game
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

- eventEmitter.onKick
  - update game state
