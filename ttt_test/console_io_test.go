package ttt_test

import (
	"bytes"
	. "github.com/andrewzures/tictactoe/ttt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Console UI", func() {
	var writer bytes.Buffer
	var reader bytes.Buffer
	var inOut InOut

	BeforeEach(func() {
		inOut = InOut(ConsoleIO{&writer, &reader})
	})

	Context("when receiving input from user", func() {

		It("reads from console", func() {
			SetMockInput(&reader, "Test\nConsole\nInput\n")
			Expect(inOut.Read()).To(Equal("Test"))
			Expect(inOut.Read()).To(Equal("Console"))
			Expect(inOut.Read()).To(Equal("Input"))
		})

		It("provides Println option", func() {
			inOut.Println("Test String")
			Expect(writer.String()).To(ContainSubstring("Test String"))
		})

		It("provides Printf option", func() {
			inOut.Printf("Test String %v\n", 123)
			Expect(writer.String()).To(ContainSubstring("Test String 123"))
		})
	})

	AfterEach(func() {
		writer.Reset()
		reader.Reset()
	})

})
