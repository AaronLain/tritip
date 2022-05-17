# Tritip

## Another specific solution to a very specific problem

Tritip is used in conjuction with [Tenderloin](https://github.com/AaronLain/tenderloin) to update orders in ShipStation using their API. The data produced by Tenderloin is what Tritip uses to update orders. The goal is to eventually merge these two solutions, but for now we are living that fast/hard micro-microservice life. YOLO.

You will need a ShipStation account and an API key for this to work. 

To install make sure you have Golang installed then clone this repo. Then you will build `tritip.go`.
```
go build tritip.go
```

Similar to Tenderloin, you will now have a binary you can run using a csv file as an argument.
```
./tritip /path/to/your/file.csv
```

The results of the updates will appear in your terminal.
