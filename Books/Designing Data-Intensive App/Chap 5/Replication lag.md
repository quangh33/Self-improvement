##### Read your own Writes
Techniques:
- when reading something that the user have modified, read from Master, otherwise,
read it from Follower
- track the time of last update, for 1 min after last update, make all read from the
leaders.
=> this may not work if users access the app by multiple devices. This metadata need
to be centralized.

##### Monotonic Reads
- issue: moving backward in time when user makes several reads from diff replicas
- Monotonic reads is a guarantee that this kind of anomaly does not happen.
- Strong consistency > Monotonic reads > Eventually Consistency
- Monotonic reads only guarantee that if user makes several reads in sequence, they
will not see time go backward, i.e they will not read older data after 
having read newer data previously. `They still can see old data`
- How: make sure that each user always read from same replica. Example: replica
can be chosen based on the hash of user ID.