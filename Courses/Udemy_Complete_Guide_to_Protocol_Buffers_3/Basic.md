#### Why proto buffer?
- To share data across languages
- JSON's disadvantage
    - no scheme enforcing
    - big in size because of repeated key
    - no comment, docs
- Proto buffer's advantage
    - fully type
    - compressed automatically
    - schema
    - docs, comment
    - cross language    
    - 3-10x smaller, 20-100x faster than XML

#### Default values
- bool: false
- number: 0
- string: ""
- bytes: empty bytes
- enum: first value
- repeated: empty list

#### Enum
