all:
	cd four_in_row && go build .

tests: test test2 test3 test4

test:
	time -p ./test-minimax.sh

test2:
	time -p ./test2-minimax.sh

test3:
	time -p ./test3-minimax.sh

test4:
	time -p ./test4-minimax.sh

test5:
	time -p ./test5-minimax.sh

test6:
	time -p ./test6-minimax.sh



zip:
	zip -r my_bot.zip four_in_row/
