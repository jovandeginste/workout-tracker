package app

import (
	"net/http"

	"github.com/jovandeginste/workout-tracker/v2/pkg/database"
	"github.com/jovandeginste/workout-tracker/v2/views/equipment"
	"github.com/labstack/echo/v5"
	"github.com/spf13/cast"
)

func (a *App) addRoutesEquipment(g *echo.Group) {
	equipmentGroup := g.Group("/equipment")
	a.GET(equipmentGroup, "", a.equipmentHandler, "equipment")
	a.POST(equipmentGroup, "", a.addEquipment, "equipment-create")
	a.GET(equipmentGroup, "/:id", a.equipmentShowHandler, "equipment-show")
	a.POST(equipmentGroup, "/:id", a.equipmentUpdateHandler, "equipment-update")
	a.GET(equipmentGroup, "/:id/edit", a.equipmentEditHandler, "equipment-edit")
	a.GET(equipmentGroup, "/:id/delete", a.equipmentDeleteConfirmHandler, "equipment-delete-confirm")
	a.POST(equipmentGroup, "/:id/delete", a.equipmentDeleteHandler, "equipment-delete")
	a.GET(equipmentGroup, "/add", a.equipmentAddHandler, "equipment-add")
}

func (a *App) addEquipment(c *echo.Context) error {
	u := a.getCurrentUser(c)
	p := database.Equipment{}

	if err := c.Bind(&p); err != nil {
		return a.redirectWithError(c, a.Reverse("add-equipment"), err)
	}

	p.UserID = u.ID

	if err := p.Save(a.db); err != nil {
		return a.redirectWithError(c, a.Reverse("add-equipment"), err)
	}

	return c.Redirect(http.StatusFound, a.Reverse("equipment"))
}

func (a *App) equipmentHandler(c *echo.Context) error {
	username := c.QueryParam("username")

	target, err := a.userForUsername(c, username)
	if err != nil {
		a.addErrorT(c, "translation.No_user_found_with_username_s", username)

		return Render(c, http.StatusOK, equipment.List(nil, username))
	}

	e, err := target.GetAllEquipment(a.db)
	if err != nil {
		return a.redirectWithError(c, a.Reverse("dashboard"), err)
	}

	return Render(c, http.StatusOK, equipment.List(e, target.Username))
}

func (a *App) equipmentShowHandler(c *echo.Context) error {
	id, err := cast.ToUint64E(c.Param("id"))
	if err != nil {
		return a.redirectWithError(c, a.Reverse("equipment"), err)
	}

	e, err := database.GetEquipment(a.db, id)
	if err != nil {
		return a.redirectWithError(c, a.Reverse("equipment"), err)
	}

	e.FilterVisibleWorkouts(a.getCurrentUser(c))

	return Render(c, http.StatusOK, equipment.Show(e))
}

func (a *App) equipmentAddHandler(c *echo.Context) error {
	return Render(c, http.StatusOK, equipment.Add())
}

func (a *App) equipmentDeleteHandler(c *echo.Context) error {
	e, err := a.getEquipment(c)
	if err != nil {
		return a.redirectWithError(c, a.Reverse("equipment-show", c.Param("id")), err)
	}

	if err := e.Delete(a.db); err != nil {
		return a.redirectWithError(c, a.Reverse("equipment-show", c.Param("id")), err)
	}

	a.addNoticeT(c, "translation.The_equipment_s_has_been_deleted", e.Name)

	if isHtmx(c) {
		c.Response().Header().Set("Hx-Redirect", a.Reverse("equipment"))
		return c.String(http.StatusFound, "ok")
	}

	return c.Redirect(http.StatusFound, a.Reverse("equipment"))
}

func (a *App) equipmentUpdateHandler(c *echo.Context) error {
	e, err := a.getEquipment(c)
	if err != nil {
		return a.redirectWithError(c, a.Reverse("equipment-edit", c.Param("id")), err)
	}

	e.DefaultFor = nil
	e.Active = (c.FormValue("active") == "true")

	if err := c.Bind(e); err != nil {
		return a.redirectWithError(c, a.Reverse("equipment-edit", c.Param("id")), err)
	}

	if err := e.Save(a.db); err != nil {
		return a.redirectWithError(c, a.Reverse("equipment-edit", c.Param("id")), err)
	}

	a.addNoticeT(c, "translation.The_equipment_s_has_been_updated", e.Name)

	return c.Redirect(http.StatusFound, a.Reverse("equipment-show", c.Param("id")))
}

func (a *App) equipmentEditHandler(c *echo.Context) error {
	e, err := a.getEquipment(c)
	if err != nil {
		return a.redirectWithError(c, a.Reverse("equipment"), err)
	}

	return Render(c, http.StatusOK, equipment.Edit(e))
}

func (a *App) equipmentDeleteConfirmHandler(c *echo.Context) error {
	e, err := a.getEquipment(c)
	if err != nil {
		return a.redirectWithError(c, a.Reverse("equipment"), err)
	}

	return Render(c, http.StatusOK, equipment.DeleteModal(e))
}
