Use the go mod init command to create a new module called github.com/example/processor.

Create a package within the module called count. Create a file in this package named count.go. This file has a single function called FromReader that takes in an io.Reader and returns an int and an error. On success, this function should return a count of the bytes in the io.Reader and nil for the error. If an error (other than io.EOF) occurs while reading, a count of 0 and the error should be returned.

Then, create a package within the module called read. Create a file in this package named read.go. This file has two functions:

- The first is called FromWeb. It takes in a string and returns an io.ReadCloser and an error. The function should use the net.http package to perform a GET using the input parameter. On success, the function should return an io.ReadCloser representing the contents of the web page and nil for the error. If an error occurs, it should return nil for the io.ReadCloser and the error.

- The second is called FromFile. It takes in a string and returns an io.ReadCloser and an error. The function should use the os package to read a local file at the specified path. On success, the function should return an io.ReadCloser representing the contents of the web page and nil for the error. If an error occurs, it should return nil for the io.ReadCloser and the error.

Next, create a file called main.go in the main package (at the root of our module). This file should declare a main function that checks to see if os.Args has a length of exactly two. If it does not, it should print out the message "no source specified". If it does, it should pass the second value in os.Args to the process function. If process returns an error, print out the error and exit. Otherwise, print out the result from process.

The process function has a string input parameter and returns a string and an error. It checks to see if the parameter starts with http:// or https:// . If so, it gets an io.ReadCloser and an error from read.FromWeb. Otherwise, it gets an io.ReadCloser and an error from read.FromFile. If either function call returns a non-nil error, return back an empty string and the error. Otherwise, call count.FromReader, passing it the io.ReadCloser. If count.FromReader returns a non-nil error, return back an empty string and the error. Otherwise, convert the result to a string and return it and nil for the error.

Build your program with go build -o processor main.go.

Invoke this program four times with the parameters:

    http://example.com
    /tmp/file1.txt
    /tmp/file2.txt
    https://foo.test

Put the output for all four runs into a file called library1.txt.

(Hint: you can use >> to append to an existing file.)
