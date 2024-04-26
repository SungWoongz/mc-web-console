package actions

import (
	"log"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/mitchellh/mapstructure"

	mcmodels "mc_web_console_common_models"
)

func UserLoginHandler(c buffalo.Context) error {
	if c.Request().Method == "POST" {

		// user bind ------------------------------ 1
		user := &mcmodels.UserLogin{}
		if err := c.Bind(user); err != nil {
			return c.Render(http.StatusBadRequest,
				defaultRender.JSON(map[string]string{"err": err.Error()}))
		}

		// validateErr ------------------------------ 2
		validateErr := validate.Validate(
			&validators.StringIsPresent{Field: user.Id, Name: "id"},
			&validators.StringIsPresent{Field: user.Password, Name: "password"},
		)
		if validateErr.HasAny() {
			log.Println(validateErr)
			return c.Render(http.StatusBadRequest,
				defaultRender.JSON(map[string]string{"err": validateErr.Error()}))
		}

		// commonRequest RequestData 할당 ------------------------------ 3
		commonRequest := &mcmodels.CommonRequest{}
		commonRequest.RequestData = user

		// API 호출 ------------------------------ 4
		status, commonRes, err := CommonAPIPostWithoutAccessToken(APILoginPath, commonRequest)
		if err != nil {
			return c.Render(status.StatusCode,
				defaultRender.JSON(map[string]string{"err": err.Error()}))
		}
		if status.StatusCode != 200 {
			return c.Render(status.StatusCode,
				defaultRender.JSON(map[string]string{"err": status.Status}),
			)
		}

		// commonRes to return 할당 ------------------------------ 5
		accessTokenResponse := &mcmodels.AccessTokenResponse{}
		decodeerr := mapstructure.Decode(commonRes.ResponseData, accessTokenResponse)
		if decodeerr != nil {
			return c.Render(status.StatusCode,
				defaultRender.JSON(map[string]string{"err": decodeerr.Error()}))
		}

		c.Session().Set("Authorization", accessTokenResponse.AccessToken)

		return c.Render(http.StatusOK,
			defaultRender.JSON(map[string]string{
				"redirect": RootPathForRedirectString,
			}))
	}

	accessToken := c.Session().Get("Authorization")
	if accessToken != nil {
		c.Flash().Add("success", "You are already logged in.")
		return c.Redirect(http.StatusSeeOther,
			"webconsoleDepth1Depth2Depth3Path()",
			RootPathForRedirect,
		)
	}

	return c.Render(http.StatusOK, webconsoleRender.HTML("auth/login.html"))
}

func UserLogoutHandler(c buffalo.Context) error {
	status, _, err := CommonAPIPost(APILogoutPath, &mcmodels.CommonRequest{}, c)
	if err != nil {
		log.Println(err.Error())
		return c.Render(status.StatusCode,
			defaultRender.JSON(map[string]string{"err": err.Error()}))
	}
	if status.StatusCode != 200 {
		return c.Render(status.StatusCode,
			defaultRender.JSON(map[string]string{"err": status.Status}),
		)
	}

	c.Session().Clear()
	c.Flash().Add("success", "Logout")
	return c.Redirect(http.StatusSeeOther, "authLoginPath()")
}

func UserRegisterpageHandler(c buffalo.Context) error {
	return c.Render(http.StatusOK, webconsoleRender.HTML("auth/login.html"))
}

// func GetUserRefreshTokenHandler(c buffalo.Context) (mcmodels.AccessTokenResponse, string, error) {
// 	status, respBody, err := CommonAPIPost(APILoginRefreshPath, accessTokenRequest, c)
// }
