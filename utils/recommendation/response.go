package recommendation

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/satryanararya/go-chefbot/entities"
)

func ToStruct(result string, userID uuid.UUID) *[]entities.Recommendation {
	var recommendations []entities.Recommendation
	splittedResult := strings.Split(result, "\n")

	fmt.Println("splittedResult", splittedResult)

	for _, line := range splittedResult {
		trim := strings.TrimLeft(line, "-12345. ")
		if trim != "" {
			recommendations = append(recommendations, entities.Recommendation{
				UserID:     userID,
				RecipeName: trim,
			})
		}
	}
	return &recommendations
}
