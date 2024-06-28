go build -tags=master -o master -gcflags "all=-N -l" main.go
./master master --path=./config/profile-gc.json --node=gc-master &
go build -tags=center -o center -gcflags "all=-N -l" main.go
./center center --path=./config/profile-gc.json --node=gc-center &
go build -tags=gate -o gate -gcflags "all=-N -l" main.go
./gate gate --path=./config/profile-gc.json --node=gc-gate-1 &
go build -tags=game -o game -gcflags "all=-N -l" main.go
./game game --path=./config/profile-gc.json --node=10001 &
go build -tags=leaf -o leaf -gcflags "all=-N -l" main.go
./leaf leaf --path=./config/profile-gc.json --node=20001 &
go build -tags=web -o web -gcflags "all=-N -l" main.go
./web web --path=./config/profile-gc.json --node=gc-web-1