package ttt_test

import (
	. "github.com/andrewzures/tictactoe/ttt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Console UI", func() {
    var console *ConsoleUI
    var board []int

    BeforeEach(func(){
      board = make([]int, 9, 9)
    })

    It("meets player interface requirements", func() {
        Expect(console.PrintBoard(board)).To(Equal(true))
    })

})

