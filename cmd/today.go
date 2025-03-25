/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"time"
	"github.com/spf13/cobra"
)

// todayCmd represents the today command
var todayCmd = &cobra.Command{
	Use:   "today",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: showToday,
}

func init() {
	rootCmd.AddCommand(todayCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// todayCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// todayCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type courses struct{
	name string
	startEnd string
}

func showToday(cmd *cobra.Command, args []string) {

    mon := []courses{
        {"KLR", "8:30 - 11:45"},
        {"Programmierung II", "15:30 - 20:00"},
    }
    tue := []courses{
        {"Systemanalyse", "8:30 - 12:30"},
        {"Programm. Konzepte", "13:45 - 17:00"},
    }
    wed := []courses{
        {"Mathematik", "9:30 - 13:00"},
        {"Recht", "15:00 - 18:15"},
    }
    thu := []courses{
        {"FiBu", "9:00 - 12:45"},
        {"GdIT II", "13:30 - 17:45"},
    }
    fri := []courses{
        {"Wiss. Arbeiten", "9:00 - 13:00"},
        {"Präsikompetenz", "14:00 - 17:30"},
    }

    switch currDate := time.Now().Weekday(); currDate {
        case time.Monday:
            printDay(mon)
        case time.Tuesday:
            printDay(tue)
        case time.Wednesday:
            printDay(wed)
        case time.Thursday:
            printDay(thu)
        case time.Friday:
            printDay(fri)
        default:
            fmt.Println("You are free to do whatevs today! :D")
    }
}

func printDay(course []courses){
    fmt.Println("Der Vorlesungsplan", time.Now().Format("2006-01-02"))
    fmt.Println("")

    for _, course := range course {
        fmt.Printf("%-25s%s\n", course.name, course.startEnd)
    }
}

