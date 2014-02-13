Tic Tac Toe in Go
=======

###To Run:

1.  Install [Go](http://golang.org/doc/install)

     a.  A Good tutorial on how to set up your environment can be found [here](http://skife.org/golang/2013/03/24/go_dev_env.html).

     b.  Also, you may try to use the `go_setup` script, although it may not work on all systems

2.  move to your golang GOPATH location
3.  run "go get github.com/AndrewZures/tictactoe"
4.  move down to src/github.com/AndrewZures/tictactoe.  You are now ready to run/test the code

#####options
        1. ./tictactoe
        2. go build; ./tictactoe

###To Test

before: run `.test_setup` which will download ginkgo and gomega test suite

#####options
        1. cd into ttt_test/, run 'go test'

###Game Rules:
    Simply choose board spot if using Human Player:
        1. 1-9 for 3x3 board
        2. 1-16 for 4x4 board
        3  1-25 for 5x5 board


