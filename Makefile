.PHONY: test, new

test:
	go test ./...

new:
	cp -r DayTemplate Day$(NUMBER)
	sed -i '' 's/DayTemplate/Day$(NUMBER)/g' Day$(NUMBER)/*.go