package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Welcome to Roulette!\nEnter the total amount you have?")
	var bank_roll float64
	fmt.Scanf("%v", &bank_roll)

	fmt.Print("Look at the our menu and choose your bets:\n 1. Bet on red or balck \n 2. Bet on high ot low (1-18 or 19-36) \n 3. Bet on a specific number\nChoose:")
	var bet_choice int
	fmt.Scanf("%d", &bet_choice)

	switch bet_choice {
	case 1:
		fmt.Print("Enter the betting amount: ")
		var bet_amt float64
		fmt.Scanf("%v", &bet_amt)

		if bet_amt > bank_roll {
			fmt.Println("Your balance is too low to play the game")
			return
		}
		bank_roll -= bet_amt
		fmt.Println("Choose Red or Black: ")
		var option string
		fmt.Scanf("%s", &option)
		spin := rand.Intn(36) + 1
		if spin%2 == 0 {
			spin_res := "Red"
			if spin_res == option {

				bank_roll += bet_amt
				fmt.Println("You won the bet! Your total balance is ", bank_roll)
			} else {
				fmt.Println("You Lost. Better luck next time. The color was ", spin_res)
			}
		} else {
			spin_res := "Black"
			if spin_res == option {
				bank_roll += bet_amt
				fmt.Println("You won the bet! Your total balance is ", bank_roll)

			} else {
				fmt.Println("You Lost. Better luck next time. The color was ", spin_res)
			}
		}
	case 2:
		fmt.Print("Enter the betting amount: ")
		var bet_amt float64
		fmt.Scanf("%v", &bet_amt)

		if bet_amt > bank_roll {
			fmt.Println("Your balance is too low to play the game")
			return
		}
		bank_roll -= bet_amt
		fmt.Println("Choose High or Low: ")
		var option string
		fmt.Scanf("%s", &option)
		spin := rand.Intn(36) + 1
		fmt.Println(spin)
		if spin >= 1 && spin <= 18 {
			spin_res := "Low"
			if spin_res == option {
				bank_roll += bet_amt
				fmt.Println("You won the bet! Your total balance is ", bank_roll)
			} else {
				fmt.Println("You Lost. Better luck next time. The number was ", spin_res)
			}
		} else if spin >= 19 && spin <= 36 {
			spin_res := "High"
			if spin_res == option {
				bank_roll += bet_amt
				fmt.Println("You won the bet! Your total balance is ", bank_roll)
			} else {
				fmt.Println("You Lost. Better luck next time. The number was ", spin_res)
			}
		}
	case 3:
		fmt.Print("Enter the betting amount: ")
		var bet_amt float64
		fmt.Scanf("%v", &bet_amt)

		if bet_amt > bank_roll {
			fmt.Println("Your balance is too low to play the game")
			return
		}
		bank_roll -= bet_amt
		fmt.Println("Enter the number you like to place the bet on: ")
		var option int
		fmt.Scanf("%d", &option)
		spin := rand.Intn(38)

		if spin == option {
			bank_roll = 35 * bet_amt
			fmt.Println("You won the bet! Your total balance is ", bank_roll)
		} else if spin == 37 && option == 00 {
			bank_roll = 35 * bet_amt
			fmt.Println("You won the bet! Your total balance is ", bank_roll)
		} else {
			fmt.Println("You Lost. Better luck next time. The number was ", spin)
		}
	default:
		fmt.Println("Invalid choice, Please try again")

	}
}
