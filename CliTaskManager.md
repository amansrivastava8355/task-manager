# Exercise details
In this exercise we are going to be building a CLI tool that can be used to manage your TODOs in the terminal. The basic usage of the tool is going to look roughly like this:

```
$ task
task is a CLI for managing your TODOs.

Usage:
  task [command]

Available Commands:
  add         Add a new task to your TODO list
  do          Mark a task on your TODO list as complete
  list        List all of your incomplete tasks

Use "task [command] --help" for more information about a command.

$ task add review talk proposal
Added "review talk proposal" to your task list.

$ task add clean dishes
Added "clean dishes" to your task list.

$ task list
You have the following tasks:
1. review talk proposal
2. some task description

$ task do 1
You have completed the "review talk proposal" task.

$ task list
You have the following tasks:
1. some task description 
```
Note: Lines prefixed with $ are lines where we type into the terminal, and other lines are output from our program.

 In the bonus section we will also discuss a few extra features we could add, but for now we will stick with the three show above:

`add` - adds a new task to our list
`list` - lists all of our incomplete tasks
`do` - marks a task as complete

You are welcome to tackle the problem however you see fit, but below is the order I would recommend to start.

# 1. Build the CLI shell

For instance, let's say we defined the task list command to run the following Go code:

fmt.Println("This is a fake \"list\" command")
Then when we used that command with our CLI we should see the following:

$ task list
This is a fake "list" command
After stubbing out all 3 commands, try to also look at how to parse arguments for the task do and task add commands.

# 2. Write the DB interactions
After stubbing out your CLI commands, try writing code that will read, add, and delete data in any DB of your choice.


# 3. Putting it all together
Finally, put the two pieces your wrote together so that when someone types task add some task it adds that task to the DB.


# Bonus
As a bonus exercise, I recommend working on the following two new commands:
```
$ task rm 1
You have deleted the "review talk proposal" task.

$ task completed
You have finished the following tasks today:
- wash the dishes
- clean the car
```
The rm command will delete a task instead of completing it.

The completed command will list out any tasks completed in the same day. You can define this however you want (last 12hrs, last 24hrs, or the same calendar date).