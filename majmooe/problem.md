Normal sets, which we call one-layer sets in this question, are sets whose members are only numbers. In this question, we are dealing with multi-layer sets whose members can be another set in addition to the number, which may also contain another set in their heart.

In other words, a multi-layer set is a set whose members can be numbers or another multi-layer set, and a single-layer set is a set whose members are only numbers and has no set members.

(Every one-layer set is also a multi-layer set, but every multi-layer set is not a single-layer set.)

To sum a multi-layer set, for each multi-layer set of its members, we put the sum of that multi-layer set and add these numbers with other numbers of the set members.

A multilayer set is given as input. We want to get the sum of the members of the set and of course the sum of all the members of its nested sets. To sum a set, we act in such a way that if all the members of that set were numbers, we print the sum of those numbers. Otherwise, first do this for all the sets in it (in the order of their placement from left to right). and when we get and print the sum of all the sets inside it, we add them together as well as the other numbers of the set members. For each collection seen.

You can assume that we don't have the null set and that the numbers are all non-negative.

Entrance
You are given a multi-layered set in one line. The length of the input string is less than 100.

output
Ideally, print the problem (product of sums) on separate lines.

Example
Sample entry 1
{1, 2, {3, {4, 5, {6}}, 7}, 8}

Sample output 1
6
15
25
36

Sample entry 2
{{12, 23, {4, 0, {1}, {1}}}, 0, {1}}
Plain text
Sample output 2
1
1
6
41
1
42

Sample entry 3
{1, {2, {{6}}}, {{{7}}}}

Sample output 3
6
6
8
7
7
7
16