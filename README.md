# Kickoff

A multiplayer real-time turn-based soccer html5 game for virtual reality.

## TODO

* Analysis
** Setup flows for an automated game (this will show the basic layers and their interactions)
*** game flow
*** viewer flow
*** player flow (bot)
* Project setup
** Define milestones to reach the basic automated game
* Development
** Tests
** MVP: automated game

## Flow

- execute game
- execute match
- execute N players for match
- execute client

### Game

start
  wait for connections
end

init
  init world
    world
    stadium
    players
  init match
    duration
    start datetime?
  start match
  end match
end

### Match

start
  setup
    setup players
    setup field
    setup duration + start datetime
    setup conditions (weather?)
end

### Client

play
start
  connect
     which match? (if there is only one, guess where you're gonna get redirected to)
  loop
end



