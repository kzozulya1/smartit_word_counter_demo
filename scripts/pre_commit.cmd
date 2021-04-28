cd ..

go clean -testcache ./...
go test ./...
pause

golangci-lint.exe  run
pause