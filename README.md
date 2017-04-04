# Kickoff

A multiplayer real-time turn-based soccer html5 game for virtual reality.

## TODO

* Analysis
** Setup flows for an automated game (this will show the basic micro-services and their interactions)
*** game flow
*** match flow
*** player flow (bot)
*** viewer flow

* Project setup
** Define milestones to reach the basic automated game

* Development
** Tests

* MVP: automated game

## Flow

- execute game
- execute match
- execute N players for match
- execute viewer

### Game

A micro-service for a game simulation.
This service will receive settings to execute the game.
This service publishes game state every __config.defined__ time intervals.
This service will receive game interactions and update the game state acorrdingly.

execute
  wait for connections
    this connections will get a socket connection to an instance?
    is this a daemon type of script? always running and defining rooms? to which clients connect?
    or is this an on demand runable loop?
    either way the interface to interact with this service should be the same 
end

init
  init simulation
    world
    stadium
    players
  init limits
    duration
    start datetime?
  
start match
  listens to player interactions and updates the simulation accordingly
end match


### Match

A microservice for players and viewers to connect to.
This service works as the gateway to connect the outside world and the game simulation.
This service encapsulates the auth logic, match settings for the game, player interactions with the game and lets viewers subscribe to watch the game.

start
  setup
    setup players
    setup field
    setup duration + start datetime
    setup conditions (weather?)
end

### Player

A microservice to automate playing the game.
This service will connect to the game and send interactions randomly when needed to.
This service will receive the game state everytime the match service publishes such state.

execute
start
  connect
    which match? (if there is only one, guess where you're gonna get redirected to)
  game loop
    player interactions
      if the player has to take some action, do so
  end game loop
end

### Viewer

A microservice to watch the game.
This service will receive the game state everytime the match service publishes such state and render a visualization of the game.

execute
  connect
  update client simulation
end


