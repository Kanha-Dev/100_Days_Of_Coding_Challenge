#This imports the random library for the computer side of the game
import random
#These are Game_Images for the preview
rock = '''
    _______
---'   ____)
      (_____)
      (_____)
      (____)
---.__(___)
'''

paper = '''
    _______
---'   ____)____
          ______)
          _______)
         _______)
---.__________)
'''

scissors = '''
    _______
---'   ____)____
          ______)
       __________)
      (____)
---.__(___)
'''
#We will initialise a list for the images, Note that the index will start from 0
images = [rock, paper, scissors]
#Ask the user for a choice between Rock, Paper and Scissors
user_choice = int(input("What do you choose? Type 0 for Rock, 1 for Paper or 2 for Scissors.\n"))
#Print the choice by converting the string to integer and passing it through the list
print(images[user_choice])
#Generate a random choice between 0 and 2 for the computer side
computer_choice = random.randint(0, 2)
print("Computer chose:")
#Print the computer choice by passing the random generated number through the list
print(images[computer_choice])
#Now use if else ladder and print the Output
#Its a threeway deadlock so try to minimise the number of if statement by arranging it in a suitable trio
if user_choice >= 3 or user_choice < 0: 
  print("You typed an invalid number, you lose!") 
elif user_choice == 0 and computer_choice == 2:
  print("You win!")
elif computer_choice == 0 and user_choice == 2:
  print("You lose")
elif computer_choice > user_choice:
  print("You lose")
elif user_choice > computer_choice:
  print("You win!")
elif computer_choice == user_choice:
  print("It's a draw")