package parser

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/slonegd/go-st/source"
)

type (
	// для пробрасывания по дереву доп объектов (логгер, исходный код и тп)
	CustomContext struct {
		antlr.BaseParserRuleContext
		source.Source
		Logger
		Name string
	}
	Logger interface {
		Debug(string, ...any)
	}
	CustomContextSetter interface {
		Set(logger Logger, source source.Source)
	}
)

// вызывается в сгенерированном коде
func NewCustomContext(parent antlr.ParserRuleContext, invokingState int) *CustomContext {
	return &CustomContext{BaseParserRuleContext: *antlr.NewBaseParserRuleContext(parent, invokingState)}
}

// хелпер
func (c *CustomContext) MakeSnippet() string {
	return c.Snippet(c.GetStart().GetLine(), c.GetStart().GetColumn(), 0)
}

// для интерфейса
func (c *CustomContext) GetCustomContext() *CustomContext {
	return c
}

// Прокидываем всё по контексту
func (c *CustomContext) Set(logger Logger, source source.Source) {
	c.Logger = logger
	c.Source = source
	// Рекурсивно прокидываем логгер в дочерние контексты
	for _, child := range c.GetChildren() {
		if ctx, ok := child.(antlr.ParseTree); ok {
			if childCtx, ok := ctx.(CustomContextSetter); ok {
				childCtx.Set(logger, source)
			}
		}
	}
}
