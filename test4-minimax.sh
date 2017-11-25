#!/bin/bash


echo "SUPPOSED SOLUTION: 3"
(
cat <<EOF
settings player_names player0,player1
settings your_bot player1
settings timebank 10000
settings time_per_move 500
settings your_botid 1
settings field_width 7
settings field_height 6
update game round 9
update game field .,.,.,.,.,.,.,.,.,.,.,.,.,.,.,.,.,.,.,.,.,.,.,.,.,.,.,.,.,.,.,.,.,.,.,.,.,.,.,.,.,.
action move 10000
quit
EOF
) | ./four_in_row/four_in_row
