package models

type Operator string

const (
	OperatorEqual            Operator = "equal"
	OperatorNotEqual         Operator = "not_equal"
	OperatorIn               Operator = "in"
	OperatorNotIn            Operator = "not_in"
	OperatorLess             Operator = "less"
	OperatorLessOrEqual      Operator = "less_or_equal"
	OperatorGreater          Operator = "greater"
	OperatorGreaterOrEqual   Operator = "greater_or_equal"
	OperatorBetween          Operator = "between"
	OperatorNotBetween       Operator = "not_between"
	OperatorBeginsWith       Operator = "begins_with"
	OperatorNotBeginsWith    Operator = "not_begins_with"
	OperatorContains         Operator = "contains"
	OperatorNotContains      Operator = "not_contains"
	OperatorEndsWith         Operator = "ends_with"
	OperatorNotEndsWith      Operator = "not_ends_with"
	OperatorIsNull           Operator = "is_null"
	OperatorIsNotNull        Operator = "is_not_null"
	OperatorIsEmpty          Operator = "is_empty"
	OperatorIsNotEmpty       Operator = "is_not_empty"
	OperatorIsAnyOf          Operator = "is_any_of"
	OperatorIsNoneOf         Operator = "is_none_of"
)

type Rule struct {
	Field       string  `json:"field,omitempty"`
	Value       string  `json:"value,omitempty"`
	Operator    string  `json:"operator,omitempty"`    //
	ValueSource string  `json:"valueSource,omitempty"` // valueSource: "value" or "field"
	Combinator  string  `json:"combinator,omitempty"`
	Not         bool    `json:"not,omitempty"`
	Rules       []*Rule `json:"rules,omitempty"`
}

// rule to sql
func (r *Rule) ToSql() (string, []interface{}) {
	var sql string
	var args []interface{}
	if r.Field != "" && r.Value != "" && r.Operator != "" {
		sql = r.Field + " " + r.Operator + " ?"
		args = append(args, r.Value)
	} else if r.Field != "" && r.ValueSource != "" && r.Operator != "" {
		sql = r.Field + " " + r.Operator + " " + r.ValueSource
	} else if r.Combinator != "" {
		var subSql string
		var subArgs []interface{}
		for _, rule := range r.Rules {
			subSql, subArgs = rule.ToSql()
			if subSql != "" {
				sql += subSql + " " + r.Combinator + " "
				args = append(args, subArgs...)
			}
		}
		if sql != "" {
			sql = "(" + sql[:len(sql)-len(r.Combinator)-2] + ")"
		}
	}
	if r.Not {
		sql = "NOT " + sql
	}
	return sql, args
}
