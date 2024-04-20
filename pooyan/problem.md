Pouyan is a chubby teenager who has decided to lose weight by working hard at a marble shop.
The shop where Pouyan works offers marbles in n different sizes and each marble can have feathers or not. So you can in this shop
Found 2n different types of marbles.

Poyan shop has regular customers who visit the shop every day and look at the marbles. Each of these customers likes a bunch of marbles and if none of his favorite marbles are found in the store, he leaves the store upset. Also, if each of these customers likes a feathered marble of one size, he will never like a feathered marble of another size. (That is, among the favorite marbles of each customer, at most one of them is feathered.)

Pouyan wants to choose only one type (with or without feathers) of each marble size and continue to produce only it, without disturbing any of these fixed customers. That is, there must be at least one favorite marble of each of these customers in the production list. Meanwhile, Pouyan wants to minimize the number of feathered marbles in the production list.

By inputting the customer's interest list, if Pouyan can continue to produce marbles under the stated conditions and without upsetting his customers, output a state of the production list where the number of feathered marbles inside them is minimal.

Entrance
In the first line of the input, there is the number n, which represents the number of different sizes for the marbles of the dynamic shop.

In the second line, the number m is entered, which represents the number of regular customers of the dynamic shop.

Then, in the i-th line of each of the next m lines, the list of types of marbles that the i-th fixed customer likes is described. At the beginning of the number line
k
  It is said that it is equal to the number of types of marbles that this customer likes. Then
k
The type of marbles are described as two sizes and feathers. The size is a natural number and the maximum is equal to n, and if feathered is equal to 1, it means that this customer likes a marble with feathers of this size, and if it is equal to 0, it means that he likes an unfilled marble of this size. (A customer may like both feathered and unfeathered marbles of the same size. This is stated in two different descriptions in this row. One for a feathered marble of this size and one for a featherless marble of this size.)
Sample entry 1
1
2
1 1 0
1 1 1

Sample output 1
IMPOSSIBLE

Sample input 2
5
3
1 1 1
2 1 0 2 0
150

Sample output 2
1 0 0 0 0