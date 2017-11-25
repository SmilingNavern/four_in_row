package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "strconv"
)

type Settings struct {
    player_names [2]string
    your_bot string
    your_botid int
    opponent_botid int
    field_columns int
    field_rows int
}

type Game struct {
    round int
    field []int
    current_player_id int
    next_move int
}



var chosen_next_move int

func display_field(game *Game, settings *Settings) {
    for i := 0; i < len(game.field); i++ {
        if i % settings.field_columns == 0 {
            fmt.Printf("\n")
        }
        fmt.Printf("%3d", game.field[i])
    }

    fmt.Printf("\n")
}

func update_settings(set *Settings, s string) {
    fields := strings.Split(s, " ")
    if len(fields) <= 1 {
        return
    }

    switch fields[1] {
    case "your_bot":
        set.your_bot = fields[2]
    case "your_botid":
        set.your_botid, _ = strconv.Atoi(fields[2])
        if set.your_botid == 0 {
            set.opponent_botid = 1
        } else {
            set.opponent_botid = 0
        }
    case "field_width":
        set.field_columns, _ = strconv.Atoi(fields[2])
    case "field_height":
        set.field_rows, _ = strconv.Atoi(fields[2])
    case "player_names":
        players := strings.Split(fields[2], ",")
        set.player_names[0] = players[0]
        set.player_names[1] = players[1]
    }
}

func process_update(game *Game, s string) {
    fields := strings.Split(s, " ")
    switch fields[2] {
    case "round":
        game.round, _ = strconv.Atoi(fields[3])
    case "field":
        game_fields := strings.Split(fields[3], ",")

        for number,value := range game_fields {
            if value == "." {
                game.field[number] = -1
            } else {
                player, _ := strconv.Atoi(value)
                game.field[number] = player
            }
        }
        //fmt.Printf("%+v\n", game)
    }
}

func take_action(game *Game, settings *Settings) {
    game.current_player_id = settings.your_botid

    minimax(game, settings, settings.your_botid, 0)
    fmt.Fprintf(os.Stderr, "Actual decision: %d\n", chosen_next_move)
    fmt.Printf("place_disc %d\n", chosen_next_move)
    //display_field(game, settings)
}

func minimax(game *Game, settings *Settings, player_id int, deep int) int {
    deep += 1
    if deep > 6 {
        return 0
    }
    //fmt.Printf("YOUR BOT: %d\n", settings.your_botid)
    //fmt.Printf("OPPONENT BOT: %d\n", settings.opponent_botid)
    //fmt.Printf("CURRENT: %d\n", player_id)
    if win_game(game, settings, player_id) {
        if player_id == settings.your_botid {
            return 10
        } else if player_id == settings.opponent_botid {
            return -10
        }
    }

    var scores []int
    var moves []int


    possible_moves := possible_moves(game, settings)

    if len(possible_moves) == 0 {
        // game is over with draw
        return 0
    }

    for _, possible_move := range possible_moves {
        field_number := make_move(game, settings, player_id, possible_move)


	if player_id == settings.opponent_botid {
            scores = append(scores, minimax(game, settings, settings.your_botid, deep))
            moves = append(moves, possible_move)
        } else {
            scores = append(scores, minimax(game, settings, settings.opponent_botid, deep))
            moves = append(moves, possible_move)
        }

        game.field[field_number] = -1




        //fmt.Printf("Scores: %v\n", scores)
        //fmt.Printf("Moves: %v\n", moves)
    }


    if player_id == settings.your_botid {
        max := -100
        var max_index int
        for index, val := range scores {
            if val > max {
                max = val
                max_index = index
            }
        }

        chosen_next_move = moves[max_index]
        return scores[max_index]
    } else {
        min := 100
        var min_index int

        for index, val := range scores {
            if val < min {
                min = val
                min_index = index
            }
        }

        chosen_next_move = moves[min_index]
        return scores[min_index]
    }
}

func horizontal_check(game *Game, settings *Settings, player_id int) bool {
    var count int
    //horizontal check
    for i := 0; i < settings.field_rows; i++ {
        count = 0
        for j := 0; j < settings.field_columns; j++ {
            n := (i * settings.field_columns) + j

            if (game.field[n] == player_id) {
                count += 1
            }

            if count >= 4 {
                return true
            }

            if count > 0 && game.field[n] != player_id {
                count = 0
            }

        }

    }

    return false
}

func vertical_check(game *Game, settings *Settings, player_id int) bool {
    //fmt.Printf("CHECK FOR %d\n", player_id)

    //display_field(game, settings)
    var count int
    //vertical check
    for j := 0; j < settings.field_columns; j++ {
        count = 0
        for i := 0; i < settings.field_rows; i++ {
            n := (i * settings.field_columns) + j

            if (game.field[n] == player_id) {
                count += 1
            }

            if count >= 4 {
                //fmt.Printf("PLAYER WON: %d\n", player_id)
                return true
            }

            //fmt.Printf("n=%d, player_id_n=%d\n", n, game.field[n])
            //fmt.Printf("count=%d\n", count)
            if count > 0 && game.field[n] != player_id {
                count = 0
            }

        }


    }
    return false
}

func diagonal_check(game *Game, settings *Settings, player_id int) bool {
    for i := 0; i < settings.field_rows; i++ {
        for j := 0; j < settings.field_columns; j++ {
            n := (i * settings.field_columns) + j

            if game.field[n] != player_id {
                continue
            }

            if check_direction(game, settings, player_id, [...]int{i+1, i+2, i+3}, [...]int{j+1, j+2, j+3}) {
                return true
            }

            if check_direction(game, settings, player_id, [...]int{i-1, i-2, i-3}, [...]int{j-1, j-2, j-3}) {
                return true
            }

            if check_direction(game, settings, player_id, [...]int{i+1, i+2, i+3}, [...]int{j-1, j-2, j-3}) {
                return true
            }

            if check_direction(game, settings, player_id, [...]int{i-1, i-2, i-3}, [...]int{j+1, j+2, j+3}) {
                return true
            }

        }

    }

    return false

}

func check_direction(game *Game, settings *Settings, player_id int, i_pos [3]int, j_pos [3]int) bool {
    for k := 0; k < len(i_pos); k++ {
        n := (i_pos[k] * settings.field_columns) + j_pos[k]

        if n < 0 || n >= len(game.field) {
            return false
        }

        if game.field[n] != player_id {
            return false
        }
    }

    return true
}


func win_game(game *Game, settings *Settings, player_id int) bool {
    return (horizontal_check(game, settings, player_id) ||
               vertical_check(game, settings, player_id) ||
                  diagonal_check(game, settings, player_id))
}

func make_move(game *Game, settings *Settings, player_id int, column int) int {
    //fmt.Printf("before move\n")
    //fmt.Printf("player: %d", player_id)
    //display_field(game, settings)

    var n int
    for i := settings.field_rows - 1; i >=0; i-- {
        n = (i * settings.field_columns) + column
        if game.field[n] == -1 {
            //fmt.Printf("move: %d(%d:%d)\n", n, i, column)
            game.field[n] = player_id
            return n
        }
    }

    return -1
}

func possible_moves(game *Game, settings *Settings) []int {
    var p_moves []int

    if len(game.field) == 0 {
        return p_moves
    }
    //returns integer with column number for possible move
    for i := 0; i < settings.field_columns; i++ {
        if game.field[i] == -1 {
            p_moves = append(p_moves, i)
        }
    }

    //return -1 if no move
    return p_moves
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    var settings Settings
    var game Game

    F_CYCLE:
    for {
        switch scanner.Scan() {
        case strings.HasPrefix(scanner.Text(), "settings"):
            update_settings(&settings, scanner.Text())
            //fmt.Printf("%+v\n", settings)
        case strings.HasPrefix(scanner.Text(), "update"):
            //TODO: find out how to fix this more accurate
            if len(game.field) == 0 {
                field_len := settings.field_columns * settings.field_rows
                game.field = make([]int, field_len)
            }
            process_update(&game, scanner.Text())
        case strings.HasPrefix(scanner.Text(), "action"):
            take_action(&game, &settings)
        case strings.HasPrefix(scanner.Text(), "quit"):
            break F_CYCLE
        default:
            continue
        }
    }

}
