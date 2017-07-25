package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type IrExportsLineService struct {
	client *Client
}

func NewIrExportsLineService(c *Client) *IrExportsLineService {
	return &IrExportsLineService{client: c}
}

func (svc *IrExportsLineService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.IrExportsLineModel, name)
}

func (svc *IrExportsLineService) GetByIds(ids []int) (*types.IrExportsLines, error) {
	i := &types.IrExportsLines{}
	return i, svc.client.getByIds(types.IrExportsLineModel, ids, i)
}

func (svc *IrExportsLineService) GetByName(name string) (*types.IrExportsLines, error) {
	i := &types.IrExportsLines{}
	return i, svc.client.getByName(types.IrExportsLineModel, name, i)
}

func (svc *IrExportsLineService) GetAll() (*types.IrExportsLines, error) {
	i := &types.IrExportsLines{}
	return i, svc.client.getAll(types.IrExportsLineModel, i)
}

func (svc *IrExportsLineService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.IrExportsLineModel, fields, relations)
}

func (svc *IrExportsLineService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.IrExportsLineModel, ids, fields, relations)
}

func (svc *IrExportsLineService) Delete(ids []int) error {
	return svc.client.delete(types.IrExportsLineModel, ids)
}
