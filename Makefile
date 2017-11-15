all:
	cd four_in_row && go build .

test:
	./test-minimax.sh
