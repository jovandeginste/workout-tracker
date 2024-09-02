package app

import (
	"net/http"
	"strconv"

	"github.com/jovandeginste/workout-tracker/internal/database"
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
	data := a.defaultData(c)

	if err := a.addAllEquipment(a.getCurrentUser(c), data); err != nil {
		return a.redirectWithError(c, a.echo.Reverse("dashboard"), err)
	}

	return c.Render(http.StatusOK, "equipment_list.html", data)
}

func (a *App) equipmentShowHandler(c echo.Context) error {
	data := a.defaultData(c)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return a.redirectWithError(c, "/equipment", err)
	}

	e, err := database.GetEquipment(a.db, id)
	if err != nil {
		return a.redirectWithError(c, "/equipment", err)
	}

	data["equipment"] = e

	return c.Render(http.StatusOK, "equipment_show.html", data)
}

func (a *App) equipmentAddHandler(c echo.Context) error {
	data := a.defaultData(c)
	return c.Render(http.StatusOK, "equipment_add.html", data)
}

func (a *App) equipmentDeleteHandler(c echo.Context) error { //nolint:dupl
	equipment, err := a.getEquipment(c)
	if err != nil {
		return a.redirectWithError(c, a.echo.Reverse("equipment-show", c.Param("id")), err)
	}

	if err := equipment.Delete(a.db); err != nil {
		return a.redirectWithError(c, a.echo.Reverse("equipment-show", c.Param("id")), err)
	}

	a.setNotice(c, "The equipment '%s' has been deleted.", equipment.Name)

	return c.Redirect(http.StatusFound, a.echo.Reverse("equipment"))
}

func (a *App) equipmentUpdateHandler(c echo.Context) error {
	equipment, err := a.getEquipment(c)
	if err != nil {
		return a.redirectWithError(c, a.echo.Reverse("equipment-edit", c.Param("id")), err)
	}

	equipment.DefaultFor = nil
	equipment.Active = (c.FormValue("active") == "true")

	if err := c.Bind(equipment); err != nil {
		return a.redirectWithError(c, a.echo.Reverse("equipment-edit", c.Param("id")), err)
	}

	if err := equipment.Save(a.db); err != nil {
		return a.redirectWithError(c, a.echo.Reverse("equipment-edit", c.Param("id")), err)
	}

	a.setNotice(c, "The equipment '%s' has been updated.", equipment.Name)

	return c.Redirect(http.StatusFound, a.echo.Reverse("equipment-show", c.Param("id")))
}

func (a *App) equipmentEditHandler(c echo.Context) error {
	data := a.defaultData(c)

	equipment, err := a.getEquipment(c)
	if err != nil {
		return a.redirectWithError(c, "/equipment", err)
	}

	data["equipment"] = equipment

	return c.Render(http.StatusOK, "equipment_edit.html", data)
}
