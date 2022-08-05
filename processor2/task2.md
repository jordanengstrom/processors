We will next add a second command to our application. In addition to being able to count the bytes in a file or HTTP resource, we also are going to add the ability to return what language it is written in.

In main.go, write a new function called detect. It will take in an io.Reader and return a string and an error. Using the third-party library github.com/abadojack/whatlanggo, detect will determine the language in the text stored in the io.Reader. Convert the io.Reader to a string and use the whatlanggo.DetectLang function to find the language, represented in an instance of type whatlanggo.Lang and return the value of the instance's String method and nil for the error. If anything generates an error, return back an empty string and the error.

Modify main to require two parameters. The first is either the word "count" or the word "lang." The second is the path or URL to the resource. If there aren't two parameters, it should print out the message "expected command and resource" and exit. If there are two parameters, pass them both to process and proceed as before.

Modify the function process to take two string parameters. The first is the command and the second is the resource. Process the resource as before. If the command is count, call count.FromReader. If the command is lang, call detect. Otherwise, return an error with the text "unknown command."

Build your program with go build -o processor main.go

Invoke this program four times with the parameters:

    count http://example.com
    lang /tmp/file3.txt
    lang /tmp/file2.txt
    read /tmp/file3.txt

Put the output for all four runs into a file called library2.txt.

(Hint: you can use >> to append to an existing file.)
