package sqlbuilder

func Literal(text string) *LiteralExpr {
	return &LiteralExpr{text: text}
}

type LiteralExpr struct {
	text string
}

func (l *LiteralExpr) AsExpr(s *Serializer) {
	s.D(l.text)
}

func (e *LiteralExpr) As(alias string) *ColumnAlias {
	return AliasColumn(e, alias)
}
