package enums

// ExperienceYears is an enumeration for years of cooking experience
type ExperienceYears string

const (
    Experience0To2Years ExperienceYears = "0-2 years"
    Experience3To5Years ExperienceYears = "3-5 years"
    ExperienceMoreThan5Years ExperienceYears = ">5 years"
)

// TimeCommitment is an enumeration of cooking time commitments for the week
type TimeCommitment string

const (
    Time0To2XPerWeek TimeCommitment = "1-2x per week"
    Time3To5XPerWeek TimeCommitment = "3-5x per week"
    TimeMoreThan5XPerWeek TimeCommitment = ">5x per week"
)

// RecipeComplexity is an enumeration for the complexity of the recipe attempted
type RecipeComplexity string

const (
    ComplexitySimple RecipeComplexity = "simple"
    ComplexityModerate RecipeComplexity = "moderate"
    ComplexityComplex RecipeComplexity = "complex"
)

// IngredientDiversity is an enumeration for the diversity of ingredients tried
type IngredientDiversity string

const (
    DiversityRare IngredientDiversity = "rare"
    DiversityOccasional IngredientDiversity = "occasional"
    DiversityFrequent IngredientDiversity = "frequent"
)
