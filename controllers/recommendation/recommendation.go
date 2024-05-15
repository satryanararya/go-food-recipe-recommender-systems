package recommendation

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	msg "github.com/satryanararya/go-chefbot/constants/message"
	"github.com/satryanararya/go-chefbot/usecases"
	http_util "github.com/satryanararya/go-chefbot/utils/http"
	"github.com/satryanararya/go-chefbot/utils/token"
)

type recommendationController struct {
	recommendationUseCase usecases.RecommendationUseCase
	token                 token.TokenUtil
}

func NewRecommendationController(ruc usecases.RecommendationUseCase, t token.TokenUtil) *recommendationController {
	return &recommendationController{
		recommendationUseCase: ruc,
		token:                 t,
	}
}

func (rc *recommendationController) GetRecommendation(c echo.Context) error {
	claims := rc.token.GetClaims(c)

	recommendation, err := rc.recommendationUseCase.GetRecommendation(c, claims.ID)
	if err != nil {
		fmt.Println("Error: ", err)
		return http_util.HandleErrorResponse(c, http.StatusInternalServerError, msg.MsgGetRecommendationFailed)
	}
	return http_util.HandleSuccessResponse(c, http.StatusOK, msg.MsgGetRecommendationSuccess, recommendation)
}
