package service

import (
	"errors"
	"gorm.io/gorm"
	"order/db"
	"order/model"
	"strconv"
)

func OrderList(userId uint, name string, gameId uint, orderId int64, status string, pageSize int, pageNo int) ([]model.Order, int64, error) {
	var order []model.Order
	pagination := db.DB.Where("user_id", userId).Preload("Mode").Preload("Game").Preload("Log")
	if status != "" {
		pagination = pagination.Where("status = ?", status)
	}
	if name != "" {
		pagination = pagination.Where("name like ?", "%"+name+"%")
	}
	if gameId != 0 {
		pagination = pagination.Where("game_id = ?", gameId)
	}
	if orderId != 0 {
		pagination = pagination.Where("order_id = ?", orderId)
	}
	if err := pagination.
		Offset(pageSize * (pageNo - 1)).Order("id desc").Limit(pageSize).Find(&order).Error; err != nil {
		return nil, 0, err
	}
	var rowCount int64
	if err := pagination.Count(&rowCount).Error; err != nil {
		return nil, 0, err
	}
	return order, rowCount, nil
}

// OrderSubmit status 0 初始化 1 执行中，2 执行完毕， -1 执行失败，-2 暂停订单
func OrderSubmit(data *model.Order) error {
	if err := db.DB.Debug().Model(model.Order{}).Create(data).Error; err != nil {
		return errors.New("创建订单失败！")
	}
	return nil
}

func GetNewOrder(user model.User) (model.Order, error) {
	order := model.Order{}
	if err := db.DB.Model(model.Order{}).
		Preload("Mode").
		Preload("Game").
		Where("user_id  = ? and status = 0", user.ID).
		Where("NOT EXISTS (SELECT 1 FROM orders where `status` = 1 and user_id = ?)", user.ID).First(&order).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return order, errors.New("当前有执行的订单，或者没有新的订单，无法执行新订单！")
		}
	}
	if err := db.DB.Model(model.Order{}).Where("id = ?", order.ID).Update("status", 1).RowsAffected; err == 0 {
		return order, errors.New("订单状态更新异常！")
	}
	return order, nil
}

func GetMachineNewOrder(machine model.UserMachine) (order model.Order, err error) {
	if err = db.DB.Model(model.Order{}).
		Preload("Mode").
		Preload("Game").
		Where("user_id  = ? and status = 0 and machine = ?", machine.UserId, machine.Machine).
		Where("NOT EXISTS (SELECT 1 FROM orders where `status` = 1 and user_id = ? and machine = ?)", machine.UserId, machine.Machine).First(&order).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return order, errors.New("当前有执行的订单，或者没有新的订单，无法执行新订单！")
		}
	}
	if rows := db.DB.Model(model.Order{}).Where("id = ?", order.ID).Update("status", 1).RowsAffected; rows == 0 {
		return order, errors.New("订单状态更新异常！")
	}
	return order, nil
}

func GetIngOrder(user model.User) (model.Order, error) {
	order := model.Order{}
	if err := db.DB.Model(model.Order{}).
		Preload("Mode").
		Preload("Game").
		Where("`status` = 1 and user_id = ?", user.ID).First(&order).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return order, errors.New("当前没有执行的订单，请获取新订单！")
		}
	}
	return order, nil
}

func GetMachineIngOrder(machine model.UserMachine) (order model.Order, err error) {
	if err = db.DB.Model(model.Order{}).
		Preload("Mode").
		Preload("Game").
		Where("`status` = 1 and user_id = ? and machine = ?", machine.UserId, machine.Machine).First(&order).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return order, errors.New("当前没有执行的订单，请获取新订单！")
		}
	}
	return
}

func GetOrderByOrderID(orderId int64) (model.Order, error) {
	order := model.Order{}
	if err := db.DB.Model(model.Order{}).Where("order_id = ?", orderId).First(&order).Error; err != nil {
		return order, err
	}
	return order, nil
}

func UpdateOrder(orderId int64, status string, cnum string) error {
	order := map[string]interface{}{}
	if status != "" {
		order["status"], _ = strconv.Atoi(status)
	}
	if cnum != "" {
		order["cnum"], _ = strconv.Atoi(cnum)
	}
	database := db.DB.Model(model.Order{}).Where("status = 1 and order_id  = ?", orderId)
	if err := database.First(map[string]interface{}{}).Error; err != nil {
		return errors.New("订单不存在，或者不是执行中的订单！")
	}
	if err := database.Updates(order).RowsAffected; err == 0 {
		return errors.New("订单不存在，或者不是执行中的订单！")
	}
	return nil
}

func UpdateOrderByStatus(orderId int64, status int) (model.Order, error) {
	order := model.Order{}
	database := db.DB.Model(model.Order{})
	if status == -2 {
		database = database.Where(" order_id  = ?", orderId)
	} else {
		database.Where("status != 1 and order_id  = ?", orderId)
	}

	if err := database.First(&order).Error; err != nil {
		return order, errors.New("订单不存在，或者订单正在执行中！")
	}
	if err := database.Update("status", status).RowsAffected; err == 0 {
		return order, errors.New("订单更新失败！")
	}
	return order, nil
}
