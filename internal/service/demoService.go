package service

import (
	"context"
	"github.com/celt237/gin-enhance.demo/model"
)

// DemoService demo service
// @zService
// @zResult model.ResultGeneric
type DemoService struct {
}

func NewDemoService() *DemoService {
	return &DemoService{}
}

// CreateDemo
// @zSummary 创建demo
// @zDescription 创建demo
// @zTags demo
// @zParam        demo   body      model.Demo  true  "demo数据"
// @zRouter       /demo/add [post]
func (d *DemoService) CreateDemo(ctx context.Context, demo *model.Demo) (result int, err error) {
	return 1, err
}

// GetDemoById
// @zSummary 查询demo
// @zDescription 查询demo	path
// @zTags　demo
// @zParam        id   path      int  true  "标识"
// @zRouter       /demo/:id [get]
func (d *DemoService) GetDemoById(ctx context.Context, id int8) (result *model.Demo, err error) {

	return nil, nil
}

// GetDemoById2
// @zSummary 查询demo
// @zDescription 查询demo query
// @zTags　demo
// @zParam        id   query      int  true  "标识"
// @zRouter       /demo/getDemo [get]
func (d *DemoService) GetDemoById2(ctx context.Context, id *int) (result *model.Demo, err error) {
	return &model.Demo{
		Id:   *id,
		Name: "test",
	}, err
}

// Search
// @zSummary 搜索demo
// @zDescription 搜索demo
// @zTags　demo
// @zParam        id   path      int  true  "标识"
// @zParam        name   query      string  true  "名字"
// @zRouter       /demo/search/:id [get]
func (d *DemoService) Search(ctx context.Context, id *int, name string) (result []*model.Demo, err error) {
	list := make([]*model.Demo, 0)
	list = append(list, &model.Demo{
		Id:   *id,
		Name: name,
	})
	return list, err
}

// Search2
// @xPower 1  2  qwe
// @zSummary 搜索demo
// @zDescription 搜索demo
// @zTags　demo
// @zParam pms body model.Demo true "搜索条件"
// @zRouter       /demo/search2 [post]
func (d *DemoService) Search2(ctx context.Context, pms *model.Demo) (result []*model.Demo, err error) {
	list := make([]*model.Demo, 0)
	list = append(list, pms)
	return list, err
}
