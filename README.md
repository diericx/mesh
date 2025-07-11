# Mesh

## High level story

### Controller

- keeps record of all nodes and networks
- acts as the authority
- acts as a STUN server
- has an open http and udp port

### Node

- doesn't need any open ports
- joins a network with a static ip via REST call to controller
- gets a list of nodes with access details from the controller

### Node 2

- http and udp ports are NOT open
- Join network: main
  - Does it conflict with any locally configured networks?
  - Are we authroized to join this new network?
  - What will my ip be?
    - xxx.xxx.xxx.xxx
  - default status: admin
    - Status change can be requested but authorization must go through an admin

## HTTP API

### Controller

#### `/api/v1/networks`

Lists all networks on this node

Returns:

```javascript
// list of available networks
{
  networks: [
    {
      name: string,
      subnetRange: string,
    },
  ];
}
```

#### `/api/v1/networks/{networkId}`

Lists all networks on this node

Returns:

```javascript
// list of available networks
{
  nodes: [
    name: string,
    ip: string, // the internal ip of this node
  ];
}
```

#### `/api/v1/networks/{networkId}/join`

Nodes call this endpoint to join a network and esablish themselves in the entire network

Input:

```javascript
{
  uuid: string;
  name: string;
}
```

Returns:

```
{
  // A list of the nodes in the network, and their connectable ips
  nodes: [
    name: string;
    ip: string; // the internal ip of this node
  ]
}
```

### Node

Nodes do not have any HTTP API at this time

## UDP API

### `/api/v1/nodes/{nodeId}/stun`

This is our basic version of a stun server. Simply returns the external ip and port of the node.
It then saves these values in the app to broadcast to other nodes

Returns:

```
{
  ip: number,
  port: number
}
```
