#!/bin/bash


(
cat <<EOF
settings player_names player0,player1
settings your_bot player1
settings timebank 10000
settings time_per_move 500
settings your_botid 1
settings field_width 7
settings field_height 6
update game round 5
update game field .,.,.,.,.,.,.,.,1,.,.,.,.,.,.,0,.,.,.,.,.,.,1,.,0,.,.,.,.,1,.,0,.,.,.,0,1,.,0,.,.,.
action move 10000
EOF
) | ./four_in_row/four_in_row
