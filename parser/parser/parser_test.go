package parser

import (
	"parser/grammar"
	"parser/stack"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ParserTestSuite struct {
	suite.Suite
	g *grammar.Grammar
	p *parser
	w []grammar.Terminal
}

func (s *ParserTestSuite) SetupTest() {
	s.g = grammar.New("../grammar.txt")
	s.p = New()
	s.w = []grammar.Terminal{"aab"}
}

func (s *ParserTestSuite) TeardownTest() {
	s.p.Reset()
}

func (s *ParserTestSuite) TestExpand() {
	s.normalState()
	s.p.is.Push(s.g.StartingSymbol)

	s.p.expand(*s.g, s.w)

	assert.Equal(s.T(), s.p.state, Normal)
	assert.Equal(s.T(), s.p.index, 1)

	vws, err := s.p.ws.Top()
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), vws, s.g.StartingSymbol)

	vop := s.p.ops.Front().Value
	assert.Equal(s.T(), vop, 0)

	p1 := s.g.Productions[s.g.StartingSymbol].GetProd(0)
	for i := 0; i < len(p1); i++ {
		x, err := s.p.is.Pop()
		assert.Nil(s.T(), err)
		assert.Equal(s.T(), p1[i], x)
	}
}

func (s *ParserTestSuite) TestAdvance() {
	s.normalState()
	s.p.is.Push(s.w[0])
	s.p.advance(*s.g, s.w)

	assert.Equal(s.T(), Normal, s.p.state)
	assert.Equal(s.T(), 2, s.p.index)

	vws, err := s.p.ws.Top()
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), s.w[0], vws)

	assert.True(s.T(), s.p.is.Empty())
}

func (s *ParserTestSuite) TestMomentaryInsuccess() {
	s.normalState()
	s.p.is.Push(grammar.Terminal("_"))

	s.p.momentaryInsucces()

	assert.Equal(s.T(), Back, s.p.state)
	assert.Equal(s.T(), 1, s.p.index)
	assert.True(s.T(), s.p.ws.Empty())

	val, err := s.p.is.Top()
	assert.Nil(s.T(), err)

	assert.Equal(s.T(), grammar.Terminal("_"), val)
}

func (s *ParserTestSuite) TestBackCanGoBack() {
	s.normalState()
	s.p.state = Back
	s.p.index = 1
	s.p.ws.Push(grammar.Terminal("a"))

	s.p.back(*s.g, s.w)

	assert.Equal(s.T(), Back, s.p.state)
	assert.Equal(s.T(), 0, s.p.index)
	assert.True(s.T(), s.p.ws.Empty())

	val, err := s.p.is.Top()
	assert.Nil(s.T(), err)

	assert.Equal(s.T(), grammar.Terminal("a"), val)
}

func (s *ParserTestSuite) TestBackCannotGoBack() {
	s.normalState()
	s.p.state = Back
	s.p.index = 0
	s.p.ws.Push(grammar.Terminal("a"))

	s.p.back(*s.g, s.w)

	assert.Equal(s.T(), Error, s.p.state)
}

func (s *ParserTestSuite) TestAnotherTryOtherProductionExists() {
	s.normalState()
	s.p.state = Back
	s.p.ws.Push(s.g.StartingSymbol)
	s.p.ops.PushFront(0)

	p1 := s.g.Productions[s.g.StartingSymbol].GetProd(0)

	for i := len(p1) - 1; i >= 0; i-- {
		s.p.is.Push(p1[i])
	}

	s.p.anotherTry(*s.g, s.w)

	p2 := s.g.Productions[s.g.StartingSymbol].GetProd(1)

	assert.Equal(s.T(), Normal, s.p.state)

	val, err := s.p.ws.Top()
	assert.Nil(s.T(), err)

	assert.Equal(s.T(), 1, s.p.index)
	assert.Equal(s.T(), s.g.StartingSymbol, val)
	assert.Equal(s.T(), 1, s.p.ops.Front().Value)

	isTop, err := s.p.is.Peek(uint(len(p2)))

	assert.Nil(s.T(), err)

	for i := 0; i < len(p2); i++ {
		assert.Equal(s.T(), p2[i], isTop[i])
	}

}

func (s *ParserTestSuite) TestAnotherTryNoOtherProductions() {
	s.normalState()
	s.p.state = Back
	s.p.ws.Push(s.g.StartingSymbol)
	s.p.ops.PushFront(1)

	p2 := s.g.Productions[s.g.StartingSymbol].GetProd(1)

	for i := len(p2) - 1; i > 0; i-- {
		s.p.is.Push(p2[i])
	}

	s.p.anotherTry(*s.g, s.w)

	assert.Equal(s.T(), Back, s.p.state)
	assert.True(s.T(), s.p.ws.Empty())
	assert.Equal(s.T(), 1, s.p.index)

	val, err := s.p.is.Top()
	assert.Nil(s.T(), err)

	assert.Equal(s.T(), s.g.StartingSymbol, val)

}

func (s *ParserTestSuite) TestAnotherTryCannotGoBackAnymore() {

}

func TestParserTestSuite(t *testing.T) {
	suite.Run(t, new(ParserTestSuite))
}

func (s *ParserTestSuite) normalState() {
	s.p.state = Normal
	s.p.index = 1
	s.p.ws = stack.New()
	s.p.is = stack.New()
}
