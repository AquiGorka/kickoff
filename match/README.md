# Match

## Flow

execute
  ping game
    acknowledge pong
  setup game
    players: 6
    size: 20m, 14m
    duration: 10min
    start: now + 1min
    acknowledge id
  connect game
  wait for connections


## Interface

ping
  pong

availableGames
  return list of game objects

connectPlayer (id)
  return event emitter

connectViewer (id)
  return event emitter

asynch
  relay game state via event emitters

eventEmitter.onKick
  relay data to game.kick 
