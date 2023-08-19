#Welcome the User
print("Welcome to Treasure Island.")
#Tell the user what his mission is
print("Your mission is to find the treasure.")
#Give the user his first choice, also lower case the input so that it can match with the if statement
choice1 = input('You\'re at a cross road. Where do you want to go? Type "left" or "right" \n').lower()
#Use "if" command for making a choice
if choice1 == "left":
  choice2 = input('You\'ve come to a lake. There is an island in the middle of the lake. Type "wait" to wait for a boat. Type "swim" to swim across. \n').lower()
  #First Nested Choice
  if choice2 == "wait":
    choice3 = input("You arrive at the island unharmed. There is a house with 3 doors. One red, one yellow and one blue. Which colour do you choose? \n").lower()
    #Now ask the user for the last nested three choices
    if choice3 == "red":
      print("It's a room full of fire. Game Over.")
    elif choice3 == "yellow":
      print("You found the treasure! You Win!")
    elif choice3 == "blue":
      print("You enter a room of beasts. Game Over.")
    else:
      print("You chose a door that doesn't exist. Game Over.")
  #This is the else choice for the First nested "if" statement
  else:
    print("You get attacked by an angry trout. Game Over.")
#This is the else choice for the Parent "if" statement
else:
  print("You fell into a hole. Game Over.")