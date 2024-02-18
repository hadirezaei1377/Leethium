most important tests in GO:
unit test and integration test,

unit test related to test smallest unit in our program.
ingrate related to test after passing unit test,
for example we have a service that gives the number from users,
we check(test) whether it save correctly -----> unit test
we check(test) whether its already exists in our db -----> integration test

for example we have a project that has a service for giving name and phone number from user and check if it not exists in db sign user up.

at first we determine that the format of phone number must be correct.
in unit test file we write a test file for the function that implements this , we enter a correct format and a wrong format.

go test ./... run all tests
09138690958
0914
go test ./... -v

true
false

to determine which component is true and which is wrong we use assert package
assert.True.......09138690958
assert.False......0914

we use mock package and its ON method for unit test.
so we can write another test for when user exists in db ...

integration test:
in this case our test related to another sources out of our source and we cant using mocking.
for this we can use dockertest and its explanations.
it downloads all needed image.
