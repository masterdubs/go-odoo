package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type PurchaseOrderService struct {
	client *Client
}

func NewPurchaseOrderService(c *Client) *PurchaseOrderService {
	return &PurchaseOrderService{client: c}
}

func (svc *PurchaseOrderService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.PurchaseOrderModel, name)
}

func (svc *PurchaseOrderService) GetByIds(ids []int) (*types.PurchaseOrders, error) {
	p := &types.PurchaseOrders{}
	return p, svc.client.getByIds(types.PurchaseOrderModel, ids, p)
}

func (svc *PurchaseOrderService) GetByName(name string) (*types.PurchaseOrders, error) {
	p := &types.PurchaseOrders{}
	return p, svc.client.getByName(types.PurchaseOrderModel, name, p)
}

func (svc *PurchaseOrderService) GetAll() (*types.PurchaseOrders, error) {
	p := &types.PurchaseOrders{}
	return p, svc.client.getAll(types.PurchaseOrderModel, p)
}

func (svc *PurchaseOrderService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.PurchaseOrderModel, fields, relations)
}

func (svc *PurchaseOrderService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.PurchaseOrderModel, ids, fields, relations)
}

func (svc *PurchaseOrderService) Delete(ids []int) error {
	return svc.client.delete(types.PurchaseOrderModel, ids)
}
