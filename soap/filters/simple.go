package filters

import "github.com/hooklift/gowsdl/soap"

type SimpleOperators string

const (
	SimpleOperatorsEquals SimpleOperators = "equals"

	SimpleOperatorsNotEquals SimpleOperators = "notEquals"

	SimpleOperatorsGreaterThan SimpleOperators = "greaterThan"

	SimpleOperatorsLessThan SimpleOperators = "lessThan"

	SimpleOperatorsIsNull SimpleOperators = "isNull"

	SimpleOperatorsIsNotNull SimpleOperators = "isNotNull"

	SimpleOperatorsGreaterThanOrEqual SimpleOperators = "greaterThanOrEqual"

	SimpleOperatorsLessThanOrEqual SimpleOperators = "lessThanOrEqual"

	SimpleOperatorsBetween SimpleOperators = "between"

	SimpleOperatorsIN SimpleOperators = "IN"

	SimpleOperatorsLike SimpleOperators = "like"

	SimpleOperatorsExistsInString SimpleOperators = "existsInString"

	SimpleOperatorsExistsInStringAsAWord SimpleOperators = "existsInStringAsAWord"

	SimpleOperatorsNotExistsInString SimpleOperators = "notExistsInString"

	SimpleOperatorsBeginsWith SimpleOperators = "beginsWith"

	SimpleOperatorsEndsWith SimpleOperators = "endsWith"

	SimpleOperatorsContains SimpleOperators = "contains"

	SimpleOperatorsNotContains SimpleOperators = "notContains"

	SimpleOperatorsIsAnniversary SimpleOperators = "isAnniversary"

	SimpleOperatorsIsNotAnniversary SimpleOperators = "isNotAnniversary"

	SimpleOperatorsGreaterThanAnniversary SimpleOperators = "greaterThanAnniversary"

	SimpleOperatorsLessThanAnniversary SimpleOperators = "lessThanAnniversary"
)

type SimpleFilterPart struct {
	Property string `xml:"Property,omitempty" json:"Property,omitempty"`

	SimpleOperator SimpleOperators `xml:"SimpleOperator,omitempty" json:"SimpleOperator,omitempty"`

	Value []string `xml:"Value,omitempty" json:"Value,omitempty"`

	DateValue []soap.XSDDateTime `xml:"DateValue,omitempty" json:"DateValue,omitempty"`
}

func Equals(property string, value string) *SimpleFilterPart {
	return &SimpleFilterPart{
		Property:       property,
		SimpleOperator: SimpleOperatorsEquals,
		Value:          []string{value},
	}
}
