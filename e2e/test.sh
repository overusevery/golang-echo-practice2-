rm -rf covdatafiles
mkdir covdatafiles
GOCOVERDIR=covdatafiles HOST=localhost USER=postgres PASSWORD=postgres go run  -cover cmd/api/main.go > /dev/null &
PID=$(echo $!)
sleep 5
go test -v e2e/*/*_test.go
pkill -P ${PID}