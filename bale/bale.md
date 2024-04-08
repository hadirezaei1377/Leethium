Bale (messaging company) wants to limit its users in channels. After much research on the methods ahead, he decided to represent each user's access with a number. Each bit of this number represents one of the accesses. A value of 1 for each bit indicates access and zero indicates no access.

Ballet has considered two functions GetUserPermissions and SetUserPermissions to manage permissions in the main.go file:

The function GetUserPermissions accepts an input of type int8, each bit of that number represents each of the accesses, and as an output, it returns a structure called Permissions, which contains different accesses of type bool.
The SetUserPermissions function takes a Permissions structure as input and returns an int8 number.
Ballet has considered 6 accesses for this task, which are listed below in order from the least valuable bit to the most valuable bit:

canSeeMessages (least significant bit)
canDeleteMessages
canEditMessages
canKickMembers
canMakeMembersAdmin
canAddMembers

Now, he asked you to implement these two functions so that he can manage his channels better than before.