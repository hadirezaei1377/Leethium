Majid, who has just learned to program, pseudo code for ascending sorting of an array by name
a to length
n has written which is clearly wrong. This code snippet is as follows:

for i = 1 to n - 2
  for j = 1 to n - 1
  if a[j] > a[j + 1]
  swap(a[j], a[j + 1])
c
Milad is a permutation of the numbers 1 to

n, some of its numbers are not known and zero has been placed instead.

Now we want to know if we can put numbers instead of zeros so that at the end:

Majid's code cannot sort the permutation correctly in ascending order.
Each of the numbers is 1
𝑛
n appear exactly once in the array.
(For better understanding, read the introduction and explanation of examples.)

Entrance
In the first line of the number
You are given n which is the number of permuted numbers. In the next line

n, separated by spaces, gives you the permutation numbers of AD.

output
If we could change the zeros in a way that met the conditions of the question, in the first line of the output, print Yes and in the next line, print the complete permutation along with the numbers that have been replaced instead of zeros. (As there were several permutations, you can print one of them as you wish.)

Otherwise, print No in one line.

Example
Sample entry 1
4
1 3 0 2
Sample output 1
No

Explanation: The only number that can replace 0 is 4, and Majid's program can arrange this permutation.

Sample entry 2
4
0 0 0 0

Sample output 2
Yes
4 3 2 1

Explanation: If we give the above permutation to the Majid program, it will not sort it correctly. Note that the above permutation is one of the answers to the problem; There are other permutations on which Majid's program does not work properly.