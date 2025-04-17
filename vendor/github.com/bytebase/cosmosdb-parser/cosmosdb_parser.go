// Code generated from CosmosDBParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // CosmosDBParser
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type CosmosDBParser struct {
	*antlr.BaseParser
}

var CosmosDBParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func cosmosdbparserParserInit() {
	staticData := &CosmosDBParserParserStaticData
	staticData.LiteralNames = []string{
		"", "'*'", "'SELECT'", "'FROM'",
	}
	staticData.SymbolicNames = []string{
		"", "MULTIPLY_OPERATOR", "SELECT_SYMBOL", "FROM_SYMBOL", "IDENTIFIER",
		"WHITESPACE",
	}
	staticData.RuleNames = []string{
		"root", "select", "select_clause", "select_specification", "from_clause",
		"from_specification", "from_source", "container_expression", "container_name",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 5, 41, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7, 4,
		2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 1, 0, 1, 0, 1, 0, 1, 1,
		1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5,
		1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 0, 0, 9, 0, 2, 4, 6, 8, 10, 12,
		14, 16, 0, 0, 31, 0, 18, 1, 0, 0, 0, 2, 21, 1, 0, 0, 0, 4, 24, 1, 0, 0,
		0, 6, 27, 1, 0, 0, 0, 8, 29, 1, 0, 0, 0, 10, 32, 1, 0, 0, 0, 12, 34, 1,
		0, 0, 0, 14, 36, 1, 0, 0, 0, 16, 38, 1, 0, 0, 0, 18, 19, 3, 2, 1, 0, 19,
		20, 5, 0, 0, 1, 20, 1, 1, 0, 0, 0, 21, 22, 3, 4, 2, 0, 22, 23, 3, 8, 4,
		0, 23, 3, 1, 0, 0, 0, 24, 25, 5, 2, 0, 0, 25, 26, 3, 6, 3, 0, 26, 5, 1,
		0, 0, 0, 27, 28, 5, 1, 0, 0, 28, 7, 1, 0, 0, 0, 29, 30, 5, 3, 0, 0, 30,
		31, 3, 10, 5, 0, 31, 9, 1, 0, 0, 0, 32, 33, 3, 12, 6, 0, 33, 11, 1, 0,
		0, 0, 34, 35, 3, 14, 7, 0, 35, 13, 1, 0, 0, 0, 36, 37, 3, 16, 8, 0, 37,
		15, 1, 0, 0, 0, 38, 39, 5, 4, 0, 0, 39, 17, 1, 0, 0, 0, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// CosmosDBParserInit initializes any static state used to implement CosmosDBParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewCosmosDBParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func CosmosDBParserInit() {
	staticData := &CosmosDBParserParserStaticData
	staticData.once.Do(cosmosdbparserParserInit)
}

// NewCosmosDBParser produces a new parser instance for the optional input antlr.TokenStream.
func NewCosmosDBParser(input antlr.TokenStream) *CosmosDBParser {
	CosmosDBParserInit()
	this := new(CosmosDBParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &CosmosDBParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "CosmosDBParser.g4"

	return this
}

// CosmosDBParser tokens.
const (
	CosmosDBParserEOF               = antlr.TokenEOF
	CosmosDBParserMULTIPLY_OPERATOR = 1
	CosmosDBParserSELECT_SYMBOL     = 2
	CosmosDBParserFROM_SYMBOL       = 3
	CosmosDBParserIDENTIFIER        = 4
	CosmosDBParserWHITESPACE        = 5
)

// CosmosDBParser rules.
const (
	CosmosDBParserRULE_root                 = 0
	CosmosDBParserRULE_select               = 1
	CosmosDBParserRULE_select_clause        = 2
	CosmosDBParserRULE_select_specification = 3
	CosmosDBParserRULE_from_clause          = 4
	CosmosDBParserRULE_from_specification   = 5
	CosmosDBParserRULE_from_source          = 6
	CosmosDBParserRULE_container_expression = 7
	CosmosDBParserRULE_container_name       = 8
)

// IRootContext is an interface to support dynamic dispatch.
type IRootContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Select_() ISelectContext
	EOF() antlr.TerminalNode

	// IsRootContext differentiates from other interfaces.
	IsRootContext()
}

type RootContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRootContext() *RootContext {
	var p = new(RootContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_root
	return p
}

func InitEmptyRootContext(p *RootContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_root
}

func (*RootContext) IsRootContext() {}

func NewRootContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RootContext {
	var p = new(RootContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CosmosDBParserRULE_root

	return p
}

func (s *RootContext) GetParser() antlr.Parser { return s.parser }

func (s *RootContext) Select_() ISelectContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelectContext)
}

func (s *RootContext) EOF() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserEOF, 0)
}

func (s *RootContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RootContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RootContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.EnterRoot(s)
	}
}

func (s *RootContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.ExitRoot(s)
	}
}

func (s *RootContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CosmosDBParserVisitor:
		return t.VisitRoot(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CosmosDBParser) Root() (localctx IRootContext) {
	localctx = NewRootContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CosmosDBParserRULE_root)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(18)
		p.Select_()
	}
	{
		p.SetState(19)
		p.Match(CosmosDBParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISelectContext is an interface to support dynamic dispatch.
type ISelectContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Select_clause() ISelect_clauseContext
	From_clause() IFrom_clauseContext

	// IsSelectContext differentiates from other interfaces.
	IsSelectContext()
}

type SelectContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectContext() *SelectContext {
	var p = new(SelectContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_select
	return p
}

func InitEmptySelectContext(p *SelectContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_select
}

func (*SelectContext) IsSelectContext() {}

func NewSelectContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectContext {
	var p = new(SelectContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CosmosDBParserRULE_select

	return p
}

func (s *SelectContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectContext) Select_clause() ISelect_clauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelect_clauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelect_clauseContext)
}

func (s *SelectContext) From_clause() IFrom_clauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFrom_clauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFrom_clauseContext)
}

func (s *SelectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.EnterSelect(s)
	}
}

func (s *SelectContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.ExitSelect(s)
	}
}

func (s *SelectContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CosmosDBParserVisitor:
		return t.VisitSelect(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CosmosDBParser) Select_() (localctx ISelectContext) {
	localctx = NewSelectContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CosmosDBParserRULE_select)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(21)
		p.Select_clause()
	}
	{
		p.SetState(22)
		p.From_clause()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISelect_clauseContext is an interface to support dynamic dispatch.
type ISelect_clauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SELECT_SYMBOL() antlr.TerminalNode
	Select_specification() ISelect_specificationContext

	// IsSelect_clauseContext differentiates from other interfaces.
	IsSelect_clauseContext()
}

type Select_clauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelect_clauseContext() *Select_clauseContext {
	var p = new(Select_clauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_select_clause
	return p
}

func InitEmptySelect_clauseContext(p *Select_clauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_select_clause
}

func (*Select_clauseContext) IsSelect_clauseContext() {}

func NewSelect_clauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Select_clauseContext {
	var p = new(Select_clauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CosmosDBParserRULE_select_clause

	return p
}

func (s *Select_clauseContext) GetParser() antlr.Parser { return s.parser }

func (s *Select_clauseContext) SELECT_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserSELECT_SYMBOL, 0)
}

func (s *Select_clauseContext) Select_specification() ISelect_specificationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelect_specificationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelect_specificationContext)
}

func (s *Select_clauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Select_clauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Select_clauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.EnterSelect_clause(s)
	}
}

func (s *Select_clauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.ExitSelect_clause(s)
	}
}

func (s *Select_clauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CosmosDBParserVisitor:
		return t.VisitSelect_clause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CosmosDBParser) Select_clause() (localctx ISelect_clauseContext) {
	localctx = NewSelect_clauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, CosmosDBParserRULE_select_clause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(24)
		p.Match(CosmosDBParserSELECT_SYMBOL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(25)
		p.Select_specification()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISelect_specificationContext is an interface to support dynamic dispatch.
type ISelect_specificationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MULTIPLY_OPERATOR() antlr.TerminalNode

	// IsSelect_specificationContext differentiates from other interfaces.
	IsSelect_specificationContext()
}

type Select_specificationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelect_specificationContext() *Select_specificationContext {
	var p = new(Select_specificationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_select_specification
	return p
}

func InitEmptySelect_specificationContext(p *Select_specificationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_select_specification
}

func (*Select_specificationContext) IsSelect_specificationContext() {}

func NewSelect_specificationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Select_specificationContext {
	var p = new(Select_specificationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CosmosDBParserRULE_select_specification

	return p
}

func (s *Select_specificationContext) GetParser() antlr.Parser { return s.parser }

func (s *Select_specificationContext) MULTIPLY_OPERATOR() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserMULTIPLY_OPERATOR, 0)
}

func (s *Select_specificationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Select_specificationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Select_specificationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.EnterSelect_specification(s)
	}
}

func (s *Select_specificationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.ExitSelect_specification(s)
	}
}

func (s *Select_specificationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CosmosDBParserVisitor:
		return t.VisitSelect_specification(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CosmosDBParser) Select_specification() (localctx ISelect_specificationContext) {
	localctx = NewSelect_specificationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, CosmosDBParserRULE_select_specification)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(27)
		p.Match(CosmosDBParserMULTIPLY_OPERATOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFrom_clauseContext is an interface to support dynamic dispatch.
type IFrom_clauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FROM_SYMBOL() antlr.TerminalNode
	From_specification() IFrom_specificationContext

	// IsFrom_clauseContext differentiates from other interfaces.
	IsFrom_clauseContext()
}

type From_clauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFrom_clauseContext() *From_clauseContext {
	var p = new(From_clauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_from_clause
	return p
}

func InitEmptyFrom_clauseContext(p *From_clauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_from_clause
}

func (*From_clauseContext) IsFrom_clauseContext() {}

func NewFrom_clauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *From_clauseContext {
	var p = new(From_clauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CosmosDBParserRULE_from_clause

	return p
}

func (s *From_clauseContext) GetParser() antlr.Parser { return s.parser }

func (s *From_clauseContext) FROM_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserFROM_SYMBOL, 0)
}

func (s *From_clauseContext) From_specification() IFrom_specificationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFrom_specificationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFrom_specificationContext)
}

func (s *From_clauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *From_clauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *From_clauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.EnterFrom_clause(s)
	}
}

func (s *From_clauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.ExitFrom_clause(s)
	}
}

func (s *From_clauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CosmosDBParserVisitor:
		return t.VisitFrom_clause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CosmosDBParser) From_clause() (localctx IFrom_clauseContext) {
	localctx = NewFrom_clauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, CosmosDBParserRULE_from_clause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(29)
		p.Match(CosmosDBParserFROM_SYMBOL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(30)
		p.From_specification()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFrom_specificationContext is an interface to support dynamic dispatch.
type IFrom_specificationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	From_source() IFrom_sourceContext

	// IsFrom_specificationContext differentiates from other interfaces.
	IsFrom_specificationContext()
}

type From_specificationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFrom_specificationContext() *From_specificationContext {
	var p = new(From_specificationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_from_specification
	return p
}

func InitEmptyFrom_specificationContext(p *From_specificationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_from_specification
}

func (*From_specificationContext) IsFrom_specificationContext() {}

func NewFrom_specificationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *From_specificationContext {
	var p = new(From_specificationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CosmosDBParserRULE_from_specification

	return p
}

func (s *From_specificationContext) GetParser() antlr.Parser { return s.parser }

func (s *From_specificationContext) From_source() IFrom_sourceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFrom_sourceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFrom_sourceContext)
}

func (s *From_specificationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *From_specificationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *From_specificationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.EnterFrom_specification(s)
	}
}

func (s *From_specificationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.ExitFrom_specification(s)
	}
}

func (s *From_specificationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CosmosDBParserVisitor:
		return t.VisitFrom_specification(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CosmosDBParser) From_specification() (localctx IFrom_specificationContext) {
	localctx = NewFrom_specificationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, CosmosDBParserRULE_from_specification)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(32)
		p.From_source()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFrom_sourceContext is an interface to support dynamic dispatch.
type IFrom_sourceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Container_expression() IContainer_expressionContext

	// IsFrom_sourceContext differentiates from other interfaces.
	IsFrom_sourceContext()
}

type From_sourceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFrom_sourceContext() *From_sourceContext {
	var p = new(From_sourceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_from_source
	return p
}

func InitEmptyFrom_sourceContext(p *From_sourceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_from_source
}

func (*From_sourceContext) IsFrom_sourceContext() {}

func NewFrom_sourceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *From_sourceContext {
	var p = new(From_sourceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CosmosDBParserRULE_from_source

	return p
}

func (s *From_sourceContext) GetParser() antlr.Parser { return s.parser }

func (s *From_sourceContext) Container_expression() IContainer_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IContainer_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IContainer_expressionContext)
}

func (s *From_sourceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *From_sourceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *From_sourceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.EnterFrom_source(s)
	}
}

func (s *From_sourceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.ExitFrom_source(s)
	}
}

func (s *From_sourceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CosmosDBParserVisitor:
		return t.VisitFrom_source(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CosmosDBParser) From_source() (localctx IFrom_sourceContext) {
	localctx = NewFrom_sourceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, CosmosDBParserRULE_from_source)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(34)
		p.Container_expression()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IContainer_expressionContext is an interface to support dynamic dispatch.
type IContainer_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Container_name() IContainer_nameContext

	// IsContainer_expressionContext differentiates from other interfaces.
	IsContainer_expressionContext()
}

type Container_expressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContainer_expressionContext() *Container_expressionContext {
	var p = new(Container_expressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_container_expression
	return p
}

func InitEmptyContainer_expressionContext(p *Container_expressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_container_expression
}

func (*Container_expressionContext) IsContainer_expressionContext() {}

func NewContainer_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Container_expressionContext {
	var p = new(Container_expressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CosmosDBParserRULE_container_expression

	return p
}

func (s *Container_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Container_expressionContext) Container_name() IContainer_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IContainer_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IContainer_nameContext)
}

func (s *Container_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Container_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Container_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.EnterContainer_expression(s)
	}
}

func (s *Container_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.ExitContainer_expression(s)
	}
}

func (s *Container_expressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CosmosDBParserVisitor:
		return t.VisitContainer_expression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CosmosDBParser) Container_expression() (localctx IContainer_expressionContext) {
	localctx = NewContainer_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, CosmosDBParserRULE_container_expression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(36)
		p.Container_name()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IContainer_nameContext is an interface to support dynamic dispatch.
type IContainer_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsContainer_nameContext differentiates from other interfaces.
	IsContainer_nameContext()
}

type Container_nameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContainer_nameContext() *Container_nameContext {
	var p = new(Container_nameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_container_name
	return p
}

func InitEmptyContainer_nameContext(p *Container_nameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_container_name
}

func (*Container_nameContext) IsContainer_nameContext() {}

func NewContainer_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Container_nameContext {
	var p = new(Container_nameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CosmosDBParserRULE_container_name

	return p
}

func (s *Container_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Container_nameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserIDENTIFIER, 0)
}

func (s *Container_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Container_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Container_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.EnterContainer_name(s)
	}
}

func (s *Container_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.ExitContainer_name(s)
	}
}

func (s *Container_nameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CosmosDBParserVisitor:
		return t.VisitContainer_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CosmosDBParser) Container_name() (localctx IContainer_nameContext) {
	localctx = NewContainer_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, CosmosDBParserRULE_container_name)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(38)
		p.Match(CosmosDBParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
