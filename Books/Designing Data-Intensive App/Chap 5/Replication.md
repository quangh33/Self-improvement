#### Background
- What: keeping copy of same data on multiple machines that are connected via 
a network
- Why:
    - Keep data near your users (reduce latency)
    - Fault tolerance
    - Increase read throughput
    
#### Leader-based replication
- Write to leader only
- Leader replicates data to followers 
    - sync replicate: leader sends the data and wait until followers confirmation
    before reporting success to user.
    - async replicate: leader sends the data without waiting for response from followers
    - in practice, it is impractical for all followers to be sync: any 1 follower
    outage would cause the whole system to stop => when you enable sync replication, 
    it means that `one` of the followers is sync.

- Read from leader or any of followers
   
##### Replication log implementation
###### Statement-based replication
- leader logs every write queries that it executes and sends that statements log 
to followers
- MySQL before 5.1, VoltDB
- Issues:
    - nondeterministic func: NOW(), RAND()
    - queries use autoincrementing column
    - queries have side effect (triggers, stored procedure)

###### WAL
- Write-ahead log    
- An append-only sequence of bytes containing all writes to the DB: describe
which bytes were changed in which disk blocks.
    - log-structured storage engine (Cassandra): this log is the main place for 
    storage
    - B-tree: every modification is first written to a WAL so index can be restored
    in case of crash.
- Downside:
    - closely couple to storage engine => can be a trouble if we change the storage format

###### Logical log replication
- A sequence of records describing writes to DB at row level:
    - insert row: log contains new value for all columns
    - delete row: primary key of deleted row
    - update row: primary key of updated row and new values of all columns
- Upside:
    - easier for external app
    - decoupled from storage engines internal => backward compatible    
        