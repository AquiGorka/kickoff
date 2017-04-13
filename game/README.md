# Game Simulation (Soccer-like game)

## Soccer rules

Out of bounds
Goal scored
Ball stuck


## Flow

execute
  wait for connections


## Interface

ping
  pong

setup (id, size, players, duration, start)
  return id

connect(id)
  return event emitter

asynch
  publish game state via event emitter

eventEmitter.onKick
  update game state
