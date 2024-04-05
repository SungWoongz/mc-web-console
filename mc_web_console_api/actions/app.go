package actions

import (
	"os"
	"strconv"
	"sync"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo-pop/v3/pop/popmw"
	"github.com/gobuffalo/envy"
	contenttype "github.com/gobuffalo/mw-contenttype"
	forcessl "github.com/gobuffalo/mw-forcessl"
	"github.com/gobuffalo/x/sessions"
	"github.com/rs/cors"
	"github.com/unrolled/secure"

	// gwa "github.com/gobuffalo/gocraft-work-adapter"

	i18n "github.com/gobuffalo/mw-i18n/v2"
	paramlogger "github.com/gobuffalo/mw-paramlogger"

	// "github.com/gomodule/redigo/redis"

	"mc_web_console_api/locales"
	"mc_web_console_api/models"

	_ "mc_web_console_api/docs" //mcone의 경우

	buffaloSwagger "github.com/swaggo/buffalo-swagger"
	"github.com/swaggo/buffalo-swagger/swaggerFiles"
)

// ENV is used to help switch settings based on where the
// application is being run. Default is "development".
var ENV = envy.Get("GO_ENV", "development")

var (
	app     *buffalo.App
	appOnce sync.Once
	T       *i18n.Translator
)

func App() *buffalo.App {
	appOnce.Do(func() {
		app = buffalo.New(buffalo.Options{
			Env:          ENV,
			SessionStore: sessions.Null{},
			PreWares: []buffalo.PreWare{
				cors.Default().Handler,
			},
			SessionName: "_mc_web_console_api_session",
			Addr:        os.Getenv("API_ADDR") + ":" + os.Getenv("API_PORT"),
		})

		mciamUse, _ := strconv.ParseBool(os.Getenv("MCIAM_USE"))
		app.ANY("/alive", alive)

		app.Use(forceSSL())
		app.Use(paramlogger.ParameterLogger)
		app.Use(contenttype.Set("application/json"))
		app.Use(popmw.Transaction(models.DB))

		app.GET("/swagger/{*docs}", buffaloSwagger.WrapHandler(swaggerFiles.Handler))

		// app.Use(SkipMiddlewareByRoutePath)
		// app.Use(SetCloudProviderList)

		// RoutesManager(app)

		//// MANUAL ROUTE ////
		apiPath := "/api"

		mcis := app.Group(apiPath + "/mcis")
		mcis.GET("/mcislist", McisList)

		// DEBUG START //

		//  DEBUG END  //

		mciamauth := app.Group(apiPath + "/mciam/auth")
		// MC-IAM-MANAGER REQUIRERD
		if mciamUse {
			mciamauth.Use(McIamAuthMiddleware)
			mciamauth.Middleware.Skip(McIamAuthMiddleware, McIamAuthLoginHandler, McIamAuthGetUserInfoHandler)
		}
		mciamauth.POST("/login", McIamAuthLoginHandler)
		mciamauth.POST("/logout", McIamAuthLogoutHandler)
		mciamauth.GET("/validate", McIamAuthGetUserInfoHandler)

		protectedtest := app.Group(apiPath + "/protected")
		// MC-IAM-MANAGER REQUIRERD
		if mciamUse {
			protectedtest.Use(McIamAuthMiddleware)
		}
		protectedtest.ANY("/alive", alive)
	})

	return app
}

// middleware moved to middleware.go
// translations, forceSSL는 buffalo 에서 정의한 커스텀 미들웨어이므로 이동시키지 않음.

// translations will load locale files, set up the translator `actions.T`,
// and will return a middleware to use to load the correct locale for each
// request.
// for more information: https://gobuffalo.io/en/docs/localization
func translations() buffalo.MiddlewareFunc {
	var err error
	if T, err = i18n.New(locales.FS(), "en-US"); err != nil {
		app.Stop(err)
	}
	return T.Middleware()
}

// forceSSL will return a middleware that will redirect an incoming request
// if it is not HTTPS. "http://example.com" => "https://example.com".
// This middleware does **not** enable SSL. for your application. To do that
// we recommend using a proxy: https://gobuffalo.io/en/docs/proxy
// for more information: https://github.com/unrolled/secure/
func forceSSL() buffalo.MiddlewareFunc {
	return forcessl.Middleware(secure.Options{
		SSLRedirect:     ENV == "production",
		SSLProxyHeaders: map[string]string{"X-Forwarded-Proto": "https"},
	})
}

// Redirect에 route 정보 전달
func RedirectTool(c buffalo.Context, p string) error {
	routes := app.Routes()
	for _, route := range routes {
		if route.PathName == p {
			return c.Redirect(302, route.Path)
		}
	}
	return c.Redirect(302, "/")
}

func alive(c buffalo.Context) error {
	return c.Render(200, r.JSON(map[string]interface{}{
		"status": "OK",
		"method": c.Request().Method,
	}))
}