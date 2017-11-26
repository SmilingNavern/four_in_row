#!/bin/bash


(
cat <<EOF
settings player_names player0,player1
settings your_bot player0
settings timebank 10000
settings time_per_move 500
settings your_botid 0
settings field_width 7
settings field_height 6
update game round 18
update game field .,0,1,0,0,.,.,1,0,1,1,1,.,.,1,1,0,0,1,.,.,1,0,1,1,1,0,.,0,1,1,0,0,1,0,0,0,1,0,0,1,0
action move 6874
quit
EOF
) | ./four_in_row/four_in_row
