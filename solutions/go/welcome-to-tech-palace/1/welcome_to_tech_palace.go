package techpalace

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
    return "Welcome to the Tech Palace, " + strings.ToUpper(customer)  
}



// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
    var bor string = strings.Repeat("*", numStarsPerLine) 
    return bor + "\n" + welcomeMsg + "\n" + bor
}




// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {

   clean := oldMsg

   clean = strings.Replace(clean, "*",  "", -1)
   clean = strings.Replace(clean, "\n", "", -1)
   clean = strings.Trim(clean, " ")

   return clean 
}





