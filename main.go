package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"time"
)

type Expense struct {
	Id          int    `json:"id"`
	Date        string `json:"date"`
	Description string `json:"description"`
	Amount      int    `json:"amount"`
}

func main() {
	expenses := InitializeExpenses()

	if len(os.Args) < 1 {
		fmt.Println("Please provide an operation: add, summmary, list, or delete.")
		return
	}

	operation := os.Args[1]
	switch operation {
	case "add":
		addCommand := flag.NewFlagSet("add", flag.ExitOnError)
		description := addCommand.String("description", "", "Description of the expense")
		amount := addCommand.Int("amount", 0, "Amount of the expense")

		addCommand.Parse(os.Args[2:])

		if *description == "" || *amount <= 0 {
			fmt.Println("Please provide valid description and amount for the expense.")
			return
		}

		AddExpense(expenses, *description, *amount)
	case "list":
		ListExpenses(expenses)
	case "summary":
		summaryCommand := flag.NewFlagSet("summary", flag.ExitOnError)
		month := summaryCommand.Int("month", 0, "Month for the summary (1-12)")

		summaryCommand.Parse(os.Args[2:])

		GetSummaryAmount(expenses, *month)
	case "delete":
		deleteCommand := flag.NewFlagSet("delete", flag.ExitOnError)
		id := deleteCommand.Int("id", 0, "ID of the expense to delete")
		deleteCommand.Parse(os.Args[2:])

		if *id <= 0 {
			fmt.Println("Please provide a valid ID for the expense to delete.")
			return
		}

		DeleteExpense(expenses, *id)
	default:
		fmt.Println("Invalid operation. Use add, summary, list, or delete.")
	}
}

func InitializeExpenses() []Expense {
	fileContent, err := os.ReadFile("expenses.json")
	if err != nil {
		fmt.Println("Error reading expenses file, creating the expenses.json file . . .")

		_, err := os.Create("expenses.json")
		if err != nil {
			fmt.Println("Error creating expenses file:", err.Error())
			return nil
		}

		fmt.Println("Success creating the expenses.json file")
		return nil
	}

	var expenses []Expense
	err = json.Unmarshal(fileContent, &expenses)
	if err != nil {
		fmt.Println("Error parsing expenses JSON:", err.Error())
		return nil
	}

	return expenses
}

func AddExpense(expenses []Expense, description string, amount int) {
	newExpense := Expense{
		Id:          len(expenses) + 1,
		Date:        time.Now().Format("2006-01-02"),
		Description: description,
		Amount:      amount,
	}

	expenses = append(expenses, newExpense)

	data, err := json.MarshalIndent(expenses, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling expenses to JSON:", err.Error())
		return
	}

	err = os.WriteFile("expenses.json", data, 0644)
	if err != nil {
		fmt.Println("Error writing expenses to file:", err.Error())
		return
	}

	fmt.Printf("Expense added: %s - $%d\n", newExpense.Description, newExpense.Amount)
}

func ListExpenses(expenses []Expense) {
	if len(expenses) == 0 {
		fmt.Println("No expenses found.")
		return
	}

	fmt.Printf("%-3s %-12s %-15s %-6s\n", "ID", "Date", "Description", "Amount")
	for _, expense := range expenses {
		fmt.Printf("%-3d %-12s %-15s $%d\n", expense.Id, expense.Date, expense.Description, expense.Amount)
	}
}

func GetSummaryAmount(expenses []Expense, month int) {
	months := []string{
		"",
		"January", "February", "March",
		"April", "May", "June",
		"July", "August", "September",
		"October", "November", "December",
	}

	if len(expenses) == 0 {
		fmt.Println("No expenses found.")
		return
	}

	totalExpenses := 0
	for _, expense := range expenses {
		extractedTime, _ := time.Parse("2006-01-02", expense.Date)
		extractedMonth := int(extractedTime.Month())
		if month != 0 && month == extractedMonth {
			totalExpenses += expense.Amount
		}

		if month == 0 {
			totalExpenses += expense.Amount
		}
	}

	if month != 0 {
		fmt.Printf("Total expenses for month %s: $%d\n", months[month], totalExpenses)
	} else {
		fmt.Printf("Total expenses: $%d\n", totalExpenses)
	}
}

func DeleteExpense(expenses []Expense, id int) {
	if len(expenses) == 0 {
		fmt.Println("No expenses found.")
		return
	}

	var newExpenses []Expense
	var isFound bool
	for _, expense := range expenses {
		if expense.Id == id {
			newExpenses = append(expenses[:expense.Id-1], expenses[expense.Id:]...)
			isFound = true
			break
		}
	}

	if !isFound {
		fmt.Printf("Expense with ID %d not found.\n", id)
		return
	}

	data, err := json.MarshalIndent(newExpenses, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling expenses to JSON:", err.Error())
		return
	}

	err = os.WriteFile("expenses.json", data, 0644)
	if err != nil {
		fmt.Println("Error writing expenses to file:", err.Error())
		return
	}

	fmt.Printf("Expense with ID %d deleted successfully.\n", id)
}
