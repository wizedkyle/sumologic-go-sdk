package types

import (
	"github.com/antihax/optional"
	"time"
)

type DimensionTransformation struct {
	// Only used if TransformationType is set to "AggregateOnTransformation"
	AggregateOn []string `json:"aggregateOn"`
	// This is the base type of all dimension transformations. Valid values are: AggregateOnTransformation and AddOrReplaceTransformation
	TransformationType string `json:"transformationType"`
	// Only used if TransformationType is set to "AddOrReplaceTransformation"
	DimensionToReplace string `json:"dimensionToReplace"`
	// Only used if TransformationType is set to "AddOrReplaceTransformation"
	Value string `json:"value"`
}

type TransformationRuleDefinition struct {
	// Name of the transformation rule.
	Name string `json:"name"`
	// Selector of the transformation rule.
	Selector string `json:"selector"`
	// Dimension transformations of the transformation rule.
	DimensionTransformations []DimensionTransformation `json:"dimensionTransformations,omitempty"`
	// Retention period in days for the transformed metrics that are generated by this rule. The supported retention periods for transformed metrics are 8 days, and 400 days. If no dimension transformations are defined, this value will be set to 0.
	TransformedMetricsRetention int64 `json:"transformedMetricsRetention,omitempty"`
	// Retention period in days for the metrics that are selected by the selector. The supported retention periods for selected metrics are 8 days, 400 days, and 0 (Do not store) if this rule contains dimension transformation.
	Retention int64 `json:"retention"`
}

type TransformationRuleResponse struct {
	// Creation timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format.
	CreatedAt time.Time `json:"createdAt"`
	// Identifier of the user who created the resource.
	CreatedBy string `json:"createdBy"`
	// Last modification timestamp in UTC.
	ModifiedAt time.Time `json:"modifiedAt"`
	// Identifier of the user who last modified the resource.
	ModifiedBy     string                        `json:"modifiedBy"`
	RuleDefinition *TransformationRuleDefinition `json:"ruleDefinition"`
	// True if the rule is enabled.
	Enabled bool `json:"enabled"`
	// Unique identifier for the transformation rule.
	Id string `json:"id"`
}

type TransformationRulesResponse struct {
	// List of transformation rules.
	Data []TransformationRuleResponse `json:"data"`
	// Next continuation token.
	Next string `json:"next,omitempty"`
}

type TransformationRuleRequest struct {
	RuleDefinition *TransformationRuleDefinition `json:"ruleDefinition"`
	// True if the rule is enabled.
	Enabled bool `json:"enabled"`
}

type TransformationRulesOpts struct {
	Limit optional.Int32
	Token optional.String
}
