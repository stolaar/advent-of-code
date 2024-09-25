package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/spf13/cobra"
	"github.com/stolaar/advent-of-code/utils"
)

var (
	rootCmd = &cobra.Command{
		Use:   "aoc",
		Args:  cobra.MinimumNArgs(1),
		Short: "Run advent of code problem",
	}
	generateCmd = &cobra.Command{
		Use:   "generate",
		Args:  cobra.MinimumNArgs(1),
		Short: "Generate advent of code starter code",
	}
)

func getArgs(args []string) (string, string) {
	day := args[0]
	yearint := time.Now().Year()
	year := strconv.Itoa(yearint)

	if len(args) > 1 {
		argYear, err := strconv.Atoi(args[1])
		if err != nil {
			log.Fatal("Invalid year flag")
		}

		if argYear < 2015 || argYear > yearint {
			panic(fmt.Sprintf("Year must be between 2015 and %d", yearint))
		}

		year = strconv.Itoa(argYear)
	}

	dayInt, err := strconv.Atoi(day)
	if err != nil {
		panic("Invalid day argument")
	}

	if dayInt < 1 || dayInt > 25 {
		panic("Day should be between 1 and 25")
	}

	return day, year
}

func init() {
	rootCmd.AddCommand(generateCmd)

	rootCmd.Run = func(cmd *cobra.Command, args []string) {
		day, year := getArgs(args)
		utils.Run(year, day)
	}

	generateCmd.Run = func(cmd *cobra.Command, args []string) {
		day, year := getArgs(args)
		utils.Generate(year, day)
	}
}

func main() {
	rootCmd.Execute()
}
