#### Factors to consider
- Schema, record size
- Number of clients
- Type of queries, access patterns
- Rates of read and write queries
- Expected changes in any of these variables

#### Questions to answer
- Does DB support the required queries?
- Can DB handle the amount of data?
- How many read and write operations can a single node handle?
- How many node should the system have?
- How do we expand the cluster given expected growth rate?
- What is maintenance process?