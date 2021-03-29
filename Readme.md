# Tasky
## _Make tasks, do stuff_

[![Build Status](https://github.com/hunter32292/tasky/actions/workflows/build.yaml/badge.svg)](https://github.com/hunter32292/tasky/actions/workflows/build.yaml)

Simple Command Line example of a task organization tool

- Create a task
- See them in a list
- ✨Magic ✨

## Usage

```sh
$ tasky help

	Tasky is a simple and extensible task manager for daily use in organizing your day to day tasks.
	Example usage:
	tasky create
		title : Pick up groceries
		content : # Grocery List
		1. Eggs
		1. Fruit cups
		1. Milk
		due date: 5d
	Task created:
	Pick Up Groceries
	
	Grocery List
	- Eggs
	- Fruit cups
	- Milk

	Due - Tue, 09 Mar 2021 15:56:39 -0800
```
Create, List, and Delete tasks
```sh
$ tasky create
Title :Make a small thing on the weekend 
Content :Probably something mildy useless
Due Date :2h

$ tasky list
0)	 Title:Make a small thing on the weekend 	Due Date:Sun, 28 Mar 2021 22:36:51 -0700

$ tasky clear
Please provide a Task ID to remove:
0

$
```

## Development

Want to contribute? Great!

Make a change and submit a pull request!