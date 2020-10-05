#### Client/Server

![](../images/80dfe841.png)

- Centralize control


#### Layered Pattern

![](../images/0030c851.png)

- structure area of concern
- layers can only communicate with peer above/ below

Cons:
- Deep call chains
- May harm performance

#### N-Tier Pattern

![](../images/4ee1c025.png)

- same as Layered layer except each layer is in a separated server
Pros:
- scale out
- high abstraction, isolation
Cons:
- network = SPoF
- network may be slow
- hard to debug
