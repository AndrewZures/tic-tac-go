package ttt_test

import (
  . "github.com/andrewzures/tictactoe/ttt"
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
)

var _ = Describe("Console UI", func() {
  var messages Messages

  BeforeEach(func(){
    messages = new(ConsoleMessages)
    messages.BuildMessages()
  })

    It("returns player symbol if available", func() {
      Expect(messages.WinnerSymbol("x")).To(Equal("X"))
      Expect(messages.WinnerSymbol("o")).To(Equal("O"))
      Expect(messages.WinnerSymbol("z")).To(Equal("z"))
    })

    It("has intro message", func() {
      expectedIntroText := "Welcome to Tic Tac Go!"
      Expect(messages.IntroMessage()).To(Equal(expectedIntroText))
    })

    It("has exit message", func() {
      expectedExitText := "So Long!"
      Expect(messages.ExitMessage()).To(Equal(expectedExitText))
    })

    It("has choose move response", func() {
      expectedMovePromptText := "%v, Choose a Move!\n"
      Expect(messages.ChooseMovePrompt()).To(Equal(expectedMovePromptText))
    })

    It("has game ends in a tie response", func() {
      expectedTieText := "The Game Has Ended In A Tie\n"
      Expect(messages.GameTieResponse()).To(Equal(expectedTieText))
    })

    It("has game winner response", func() {
      expectedWinnerText := "The Winner is %v\n"
      Expect(messages.GameWinnerResponse()).To(Equal(expectedWinnerText))
    })

    It("has new game prompt", func() {
      expectedNewGameText := "Would you like to start a new game? Press (Y) for yes, any other key to exit"
      Expect(messages.NewGamePrompt()).To(Equal(expectedNewGameText))
    })

    It("has invalid choice response", func() {
      expectedInvalidChoiceText := "Whoops, that choice is invalid! Try Again"
      Expect(messages.InvalidChoiceResponse()).To(Equal(expectedInvalidChoiceText))
    })

    It("has board type prompt text", func() {
      expectedBoardTypeText := "Choose Board Type:"
      Expect(messages.BoardTypePrompt()).To(Equal(expectedBoardTypeText))
    })

    It("has board type response text", func() {
      expectedBoardTypeText := "%v. %v\n"
      Expect(messages.BoardTypesResponse()).To(Equal(expectedBoardTypeText))
    })

    It("has player type prompt text", func() {
      expectedPlayerTypeText := "Choose Type for: %v\n"
      Expect(messages.PlayerTypePrompt()).To(Equal(expectedPlayerTypeText))
    })

    It("has player type response text", func() {
      expectedPlayerTypeText := "%v. %v\n"
      Expect(messages.PlayerTypesResponse()).To(Equal(expectedPlayerTypeText))
    })

    It("has vertical divider text", func() {
      expectedVerticalDividerText := " | "
      Expect(messages.VerticalDivider()).To(Equal(expectedVerticalDividerText))
    })

    It("has horizontal divider text", func() {
      expectedHorizontalDividerText := "-"
      Expect(messages.HorizontalDivider()).To(Equal(expectedHorizontalDividerText))
    })

    It("has empty space text", func() {
      expectedEmptySpaceText := " "
      Expect(messages.EmptySpot()).To(Equal(expectedEmptySpaceText))
    })

    It("has player x symbol", func() {
      expectedPlayerXText := "X"
      Expect(messages.PlayerXSymbol()).To(Equal(expectedPlayerXText))
    })

    It("has player o symbol", func() {
      expectedPlayerOText := "O"
      Expect(messages.PlayerOSymbol()).To(Equal(expectedPlayerOText))
    })

    It("has spot width", func() {
      expectedSpotWidth := 3
      Expect(messages.SpotWidth()).To(Equal(expectedSpotWidth))
    })

    It("has yes status", func() {
      expectedYesOption := "y"
      Expect(messages.YesOption()).To(Equal(expectedYesOption))
    })

})


