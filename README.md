## Running 
Build binary:
```
cd src
go build matrix.go
```

Running generated tests
```
python3 test_it.py
```
Running program itself
```
cd src/
# if not already run
go build matrix.go
./matrix
```

Command line arguments:
-s single thread
-m multithread