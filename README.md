
# Golang Expenses Tracker CLI APP

This is a CLI application for expense tracker to manage your expenses. 

Project from: https://roadmap.sh/projects/expense-tracker



## Features

- Add expense
- Delete expense
- Summary expenses (by month)
- List expenses 


## Documentation

### Add Expense
```
expense-tracker add --description "Buy groceries" --amount 30
```

### Delete Expense
```
expense-tracker delete --id 2
```

### Summary Expenses
```
expense-tracker summary
expense-tracker summary
```

### Listing All Expenses
```
expense-tracker lsit
```


## Run Locally

Clone the project

```bash
https://github.com/mathiasyeremiaaryadi/project-expense-tracker-cli-app-golang.git
```

Go to the project directory

```bash
  cd my-project
```

Install dependencies

```bash
go build
```

Start the server

```bash
./expense-tracker [command]
```
