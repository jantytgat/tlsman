clean:
	sh scripts/clean.sh

snapshot:
	goreleaser release --snapshot --clean

run_daemon:
	make snapshot
	./dist/tlsman_darwin_arm64_v8.0/tlsman daemon --log-level trace