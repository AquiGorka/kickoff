#### Notes

* Whitelist for players.
** Only selected players can join.
** If one disconnects then he will be allowed to come back when reconnected.

## File structure
/

## Microservices

* Management
* Game
* Match
* Auth
* Gateway
* Client
** Player
** Viewer

## Ideas

* Go as server?
** What about p2p streaming?
* Use cannon from Go?
* Bot for dev-flow steps
** Tells you where to continue and how (might even fill out required info)
** Ask for status, milestones, etc <- maybe integrate with git?
*** https://hubot.github.com/

### Management

teams
  add
  edit
match
  play?
    setup?
    join?


# Game
any client that connects will receive a pub/sub stream of events (in reality update events of current simulation state)
    this connections will get a socket connection to an instance?
    is this a daemon type of script? always running and defining rooms? to which clients connect?
    or is this an on demand runable loop?
    either way the interface to interact with this service should be the same 
