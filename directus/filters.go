package directus

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strconv"
	"strings"
)

type FilterOperation string

const (
	FILTER_EQ                = FilterOperation("_eq")
	FILTER_NEQ               = FilterOperation("_neq")
	FILTER_CONTAINS          = FilterOperation("_contains")
	FILTER_LESS              = FilterOperation("_lt")
	FILTER_LESS_OR_EQUALS    = FilterOperation("_lte")
	FILTER_GREATER_OR_EQUALS = FilterOperation("_gte")
	FILTER_GREATER           = FilterOperation("_gt")

	FILTER_OR  = FilterOperation("_or")
	FILTER_AND = FilterOperation("_and")

	VAR_TIME_NOW     = "$NOW"
	VAR_CURRENT_USER = "$CURRENT_USER"
	VAR_CURRENT_ROLE = "$CURRENT_ROLE"
)

func getOperator(op token.Token) (FilterOperation, error) {
	switch op {
	case token.EQL:
		return FILTER_EQ, nil
	case token.NEQ:
		return FILTER_NEQ, nil
	case token.LSS:
		return FILTER_LESS, nil
	case token.LEQ:
		return FILTER_LESS_OR_EQUALS, nil
	case token.GTR:
		return FILTER_GREATER, nil
	case token.GEQ:
		return FILTER_GREATER_OR_EQUALS, nil
	default:
		return FilterOperation(""), fmt.Errorf("Failed to get operator from ast tree")
	}
}
func getOperand(op ast.Expr) (string, token.Token, error) {
	switch op.(type) {
	case *ast.BasicLit:
		operand := op.(*ast.BasicLit)
		litname := strings.ReplaceAll(operand.Value, "\"", "")
		litname = strings.ReplaceAll(litname, "'", "")
		return litname, operand.Kind, nil
	case *ast.SelectorExpr:
		innerOp, _, err := getOperand(op.(*ast.SelectorExpr).X)
		if err != nil {
			return "", 0, err
		}
		return fmt.Sprintf("%s.%s", innerOp, op.(*ast.SelectorExpr).Sel), 0, nil
	case *ast.Ident:
		return op.(*ast.Ident).Name, 0, nil
	}
	return "", 0, fmt.Errorf("Failed to get operand from ast tree")
}

func (h *CollectionQuery[K, V]) buildWhereFilters() (string, error) {
	fmap := make(map[string]map[FilterOperation]any)

	for _, filterString := range h.whereFilters {

		binaryExpr, err := parser.ParseExpr(filterString)
		if err != nil {
			return "", err
		}
		op, err := getOperator(binaryExpr.(*ast.BinaryExpr).Op)
		if err != nil {
			return "", err
		}
		left, _, err := getOperand(binaryExpr.(*ast.BinaryExpr).X)
		if err != nil {
			return "", err
		}
		right, rkind, err := getOperand(binaryExpr.(*ast.BinaryExpr).Y)
		if err != nil {
			return "", err
		}
		switch rkind {
		case token.STRING:
			fmap[left] = map[FilterOperation]any{op: right}
			break
		case token.INT:
			v, _ := strconv.ParseInt(right, 10, 32)
			fmap[left] = map[FilterOperation]any{op: int32(v)}
			break
		case token.FLOAT:
			v, _ := strconv.ParseFloat(right, 32)
			fmap[left] = map[FilterOperation]any{op: float32(v)}
			break
		}

	}
	result, err := json.Marshal(fmap)
	if err != nil {
		return "", err
	}
	return string(result), nil
}

func (h *CollectionQuery[K, V]) buildSelectors() string {
	fields := strings.Join(h.fieldSelectors, ",")
	return fields
}
