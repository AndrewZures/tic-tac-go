package ttt_test


import (
	. "github.com/andrewzures/tictactoe/ttt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Factory", func() {
    var factory *TTTFactory

    BeforeEach(func(){
      factory = new(TTTFactory)
    })

    Context("Player Generation", func() {

      It("provides of a list of available player types", func() {
        playerList := factory.PlayerTypes()
        descriptions := getPlayerDescriptions(playerList)

        Expect(descriptions).To(ContainElement("Human"))
        Expect(descriptions).To(ContainElement("Computer"))
      })

      It("each player type in player list is a unique type", func() {
        playerList := factory.PlayerTypes()
        descriptions := getPlayerDescriptions(playerList)
        Expect(allElementsUnique(descriptions)).To(Equal(true))
      })

      It("provides human player when given a human player type (Player)", func() {
        humanTemplate := new(HumanPlayer)
        humanTemplate.NewHumanPlayer("X", "Human")
        player := factory.Player(humanTemplate)
        comparison := player == humanTemplate

        Expect(comparison).To(Equal(false))
        Expect(player.Description()).To(Equal("Human"))
        Expect(player.Symbol()).To(Equal("X"))
      })

      It("provides computerplayer when given a computer player type (Player)", func() {
        computerTemplate := new(ComputerPlayer)
        computerTemplate.NewComputerPlayer("O", "Computer")
        player := factory.Player(computerTemplate)
        comparison := player == computerTemplate

        Expect(comparison).To(Equal(false))
        Expect(player.Description()).To(Equal("Computer"))
        Expect(player.Symbol()).To(Equal("O"))
      })

    })

    Context("Board Generation", func() {

      It("generates list of avaiable boards", func() {
        boardList := factory.BoardTypes()
        descriptions := getBoardDescriptions(boardList)
        Expect(allElementsUnique(descriptions)).To(Equal(true))
      })

      It("provides board when given a board template", func() {
        board3x3Template := new(BasicBoard)
        board3x3Template.NewBoard("X")

        newBoard := factory.Board(board3x3Template)
        Expect(newBoard.Description()).To(Equal("3x3 Board"))
      })

    })

  })

  func allElementsUnique(elementList []string) (bool) {

    for i := 0; i < len(elementList); i++ {
      for j := i+1; j < len(elementList); j++ {
        if elementList[j] == elementList[i] {
          return false
        }
      }
    }

    return true
  }

  func getPlayerDescriptions(playerList []Player) ([]string) {
    options := make([]string, len(playerList))

    for i := 0; i < len(options); i++ {
      options[i] = playerList[i].Description()
    }

    return options
  }

  func getBoardDescriptions(boardList []Board) ([]string) {
    options := make([]string, len(boardList))

    for i := 0; i < len(options); i++ {
      options[i] = boardList[i].Description()
    }

    return options
  }
