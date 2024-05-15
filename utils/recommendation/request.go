package recommendation

import (
	"fmt"
	"github.com/satryanararya/go-chefbot/entities"
)

const RECOMMENDATION_LIMIT = 5

func ToString(recommendations *[]entities.Recommendation, isCron bool) []string {
	var result []string
	var temp string
	for _, recommendation := range *recommendations {
		temp += fmt.Sprintf("%s\n", recommendation.RecipeName)
		if recommendation.ID%RECOMMENDATION_LIMIT == 1 {
			result = append(result, temp[:len(temp)-1])
			if !isCron {
				break
			}
			temp = ""
		}
	}
	return result
}