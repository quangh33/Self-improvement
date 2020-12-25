##### Word count example

Input 1 -> Map: (a,1) (b,1)

Input 2 -> Map:       (b,1)

Input 3 -> Map: (a,1)      (c,1)

Reduce -> (a,2)

Reduce -> (b,2)

Reduce -> (c,1)

```
Map(k,v)
    # k is filename, v is file's content
    split v into words
    for word w
        emit(w,1)
```

```
Reduce(k,v)
    # k is word, v is list of 1
    emit(len(v))
```

