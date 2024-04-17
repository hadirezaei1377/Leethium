Codecap team has encountered a problem to prepare the keyboards for the Codecap live stage. These keyboards are of very poor quality and their keys break quickly. After a key is broken nothing happens when you press it.

In fact, each key on the keyboard has a specific life and the life of all the keys on a keyboard is the same. We call the number of acceptable times before a key on the keyboard breaks down, the "quality number". For example, a keyboard's "quality number" may be 2, in which case each key will work 2 times and then fail.

That is, in this keyboard for the following input text:

Welcome to CodeCup7

You will receive the following text:

Welcome to Cdup7

Here are some of the letters that were corrupted; For example, o crashed and stopped working.

Now, in this matter, we want you to specify the "quality number" of a keyboard, if the user wants to enter an initial text with it, what text will be entered instead.

These keyboards do not have Caps Lock at all.
These keyboards have only one shift.
The keyboard basically cannot hold a key and for example to type "aa" you have to press the a key twice.

There are two functions defined in the main.go file that you need to complete:

NewKeyboard function
This function is supposed to initialize and return an object of type Keyboard. The input of this method is the "quality number" of the keyboard. You can also design and use the internal structure of the Keyboard as you wish.


type Keyboard struct {
      // TODO
}

func NewKeyboard(dure int) *Keyboard {
return &Keyboard{
// TODO
}
}

Enter function
This function takes an input string and pretends that the user has typed that string. Return what is actually typed in the output.



func (keyboard *Keyboard) Enter(inp string) string {
// TODO
return ""
}


For simplicity, let's assume that the input inp consists only of the following characters:

English uppercase and lowercase letters
English numbers
normal distance
Special characters opposite: ?!'