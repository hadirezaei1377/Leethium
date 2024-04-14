In this issue, Salib wants to define a game to escape from his eternal loneliness and have some fun. This time he wants to write an interactive program to find the number that the other code selected :). The game is that Salib wants to implement the GuessMyNumber function in the main.go file. The signature of this function is as follows:

func GuessMyNumber(game Game) string {

This function receives an object of type Game. The purpose of this function is to find the number selected by the game and stored inside the game. But it is not possible to find directly and we must play! In fact, we only have access to a function that we pass our guesses to, and it checks our guesses and guides us to get closer to the original answer.

Finally, the GuessMyNumber function, after making various guesses and finally finding the answer, returns its answer with a message in the form of a string:

Your Number was <NUMBER>

For example, if the game number is equal to 269, this function should return the following expression after finding out the number.

Your number was 269

Game details
In the utils.go file, the Game interface is mentioned, which we can also see here:

package main

type Game interface {
CheckNumber(n int) string
}


The CheckNumber function checks the number it takes as a guess and returns 3 different states:

If the input number of the function, i.e. n, is less than the desired number of the game, it will return the expression My Number is Greater.
If the input number of the function, i.e. n, is greater than the desired number of the game, it will return the statement My Number is Lower.
If the input number of the function or n is exactly equal to the desired number of the game, it will return the expression CORRECT.

Since the implementation of the CheckNumber function is expensive for Salib, Salib has defined the problem in 4 parts:

Score issue
By asking a maximum of 360 questions, the number will reach 25
By asking a maximum of 180 questions, the number will reach 25
By asking a maximum of 60 questions, the number will reach 25
By asking a maximum of 10 questions, the number will reach 25

  The structure of the project files is as follows:


├── main.go
├── utils.go


The utils.go file contains the Game interface. file

The number range you have to guess is between 1 and 360
