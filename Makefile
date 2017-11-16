all:
	cd four_in_row && go build .

test:
	./test-minimax.sh

test2:
	./test2-minimax.sh


zip:
	zip -r my_bot.zip four_in_row/
