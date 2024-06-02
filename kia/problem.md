"Kia" has fallen in love with "Lily" and wants to buy a bouquet for "Lily"! For this purpose, he decided to pick a bunch of flowers from the garden next to the street, which has n flowers in a row. A bouquet consists of a subset of garden flowers. He knows that "Lily" likes a bouquet when it meets the following conditions.

The height of the flowers he picks should be equal.
The height of all the flowers between the selected flowers must be greater than or equal to the height of the selected flowers.
He must give all the flowers he picks to Lily. (That is, he does not have the right not to use a flower that has been arranged.)

Now Kia wants to win the most number of goals for "Lily" so that it has high conditions. "Kia" has fallen in love and cannot solve this problem. Tell him how many goals he can score for Lily.

Entrance
In the first line of input, an integer n is given, and in the next line n is a natural number, which is the number
The i-th is equal to the height of the i-th flower.


output
Print on one output line the maximum number of goals that Kia can score for Lily.

Example
Sample entry 1
5
8 9 8 3 8

Sample output 1
2

Explanation of sample 1
Kia can score the first and third goals. He cannot score the first, third and last goal; Because there is a flower with a height of 3 between the last and the third flower and it is shorter than them.

Sample entry 2
2
8 2

Sample output 2
1