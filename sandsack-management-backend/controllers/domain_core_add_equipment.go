package controllers

// AddEquipment
// @Description This endpoint adds new equipment
// @Summary This endpoint adds new equipment
// @Accept json
// @Produce json
// @Param Authorization header string true " "
// @Param input body models.OrderEquipment true "name, quantity"
// @Success 200 {array} models.OrderEquipment
// @Failure 500 {object} models.ErrorResponse
// @Failure 400 {object} models.ErrorResponse
// @Tags Core
// @Router /core/equipment/add [post]
/*
func (a *App) AddEquipment(c *gin.Context) {
	var input models.OrderEquipment
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println("AddEquipment error: ", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrCode:    http.StatusBadRequest,
			ErrMessage: "incorrect request",
		})
		return
	}

	err := service.AddEquipment(a.DB, input.Name, input.Quantity)
	if err != nil {
		log.Println("AddEquipment error: ", err.Error())
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			ErrCode:    http.StatusInternalServerError,
			ErrMessage: "something went wrong",
		})
		return
	}

	equipments, err := service.GetEquipment(a.DB)
	c.JSON(http.StatusOK, equipments)
	return
}


 */