clean:
	sh scripts/clean.sh

snapshot:
	goreleaser release --snapshot --clean

run:
	make snapshot
	./dist/tlsman_darwin_arm64_v8.0/tlsman