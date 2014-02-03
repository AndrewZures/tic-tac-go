package ttt_test

import (
	. "github.com/andrewzures/tictactoe/ttt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Human Player", func() {
    var human *HumanPlayer
    var board []int

    BeforeEach(func(){
		    board = make([]int, 9, 9)
    })

    It("makes a move", func() {
        Expect(human.MakeMove(board)).To(Equal(1))
        Expect(human.MakeMove(board)).To(Equal(4))
    })

})
