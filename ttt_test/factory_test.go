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
