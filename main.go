package main

import (
	"fmt"
	"time"

	"github.com/Alfian57/belajar_golang_gorm/model"
	"github.com/Alfian57/belajar_golang_gorm/repository"
	"github.com/Alfian57/belajar_golang_gorm/util"
	"github.com/inancgumus/screen"
)

func main() {
	db := util.OpenDbConnection()
	userRepository := repository.NewUserRepository(db)

	for {
		var menu string

		screen.Clear()

		fmt.Println("======== User Management App  ========")
		fmt.Println("1. Show User")
		fmt.Println("2. Add User")
		fmt.Println("3. Delete User")
		fmt.Println("4. Exit")
		fmt.Print("Choose menu [1-4] : ")
		fmt.Scan(&menu)

		if menu == "1" {
			showUser(userRepository)
		} else if menu == "2" {
			addUser(userRepository)
		} else if menu == "3" {
			deleteUser(userRepository)
		} else if menu == "4" {
			break
		} else {
			screen.Clear()
			fmt.Println("==============")
			fmt.Println("Menu not valid")
			fmt.Println("==============")
			time.Sleep(3 * time.Second)
		}
	}
}

func showUser(userRepository repository.UserRepository) {
	var next string
	users := userRepository.FindAll()

	screen.Clear()
	printUser(users)

	fmt.Print("\nEnter to continue")
	fmt.Scanln(&next)
}

func addUser(userRepository repository.UserRepository) {
	var newName string

	screen.Clear()
	fmt.Print("Enter new user's name : ")
	fmt.Scan(&newName)

	userRepository.Create(newName)

	screen.Clear()
	fmt.Println("==========================")
	fmt.Println("Success to create new user")
	fmt.Println("==========================")
	time.Sleep(3 * time.Second)
}

func deleteUser(userRepository repository.UserRepository) {
	var userId int
	users := userRepository.FindAll()

	screen.Clear()
	printUser(users)
	fmt.Print("Enter user's id : ")
	fmt.Scan(&userId)

	userRepository.Delete(userId)

	screen.Clear()
	fmt.Println("======================")
	fmt.Println("Success to delete user")
	fmt.Println("======================")
	time.Sleep(3 * time.Second)
}

func printUser(users []model.User) {
	for _, user := range users {
		fmt.Printf("%d. %s\n", user.ID, user.Name)
	}
}
