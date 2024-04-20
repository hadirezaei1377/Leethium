"Ali" has decided to travel from Shukaristan to Namakistan. He knows that in this vicinity there are n cities with numbers 1 to n (Shakrestan city number 1 and Namakestan city number n) and there are m two-way roads between these cities. We know that there is at most one road between any two cities and for each road we know how long it takes to travel. He does not want to pass through the same city twice on his journey.
Recently, the scientists of Shukrastan launched a strange communication system called "Run and Come" in some of these n cities. This system works in such a way that if a city has this system, it can go from one city to another city that has this system after t seconds. The strange part of this system is that this t number may be negative!
If he can make such a journey, what is the minimum time it takes him to travel from Shukaristan to Namakistan? (Strangely, this value can also be negative!)

Entrance
In the first line of the input, there are three numbers n, m and t, which respectively represent the number of cities, roads and the value of t. Shukrastan is city number 1 and Namakestan is city number n.
In the second line, there is a string with n letters, the i-th letter of which is equal to 1 if and only if there is a "run" communication system at the top of number i, and otherwise the i-th letter of the string is equal to 0.

Then, in the i-th line of each of the next m lines, three numbers u, v, and w
   It is said that it shows the two final cities and the time to travel the road between these two cities.
output
On the only output line, print a number that is equal to the minimum time it takes to travel from Shukaristan to Namakistan. If he cannot make such a trip, the word

Print "Impossible".

Example
Sample entry 1
3 2 3
101
1 2 2
2 3 2

Sample output 1
3

Sample input 2
4 4 -6
0110
1 2 2
2 3 2
3 4 2
1 4 1

Sample output 2
-2

Sample entry 3
4 2 10
1000
1 2 1
2 3 1

Sample output 3
Impossible