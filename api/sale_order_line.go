package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type SaleOrderLineService struct {
	client *Client
}

func NewSaleOrderLineService(c *Client) *SaleOrderLineService {
	return &SaleOrderLineService{client: c}
}

func (svc *SaleOrderLineService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.SaleOrderLineModel, name)
}

func (svc *SaleOrderLineService) GetByIds(ids []int) (*types.SaleOrderLines, error) {
	s := &types.SaleOrderLines{}
	return s, svc.client.getByIds(types.SaleOrderLineModel, ids, s)
}

func (svc *SaleOrderLineService) GetByName(name string) (*types.SaleOrderLines, error) {
	s := &types.SaleOrderLines{}
	return s, svc.client.getByName(types.SaleOrderLineModel, name, s)
}

func (svc *SaleOrderLineService) GetAll() (*types.SaleOrderLines, error) {
	s := &types.SaleOrderLines{}
	return s, svc.client.getAll(types.SaleOrderLineModel, s)
}

func (svc *SaleOrderLineService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.SaleOrderLineModel, fields, relations)
}

func (svc *SaleOrderLineService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.SaleOrderLineModel, ids, fields, relations)
}

func (svc *SaleOrderLineService) Delete(ids []int) error {
	return svc.client.delete(types.SaleOrderLineModel, ids)
}
