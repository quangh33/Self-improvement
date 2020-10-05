#### Service-oriented pattern

![](../images/06472a1a.png)
- use network to communicate
- HTTP, XML, SOAP, Binary

#### Micro-service

![](../images/949ce510.png)

- services calling services
- can use fast private network
- deployed on multiple servers

#### Message-bus pattern

![](../images/caa3aad9.png)

- services connected to shared data bus
- user message for communication
- support discovery, failover

Cons
- bus == SPoF
- can be slow
- hard to test, debug
