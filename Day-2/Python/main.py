#Welcome Message
print("Welcome to the Tip Calculator.")
#Prompt the User for the Total Bill
bill=int(input("What was the total bill:$ "))
#Prompt the User for Percentage of Tip
percent=int(input("What Percentage tip would you like to give? 10,12 or 15? "))
#Prompt the User for Number of People
num_Person=int(input("How many people to spilt the bill? "))
#Calculate the Final Bill
final_Bill=round((bill/num_Person)*(1+(percent/100)),2)
#Print the Final Bill
print(final_Bill)