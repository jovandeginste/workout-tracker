package app

import (
	"net/http"
	"strconv"

	"github.com/jovandeginste/workout-tracker/v2/pkg/database"
	"github.com/jovandeginste/workout-tracker/v2/views/equipment"
	"github.com/labstack/echo/v4"
)

func (a *App) addRoutesEquipment(e *echo.Group) {
	equipmentGroup := e.Group("/equipment")
	equipmentGroup.GET("", a.equipmentHandler).Name = "equipment"
	equipmentGroup.POST("", a.addEquipment).Name = "equipment-create"
	equipmentGroup.GET("/:id", a.equipmentShowHandler).Name = "equipment-show"
	equipmentGroup.POST("/:id", a.equipmentUpdateHandler).Name = "equipment-update"
	equipmentGroup.GET("/:id/edit", a.equipmentEditHandler).Name = "equipment-edit"
	equipmentGroup.POST("/:id/delete", a.equipmentDeleteHandler).Name = "equipment-delete"
	equipmentGroup.GET("/add", a.equipmentAddHandler).Name = "equipment-add"
}

func (a *App) addEquipment(c echo.Context) error {
	u := a.getCurrentUser(c)
	p := database.Equipment{}

	if err := c.Bind(&p); err != nil {
		return a.renderError(c, http.StatusBadRequest, err, "add-equipment")
	}

	p.UserID = u.ID

	if err := p.Save(a.db); err != nil {
		return a.renderError(c, http.StatusInternalServerError, err, "add-equipment")
	}

	return c.Redirect(http.StatusFound, a.echo.Reverse("equipment"))
}

func (a *App) equipmentHandler(c echo.Context) error {
	a.setContext(c)

	u := a.getCurrentUser(c)

	e, err := u.GetAllEquipment(a.db)
	if err != nil {
		return a.renderError(c, http.StatusInternalServerError, err, "dashboard")
	}

	return Render(c, http.StatusOK, equipment.List(e))
}

func (a *App) equipmentShowHandler(c echo.Context) error {
	a.setContext(c)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return a.renderError(c, http.StatusBadRequest, err, "equipment")
	}

	e, err := database.GetEquipment(a.db, id)
	if err != nil {
		return a.renderError(c, http.StatusNotFound, err, "equipment")
	}

	return Render(c, http.StatusOK, equipment.Show(e))
}

func (a *App) equipmentAddHandler(c echo.Context) error {
	a.setContext(c)
	return Render(c, http.StatusOK, equipment.Add())
}

func (a *App) equipmentDeleteHandler(c echo.Context) error {
	e, err := a.getEquipment(c)
	if err != nil {
		return a.renderError(c, http.StatusNotFound, err, "equipment-show", c.Param("id"))
	}

	if err := e.Delete(a.db); err != nil {
		return a.renderError(c, http.StatusInternalServerError, err, "equipment-show", c.Param("id"))
	}

	a.addNotice(c, "The equipment '%s' has been deleted.", e.Name)

	return c.Redirect(http.StatusFound, a.echo.Reverse("equipment"))
}

func (a *App) equipmentUpdateHandler(c echo.Context) error {
	e, err := a.getEquipment(c)
	if err != nil {
		return a.renderError(c, http.StatusNotFound, err, "equipment-edit", c.Param("id"))
	}

	e.DefaultFor = nil
	e.Active = (c.FormValue("active") == "true")

	if err := c.Bind(e); err != nil {
		return a.renderError(c, http.StatusBadRequest, err, "equipment-edit", c.Param("id"))
	}

	if err := e.Save(a.db); err != nil {
		return a.renderError(c, http.StatusInternalServerError, err, "equipment-edit", c.Param("id"))
	}

	a.addNotice(c, "The equipment '%s' has been updated.", e.Name)

	return c.Redirect(http.StatusFound, a.echo.Reverse("equipment-show", c.Param("id")))
}

func (a *App) equipmentEditHandler(c echo.Context) error {
	a.setContext(c)

	e, err := a.getEquipment(c)
	if err != nil {
		return a.renderError(c, http.StatusNotFound, err, "equipment")
	}

	return Render(c, http.StatusOK, equipment.Edit(e))
}

func (a *App) renderError(c echo.Context, statusCode int, err error, route string, params ...any) error {
	a.addError(c, err.Error())
	if len(params) > 0 {
		return c.Redirect(statusCode, a.echo.Reverse(route, params...))
	}
	return c.Redirect(statusCode, a.echo.Reverse(route))
}
