package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type MailAliasService struct {
	client *Client
}

func NewMailAliasService(c *Client) *MailAliasService {
	return &MailAliasService{client: c}
}

func (svc *MailAliasService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.MailAliasModel, name)
}

func (svc *MailAliasService) GetByIds(ids []int) (*types.MailAliass, error) {
	m := &types.MailAliass{}
	return m, svc.client.getByIds(types.MailAliasModel, ids, m)
}

func (svc *MailAliasService) GetByName(name string) (*types.MailAliass, error) {
	m := &types.MailAliass{}
	return m, svc.client.getByName(types.MailAliasModel, name, m)
}

func (svc *MailAliasService) GetAll() (*types.MailAliass, error) {
	m := &types.MailAliass{}
	return m, svc.client.getAll(types.MailAliasModel, m)
}

func (svc *MailAliasService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.MailAliasModel, fields, relations)
}

func (svc *MailAliasService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.MailAliasModel, ids, fields, relations)
}

func (svc *MailAliasService) Delete(ids []int) error {
	return svc.client.delete(types.MailAliasModel, ids)
}
