package commands

import (
    "strings"
)

func parseCommand(rawcmd string)([]string, string){
    args := strings.Fields(rawcmd)
    cmd := args[0]
    args = append(args[:0], args[0+1:]...)

    return args, cmd
}
