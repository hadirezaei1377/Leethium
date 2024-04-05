The job of a person is that they give him some integers that are the amount of profit or loss of the company in consecutive days. (The unit of numbers is million tomans.) He must say how much the company's maximum profit was. For example, on the first day, they gave him these numbers:
1,2,−5,4,−3,2.
It is clear that the highest profit of the company was on the fourth day, which is equal to 4 million Tomans. Because the sum of the members of any other subarray of this given array has a value smaller than 4. Pay attention that if all the numbers were negative (loss), the amount of profit is equal to 0. Write a program by which Farzad can do his work without mental calculations.

Entrance
In the first line of input, the number of days for which profit and loss are to be taken, and then the array of profit and loss in these days is taken.

output
In your output, you should state the maximum profit. Pay attention to the input and output of the sample.

Example
Sample entry 1
12
7 -1 -2 1 5 -11 9 1 4 -1 3 -10

Sample output 1
16

Explanation of the output: the company's maximum profit is on days 7 to 11, when the sum of numbers 7 to 11 is equal to 16.

Sample input 2
5
-5 -2 -9 -1 -3

Sample output 2
0