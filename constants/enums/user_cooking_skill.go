package enums

// ExperienceYears adalah enumerasi untuk tahun pengalaman memasak
type ExperienceYears string

const (
    Experience0To2Years ExperienceYears = "0-2 tahun"
    Experience3To5Years ExperienceYears = "3-5 tahun"
    ExperienceMoreThan5Years ExperienceYears = ">5 tahun"
)

// TimeCommitment adalah enumerasi untuk komitmen waktu memasak dalam seminggu
type TimeCommitment string

const (
    Time0To2XPerWeek TimeCommitment = "1-2x per minggu"
    Time3To5XPerWeek TimeCommitment = "3-5x per minggu"
    TimeMoreThan5XPerWeek TimeCommitment = ">5x per minggu"
)

// RecipeComplexity adalah enumerasi untuk kompleksitas resep yang dicoba
type RecipeComplexity string

const (
    ComplexitySimple RecipeComplexity = "sederhana"
    ComplexityModerate RecipeComplexity = "sedang"
    ComplexityComplex RecipeComplexity = "rumit"
)

// IngredientDiversity adalah enumerasi untuk keragaman bahan yang dicoba
type IngredientDiversity string

const (
    DiversityRare IngredientDiversity = "jarang"
    DiversityOccasional IngredientDiversity = "kadang"
    DiversityFrequent IngredientDiversity = "sering"
)
