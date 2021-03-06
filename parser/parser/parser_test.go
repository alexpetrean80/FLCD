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

}

func (s *ParserTestSuite) TestBack() {

}

func (s *ParserTestSuite) TestError() {

}

func (s *ParserTestSuite) TestAnotherTry() {

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
