all:
	cd four_in_row && go build .

tests: test test2 test3 test4

test:
	./test-minimax.sh

test2:
	./test2-minimax.sh

test3:
	./test3-minimax.sh

test4:
	./test4-minimax.sh

zip:
	zip -r my_bot.zip four_in_row/
