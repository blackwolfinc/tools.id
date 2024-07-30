package cli

import (
	"app/handler"
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func RunApp(db *sql.DB) {
	fmt.Println("Select report to generate:")
	fmt.Println("1. Total Game Sales Report")
	fmt.Println("2. Most Popular Game Report")
	fmt.Println("3. Total Revenue Per Game Report")
	fmt.Println("4. Player Count Per Game Report")
	fmt.Print("Enter the number of the report you want to generate: ")
	reader := bufio.NewReader(os.Stdin)
	choiceInp, _ := reader.ReadString('\n')
	choiceInp = strings.TrimSpace(choiceInp)
	choice, err := strconv.Atoi(choiceInp)
	if err != nil || choice < 1 || choice > 4 {
		fmt.Println("Invalid input. Please enter a number between 1 to 4")
		return
	}
	switch choice {
	case 1:
		handler.ShowTotalGameSalesReport(db)
	case 2:
		handler.ShowMostPopularGameReport(db)
	case 3:
		handler.ShowTotalRevenuePerGameReport(db)
	case 4:
		handler.ShowPlayerCountPerGameReport(db)
	}

}