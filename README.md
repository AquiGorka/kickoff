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

* MVP
** Automated game


## Flow

- execute game
- execute match
- execute N players for match
- execute viewer


### Game Simulation (Soccer-like game)

A micro-service for a soccer game-like simulation.
This service will wait for connections to setup and initiate an instance of a soccer-like game where a physics simulation representing a field (with defined size & boundaries defined by parameters sent to the service), a number of players (defined by parameters sent to the service) and a ball will take place.
This service will receive game interactions from the connected peer that will modify the game state (in a soccer-like fashion: giving the ball direction & force to move it).
This service will publlish to the connected peer the current state of the physics simulation in a defined interval (defined via parameters sent to the service).
The soccer-like game simulation will take responsibility for soccer rules (what to do when the ball goes out of the defined boundaries, how to proceed after a goal has been scored, duration of the match, etc).

- setup
- parameters
- interactions
- game rules
- match duration

execute
  wait for connections
    any client that connects will receive a pub/sub stream of events (in reality update events of current simulation state)
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

execute
  setup automated match
    setup players
    setup field
    setup duration + start datetime
    setup conditions (weather?)
  wait for connections
    player
    viewer
end

connection
  player
    Send to player game state and user interactions
    Receive from player game interactions
  viewer
    Send to viewer game state and user interactions


### Player

A microservice to automate playing the game.
This service will connect to the game and send interactions randomly when needed to.
This service will receive the game state everytime the match service publishes such state.

execute
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


