Write a program that takes an integer n from the user and then takes n strings from the input. The output of the program will be the largest string like s, each string having s or its inverse as a substring. If there is no shared substring, print nothing.

Definition of substring: "consecutive" characters from the string whose beginning and end can be anywhere in the string and the order of the characters is exactly the same as the order in the main string. For example, the substrings of ABC are: A, B, C, AB, BC, ABC

Entrance
In a line you will be given n numbers then n strings.
Also, the length of the strings is a maximum of 20.

output
The substring printed on the output must be of the same form as in the first string, eg in the example it should be printed on the output CDEF, not FEDC.

Example:
Sample input
3
ABCDEF
FEDCAB
GHCDEFJK

Sample output
CDEF