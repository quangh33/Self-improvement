#### Advanced types
- oneof
```protobuf
    message MyMes {
      int32 id = 1;
      oneof oneOfThose {
        string first_name = 2;
        string last_name = 3;
      }   
    }
```

- maps
```protobuf
    message MyMes {
      map<string, int32> results = 1;
    } 
```

map scalar (except float/double) to values of any type

- timestamp
- duration