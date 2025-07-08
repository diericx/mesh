# Mesh

## Controller

- keeps record of all nodes and networks
- acts as the authority
- acts as a STUN server
- has an open http and udp port

## Node

- doesn't need any open ports
- joins a network with a static ip via REST call to controller
- gets a list of nodes with access details from the controller

## Node 2

- http and udp ports are NOT open
- Join network: main
  - Does it conflict with any locally configured networks?
  - Are we authroized to join this new network?
  - What will my ip be?
    - xxx.xxx.xxx.xxx
  - default status: admin
    - Status change can be requested but authorization must go through an admin
