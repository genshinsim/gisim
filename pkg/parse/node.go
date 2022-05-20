package parse

import (
	"strconv"
	"strings"
)

type Node interface {
	String() string
	// Copy does a deep copy of the Node and all its components.
	// To avoid type assertions, some XxxNodes also have specialized
	// CopyXxx methods that return *XxxNode.
	Copy() Node
	Position() Pos // byte position of start of node in full original input string
	// writeTo writes the String output to the builder.
	writeTo(*strings.Builder)
}

type Pos int

func (p Pos) Position() Pos {
	return p
}

// Stmt.

type Stmt interface {
	Node
	stmtNode()
}

type (

	// BlockStmt represents a brace statement list
	BlockStmt struct {
		List []Node
		Pos
	}

	// ActionStmt represents a sim action; Does not produce a value
	ActionStmt struct {
		Pos
	}

	// AssignStmt represents assigning of a value to a previously declared variable
	AssignStmt struct {
		Pos
		Ident Token
		Val   Expr
	}

	// LetStmt represents a variable assignment. Number only
	LetStmt struct {
		Pos
		Ident Token
		Val   Expr
	}

	// ReturnStmt represents return <expr>.
	ReturnStmt struct {
		Pos
		Val Expr
	}

	// CtrlStmt represents continue, break, and fallthrough
	CtrlStmt struct {
		Pos
		Typ CtrlTyp
	}

	// IfStmt represents an if block
	IfStmt struct {
		Pos
		Condition Expr       //TODO: this should be an expr?
		IfBlock   *BlockStmt // What to execute if true
		ElseBlock *BlockStmt // What to execute if false
	}

	// SwitchStmt represent a switch block
	SwitchStmt struct {
		Pos
		Condition Expr // the condition to switch on
		Cases     []*CaseStmt
		Default   *BlockStmt // default case
	}

	// CaseStmt represents a case in a switch block
	CaseStmt struct {
		Pos
		Condition Expr
		Body      *BlockStmt
	}

	// A FnStmt node represents a function
	FnStmt struct {
		Pos
		FunVal Token
		Args   []*Ident
		Body   *BlockStmt
	}
	// WhileStmt represents a while block
	WhileStmt struct {
		Pos
		Condition  Expr       //TODO: this should be an expr?
		WhileBlock *BlockStmt // What to execute if true
	}
)

type CtrlTyp int

const (
	InvalidCtrl CtrlTyp = iota
	CtrlBreak
	CtrlContinue
	CtrlFallthrough
)

// stmtNode()
func (*BlockStmt) stmtNode()  {}
func (*AssignStmt) stmtNode() {}
func (*LetStmt) stmtNode()    {}
func (*CtrlStmt) stmtNode()   {}
func (*ReturnStmt) stmtNode() {}
func (*IfStmt) stmtNode()     {}
func (*SwitchStmt) stmtNode() {}
func (*CaseStmt) stmtNode()   {}
func (*FnStmt) stmtNode()     {}
func (*WhileStmt) stmtNode()  {}

// BlockStmt.
func newBlockStmt(pos Pos) *BlockStmt {
	return &BlockStmt{Pos: pos}
}
func (b *BlockStmt) append(n Node) {
	b.List = append(b.List, n)
}

func (b *BlockStmt) String() string {
	var sb strings.Builder
	b.writeTo(&sb)
	return sb.String()
}

func (b *BlockStmt) writeTo(sb *strings.Builder) {
	for _, n := range b.List {
		n.writeTo(sb)
		sb.WriteString(";\n")
	}
}

func (b *BlockStmt) CopyBlock() *BlockStmt {
	if b == nil {
		return b
	}
	n := newBlockStmt(b.Pos)
	for _, elem := range b.List {
		n.append(elem.Copy())
	}
	return n
}

func (b *BlockStmt) Copy() Node {
	return b.CopyBlock()
}

// AssignStmt.

func (a *AssignStmt) String() string {
	var sb strings.Builder
	a.writeTo(&sb)
	return sb.String()
}

func (a *AssignStmt) writeTo(sb *strings.Builder) {
	sb.WriteString(a.Ident.String())
	sb.WriteString(" = ")
	a.Val.writeTo(sb)
}

func (a *AssignStmt) CopyAssign() *AssignStmt {
	if a == nil {
		return a
	}
	n := &AssignStmt{
		Pos:   a.Pos,
		Ident: a.Ident,
	}
	n.Val = a.Val.CopyExpr()
	return n
}

func (a *AssignStmt) Copy() Node {
	return a.CopyAssign()
}

// LetStmt.

func (l *LetStmt) String() string {
	var sb strings.Builder
	l.writeTo(&sb)
	return sb.String()
}

func (l *LetStmt) writeTo(sb *strings.Builder) {
	sb.WriteString("let ")
	sb.WriteString(l.Ident.String())
	sb.WriteString(" = ")
	if l.Val != nil {

		l.Val.writeTo(sb)
	}
}

func (l *LetStmt) CopyLet() *LetStmt {
	if l == nil {
		return l
	}
	n := &LetStmt{
		Pos:   l.Pos,
		Ident: l.Ident,
	}
	n.Val = l.Val.CopyExpr()
	return n
}

func (l *LetStmt) Copy() Node {
	return l.CopyLet()
}

// ReturnStmt.

func (r *ReturnStmt) String() string {
	var sb strings.Builder
	r.writeTo(&sb)
	return sb.String()
}

func (r *ReturnStmt) writeTo(sb *strings.Builder) {
	sb.WriteString("return ")
	r.Val.writeTo(sb)
}

func (r *ReturnStmt) CopyReturn() *ReturnStmt {
	if r == nil {
		return r
	}
	n := &ReturnStmt{
		Pos: r.Pos,
	}
	n.Val = r.Val.CopyExpr()
	return n
}

func (r *ReturnStmt) Copy() Node {
	return r.CopyReturn()
}

// CtrlStmt.

func (c *CtrlStmt) String() string {
	var sb strings.Builder
	c.writeTo(&sb)
	return sb.String()
}

func (c *CtrlStmt) writeTo(sb *strings.Builder) {
	switch c.Typ {
	case CtrlContinue:
		sb.WriteString("continue")
	case CtrlBreak:
		sb.WriteString("break")
	case CtrlFallthrough:
		sb.WriteString("fallthrough")
	}
}

func (c *CtrlStmt) CopyControl() *CtrlStmt {
	if c == nil {
		return c
	}
	n := &CtrlStmt{
		Pos: c.Pos,
		Typ: c.Typ,
	}
	return n
}

func (c *CtrlStmt) Copy() Node {
	return c.CopyControl()
}

// IfStmt.

func (i *IfStmt) String() string {
	var sb strings.Builder
	i.writeTo(&sb)
	return sb.String()
}

func (i *IfStmt) writeTo(sb *strings.Builder) {
	sb.WriteString("if ")
	i.Condition.writeTo(sb)
	sb.WriteString(" {\n")
	i.IfBlock.writeTo(sb)
	sb.WriteString("}")
	if i.ElseBlock != nil {
		sb.WriteString("else {\n")
		sb.WriteString(i.ElseBlock.String())
		sb.WriteString("}")
	}
}

func (i *IfStmt) Copy() Node {
	if i == nil {
		return nil
	}
	return &IfStmt{
		Pos:       i.Pos,
		Condition: i.Condition.CopyExpr(),
		IfBlock:   i.IfBlock.CopyBlock(),
		ElseBlock: i.ElseBlock.CopyBlock(),
	}
}

// SwitchStmt.

func (s *SwitchStmt) String() string {
	var sb strings.Builder
	s.writeTo(&sb)
	return sb.String()
}

func (s *SwitchStmt) writeTo(sb *strings.Builder) {
	sb.WriteString("switch ")
	s.Condition.writeTo(sb)
	sb.WriteString(" {\n")
	for _, v := range s.Cases {
		v.writeTo(sb)
	}
	if s.Default != nil {
		sb.WriteString("default: {\n")
		s.Default.writeTo(sb)
		sb.WriteString("}")
	}
	sb.WriteString("}")
}

func (s *SwitchStmt) CopySwitch() *SwitchStmt {
	if s == nil {
		return nil
	}
	n := &SwitchStmt{
		Pos:       s.Pos,
		Condition: s.Condition.CopyExpr(),
		Cases:     make([]*CaseStmt, 0, len(s.Cases)),
	}
	for i := range s.Cases {
		n.Cases = append(n.Cases, s.Cases[i].CopyCase())
	}
	return n
}

func (s *SwitchStmt) Copy() Node {
	return s.CopySwitch()
}

// CaseStmt.

func (c *CaseStmt) String() string {
	var sb strings.Builder
	c.writeTo(&sb)
	return sb.String()
}

func (c *CaseStmt) writeTo(sb *strings.Builder) {
	sb.WriteString("case ")
	c.Condition.writeTo(sb)
	sb.WriteString(" {\n")
	c.Body.writeTo(sb)
	sb.WriteString("}")
}

func (c *CaseStmt) CopyCase() *CaseStmt {
	if c == nil {
		return nil
	}
	return &CaseStmt{
		Pos:       c.Pos,
		Condition: c.Condition.CopyExpr(),
		Body:      c.Body.CopyBlock(),
	}
}

func (c *CaseStmt) Copy() Node {
	return c.CopyCase()
}

// FnExpr.

func (f *FnStmt) CopyFn() Stmt {
	if f == nil {
		return nil
	}
	n := &FnStmt{
		Pos:    f.Pos,
		FunVal: f.FunVal,
		Body:   f.Body.CopyBlock(),
		Args:   make([]*Ident, 0, len(f.Args)),
	}
	for i := range f.Args {
		n.Args = append(n.Args, f.Args[i].CopyIdent())
	}

	return n
}

func (f *FnStmt) CopyStmt() Stmt {
	return f.CopyFn()
}

func (f *FnStmt) Copy() Node {
	return f.CopyStmt()
}

func (f *FnStmt) String() string {
	var sb strings.Builder
	f.writeTo(&sb)
	return sb.String()
}

func (f *FnStmt) writeTo(sb *strings.Builder) {
	sb.WriteString("fn(")
	for i, v := range f.Args {
		if i > 0 {
			sb.WriteString(", ")
		}
		v.writeTo(sb)
	}
	sb.WriteString(") {\n")
	f.Body.writeTo(sb)
	sb.WriteString("}")
}

// WhileStmt.

func (w *WhileStmt) String() string {
	var sb strings.Builder
	w.writeTo(&sb)
	return sb.String()
}

func (w *WhileStmt) writeTo(sb *strings.Builder) {
	sb.WriteString("while ")
	w.Condition.writeTo(sb)
	sb.WriteString(" {\n")
	w.WhileBlock.writeTo(sb)
	sb.WriteString("}")
}

func (w *WhileStmt) Copy() Node {
	if w == nil {
		return nil
	}
	return &WhileStmt{
		Pos:        w.Pos,
		Condition:  w.Condition.CopyExpr(),
		WhileBlock: w.WhileBlock.CopyBlock(),
	}
}

// Expr.

type Expr interface {
	Node
	exprNode()
	CopyExpr() Expr
}

// An expression is represented by a tree consisting of one or
// more of the following concrete expression nodes
type (
	NumberLit struct {
		Pos
		IntVal   int64
		FloatVal float64
		IsInt    bool
	}

	StringLit struct {
		Pos
		Value float64
	}
	BoolLit struct {
		Pos
		Value float64
	}

	Ident struct {
		Pos
		Value string
	}

	// A CallExpr node represents an expression followed by an argument list.
	CallExpr struct {
		Pos
		Fun  Expr   // function expression
		Args []Expr // function arguments; or nil
	}

	// A UnaryExpr node represents a unary expression.
	UnaryExpr struct {
		Pos
		Op    Token
		Right Expr // operand
	}

	//A BinaryExpr node represents a binary expression i.e. a > b, 1 + 1, etc..
	BinaryExpr struct {
		Pos
		Left  Expr
		Right Expr  // need to evalute to same type as lhs
		Op    Token //should be > itemCompareOP and < itemDot
	}
)

//exprNode()
func (*NumberLit) exprNode()  {}
func (*Ident) exprNode()      {}
func (*CallExpr) exprNode()   {}
func (*UnaryExpr) exprNode()  {}
func (*BinaryExpr) exprNode() {}

// NumberLit.

func (n *NumberLit) CopyExpr() Expr {
	if n == nil {
		return nil
	}
	return &NumberLit{Pos: n.Pos, IntVal: n.IntVal}
}

func (n *NumberLit) Copy() Node {
	return n.CopyExpr()
}

func (n *NumberLit) String() string {
	var sb strings.Builder
	n.writeTo(&sb)
	return sb.String()
}

func (n *NumberLit) writeTo(sb *strings.Builder) {
	if n.IsInt {
		sb.WriteString(strconv.FormatInt(n.IntVal, 10))
	} else {
		sb.WriteString(strconv.FormatFloat(n.FloatVal, 'f', -1, 64))
	}
}

// Ident.

func (i *Ident) CopyIdent() *Ident {
	if i == nil {
		return nil
	}
	return &Ident{Pos: i.Pos, Value: i.Value}
}

func (i *Ident) CopyExpr() Expr {
	return i.CopyIdent()
}

func (i *Ident) Copy() Node {
	return i.CopyIdent()
}

func (b *Ident) String() string {
	var sb strings.Builder
	b.writeTo(&sb)
	return sb.String()
}

func (b *Ident) writeTo(sb *strings.Builder) {
	sb.WriteString(b.Value)
}

// CallExpr.

func (c *CallExpr) CopyFn() Expr {
	if c == nil {
		return nil
	}
	n := &CallExpr{
		Pos:  c.Pos,
		Fun:  c.Fun.CopyExpr(),
		Args: make([]Expr, 0, len(c.Args)),
	}
	for i := range c.Args {
		n.Args = append(n.Args, c.Args[i].CopyExpr())
	}

	return n
}

func (f *CallExpr) CopyExpr() Expr {
	return f.CopyFn()
}

func (f *CallExpr) Copy() Node {
	return f.CopyExpr()
}

func (f *CallExpr) String() string {
	var sb strings.Builder
	f.writeTo(&sb)
	return sb.String()
}

func (b *CallExpr) writeTo(sb *strings.Builder) {
	b.Fun.writeTo(sb)
	sb.WriteString("(")
	for i, v := range b.Args {
		if i > 0 {
			sb.WriteString(", ")
		}
		v.writeTo(sb)
	}
	sb.WriteString(")")
}

// UnaryExpr.

func (u *UnaryExpr) CopyUnaryExpr() *UnaryExpr {
	if u == nil {
		return u
	}
	n := &UnaryExpr{Pos: u.Pos}
	n.Right = u.Right.CopyExpr()
	n.Op = u.Op
	return n
}

func (u *UnaryExpr) CopyExpr() Expr {
	return u.CopyUnaryExpr()
}

func (u *UnaryExpr) Copy() Node {
	return u.CopyUnaryExpr()
}

func (u *UnaryExpr) String() string {
	var sb strings.Builder
	u.writeTo(&sb)
	return sb.String()
}

func (u *UnaryExpr) writeTo(sb *strings.Builder) {
	sb.WriteString("(")
	sb.WriteString(u.Op.String())
	u.Right.writeTo(sb)
	sb.WriteString(")")
}

// BinaryExpr.

func (b *BinaryExpr) CopyBinaryExpr() *BinaryExpr {
	if b == nil {
		return b
	}
	n := &BinaryExpr{Pos: b.Pos}
	n.Left = b.Left.CopyExpr()
	n.Right = b.Right.CopyExpr()
	n.Op = b.Op
	return n
}

func (b *BinaryExpr) CopyExpr() Expr {
	return b.CopyBinaryExpr()
}

func (b *BinaryExpr) Copy() Node {
	return b.CopyBinaryExpr()
}

func (b *BinaryExpr) String() string {
	var sb strings.Builder
	b.writeTo(&sb)
	return sb.String()
}

func (b *BinaryExpr) writeTo(sb *strings.Builder) {
	sb.WriteString("(")
	b.Left.writeTo(sb)
	sb.WriteString(" ")
	sb.WriteString(b.Op.String())
	sb.WriteString(" ")
	b.Right.writeTo(sb)
	sb.WriteString(")")
}
