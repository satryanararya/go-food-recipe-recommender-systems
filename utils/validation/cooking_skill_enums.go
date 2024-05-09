package validation

import (
	"github.com/satryanararya/go-chefbot/constants/enums"
)

func IsValidEnumValue(enumType string, value string) bool {
	switch enumType {
	case "ExperienceYears":
		return isValidExperienceYears(value)
	case "TimeCommitment":
		return isValidTimeCommitment(value)
	case "RecipeComplexity":
		return isValidRecipeComplexity(value)
	case "IngredientDiversity":
		return isValidIngredientDiversity(value)
	default:
		return false
	}
}

func isValidExperienceYears(value string) bool {
	return enums.ExperienceYears(value) == enums.Experience0To2Years ||
		enums.ExperienceYears(value) == enums.Experience3To5Years ||
		enums.ExperienceYears(value) == enums.ExperienceMoreThan5Years
}

func isValidTimeCommitment(value string) bool {
	return enums.TimeCommitment(value) == enums.Time0To2XPerWeek ||
		enums.TimeCommitment(value) == enums.Time3To5XPerWeek ||
		enums.TimeCommitment(value) == enums.TimeMoreThan5XPerWeek
}

func isValidRecipeComplexity(value string) bool {
	return enums.RecipeComplexity(value) == enums.ComplexitySimple ||
		enums.RecipeComplexity(value) == enums.ComplexityModerate ||
		enums.RecipeComplexity(value) == enums.ComplexityComplex
}

func isValidIngredientDiversity(value string) bool {
	return enums.IngredientDiversity(value) == enums.DiversityRare ||
		enums.IngredientDiversity(value) == enums.DiversityOccasional ||
		enums.IngredientDiversity(value) == enums.DiversityFrequent
}
