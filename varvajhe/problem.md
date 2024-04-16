T is a crossword if T can be reached by permuting the letters of the string. For example, aba is a wildcard for the string aab, but not a wildcard for the string aaa.

S is a string that contains English lowercase letters and some characters. Also, the P string is a string that only contains English lowercase letters. A substring of S is called a good substring if it can be obtained by inserting arbitrary letters instead of P.

Entrance
In this question, you are first given the string S and then p.
output
The output consists of only one line in which an integer equal to the number of good substrings of string S is printed.

Example
Sample entry 1
bb??x???
aab

Sample output 1
2

In this example, two substrings b?? and ??? can be paraphrased by replacing their question marks and becoming baa and aab.

Sample input 2
ab? c
acb

Sample output 2
2

In this example, two substrings ab? and b?c by becoming abc and bac can be transformed into the string acb.