package apps

import (
	"encoding/json"
	"net/http"

	"github.com/aerogear/mobile-security-service/pkg/helpers"
	"github.com/aerogear/mobile-security-service/pkg/httperrors"
	"github.com/aerogear/mobile-security-service/pkg/models"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

type (
	HTTPHandler interface {
		GetApps(c echo.Context) error
		GetActiveAppByID(c echo.Context) error
		UpdateAppVersions(c echo.Context) error
		DisableAllAppVersionsByAppID(c echo.Context) error
		DeleteAppById(c echo.Context) error
		CreateApp(c echo.Context) error
		UpdateAppNameByID(c echo.Context) error
	}

	// httpHandler instance
	httpHandler struct {
		Service Service
	}
)

// NewHTTPHandler returns a new instance of app.Handler
func NewHTTPHandler(e *echo.Echo, s Service) HTTPHandler {
	return &httpHandler{
		Service: s,
	}
}

// GetApps returns all apps as JSON from the AppService
func (a *httpHandler) GetApps(c echo.Context) error {
	apps, err := a.HandleGetApp(c)

	// If no apps have been found, return a HTTP Status code of 204 with no response body
	if err == models.ErrNotFound {
		return c.NoContent(http.StatusNoContent)
	}

	if err != nil {
		return httperrors.GetHTTPResponseFromErr(c, err)
	}

	return c.JSON(http.StatusOK, apps)
}

// HandleGetApp will return handle the request according to the data provided
func (a *httpHandler) HandleGetApp(c echo.Context) (*[]models.App, error) {
	appId := c.QueryParam("appId")
	var apps *[]models.App
	var err error
	if len(appId) > 1 {
		var app *models.App
		app, err = a.Service.GetActiveAppByAppID(appId)
		if app != nil {
			apps = &[]models.App{*app}
		}
	} else {
		apps, err = a.Service.GetApps()
	}
	return apps, err
}

// GetActiveAppByID returns apps by id as JSON from the AppService
func (a *httpHandler) GetActiveAppByID(c echo.Context) error {

	id := c.Param("id")
	if !helpers.IsValidUUID(id) {
		return httperrors.BadRequest(c, "Invalid id supplied")
	}

	apps, err := a.Service.GetActiveAppByID(id)

	if err != nil {
		return httperrors.GetHTTPResponseFromErr(c, err)
	}
	return c.JSON(http.StatusOK, apps)

}

func (a *httpHandler) UpdateAppNameByID(c echo.Context) error {

	id := c.Param("id")
	if !helpers.IsValidUUID(id) {
		return httperrors.BadRequest(c, "Invalid id supplied")
	}

	// Transform the body request in the version struct
	app := models.App{}
	errV := json.NewDecoder(c.Request().Body).Decode(&app)

	// check if the data sent is in the correct format
	if errV != nil || len(app.AppID) < 1 {
		return httperrors.BadRequest(c, "Invalid data")
	}

	err := a.Service.UpdateAppNameByID(id, app.AppName)
	if err != nil {
		return httperrors.GetHTTPResponseFromErr(c, err)
	}

	return c.NoContent(http.StatusNoContent)
}

//UpdateApp returns a app updated with the ID in JSON format from the AppService
func (a *httpHandler) UpdateAppVersions(c echo.Context) error {
	// Validations
	id := c.Param("id")
	if !helpers.IsValidUUID(id) {
		return httperrors.BadRequest(c, "Invalid id supplied")
	}

	versions := []models.Version{}
	errV := json.NewDecoder(c.Request().Body).Decode(&versions)

	// check if the data sent is in the correct format
	if errV != nil {
		return httperrors.BadRequest(c, "Invalid data")
	}

	// Check if versions were sent all the body is empty
	if len(versions) == 0 {
		return httperrors.BadRequest(c, "No version(s) was sent.")
	}

	// Call service
	errUpdate := a.Service.UpdateAppVersions(id, versions)
	if errUpdate != nil {
		return httperrors.GetHTTPResponseFromErr(c, errUpdate)
	}

	return c.NoContent(http.StatusNoContent)
}

//UpdateApp returns a app updated with the ID in JSON format from the AppService
func (a *httpHandler) DisableAllAppVersionsByAppID(c echo.Context) error {
	id := c.Param("id")
	if !helpers.IsValidUUID(id) {
		return httperrors.BadRequest(c, "Invalid id supplied")
	}

	// Transform the body request in the version struct
	version := models.Version{}

	if err := json.NewDecoder(c.Request().Body).Decode(&version); err != nil {
		log.Error(err)
		return httperrors.BadRequest(c, "Invalid data")
	}

	err := a.Service.DisableAllAppVersionsByAppID(id, version.DisabledMessage)

	if err != nil {
		return httperrors.GetHTTPResponseFromErr(c, err)
	}

	return c.NoContent(http.StatusNoContent)

}

func (a *httpHandler) CreateApp(c echo.Context) error {

	// Transform the body request in the version struct
	app := models.App{}
	errV := json.NewDecoder(c.Request().Body).Decode(&app)

	// check if the data sent is in the correct format
	if errV != nil || len(app.AppID) < 1 {
		return httperrors.BadRequest(c, "Invalid data")
	}

	err := a.Service.CreateApp(app)

	if err != nil {
		return httperrors.GetHTTPResponseFromErr(c, err)
	}

	return c.NoContent(http.StatusCreated)

}

func (a *httpHandler) DeleteAppById(c echo.Context) error {
	id := c.Param("id")
	if !helpers.IsValidUUID(id) {
		return httperrors.BadRequest(c, "Invalid id supplied")
	}

	err := a.Service.DeleteAppById(id)

	if err != nil {
		return httperrors.GetHTTPResponseFromErr(c, err)
	}

	return c.NoContent(http.StatusNoContent)
}
