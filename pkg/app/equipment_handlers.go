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
		return a.redirectWithError(c, a.echo.Reverse("add-equipment"), err)
	}

	p.UserID = u.ID

	if err := p.Save(a.db); err != nil {
		return a.redirectWithError(c, a.echo.Reverse("add-equipment"), err)
	}

	return c.Redirect(http.StatusFound, a.echo.Reverse("equipment"))
}

func (a *App) equipmentHandler(c echo.Context) error {
	u := a.getCurrentUser(c)

	e, err := u.GetAllEquipment(a.db)
	if err != nil {
		return a.redirectWithError(c, a.echo.Reverse("dashboard"), err)
	}

	return Render(c, http.StatusOK, equipment.List(e))
}

func (a *App) equipmentShowHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return a.redirectWithError(c, a.echo.Reverse("equipment"), err)
	}

	e, err := database.GetEquipment(a.db, id)
	if err != nil {
		return a.redirectWithError(c, a.echo.Reverse("equipment"), err)
	}

	return Render(c, http.StatusOK, equipment.Show(e))
}

func (a *App) equipmentAddHandler(c echo.Context) error {
	return Render(c, http.StatusOK, equipment.Add())
}

func (a *App) equipmentDeleteHandler(c echo.Context) error {
	e, err := a.getEquipment(c)
	if err != nil {
		return a.redirectWithError(c, a.echo.Reverse("equipment-show", c.Param("id")), err)
	}

	if err := e.Delete(a.db); err != nil {
		return a.redirectWithError(c, a.echo.Reverse("equipment-show", c.Param("id")), err)
	}

	a.addNoticeT(c, "translation.The_equipment_s_has_been_deleted", e.Name)

	return c.Redirect(http.StatusFound, a.echo.Reverse("equipment"))
}

func (a *App) equipmentUpdateHandler(c echo.Context) error {
	e, err := a.getEquipment(c)
	if err != nil {
		return a.redirectWithError(c, a.echo.Reverse("equipment-edit", c.Param("id")), err)
	}

	e.DefaultFor = nil
	e.Active = (c.FormValue("active") == "true")

	if err := c.Bind(e); err != nil {
		return a.redirectWithError(c, a.echo.Reverse("equipment-edit", c.Param("id")), err)
	}

	if err := e.Save(a.db); err != nil {
		return a.redirectWithError(c, a.echo.Reverse("equipment-edit", c.Param("id")), err)
	}

	a.addNoticeT(c, "translation.The_equipment_s_has_been_updated", e.Name)

	return c.Redirect(http.StatusFound, a.echo.Reverse("equipment-show", c.Param("id")))
}

func (a *App) equipmentEditHandler(c echo.Context) error {
	e, err := a.getEquipment(c)
	if err != nil {
		return a.redirectWithError(c, a.echo.Reverse("equipment"), err)
	}

	return Render(c, http.StatusOK, equipment.Edit(e))
}
