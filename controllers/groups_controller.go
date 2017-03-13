package controllers

import (
	"net/http"

	"fmt"

	"github.com/aadeschamps/giftexchangeapi/models"
)

// GroupController exports all methods needed to act on users
type GroupController struct {
	Group *models.GroupModel
}

// Show retrieves a specific group by id
func (c *GroupController) Show(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello, groups!")
}
