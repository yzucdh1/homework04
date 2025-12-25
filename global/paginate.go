package global

import (
	"fmt"
)

type PageReq struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

type PageRes struct {
	List       interface{} `json:"list"`        // 当前页数据列表
	Total      int64       `json:"total"`       // 总条数
	Page       int         `json:"page"`        // 当前页
	PageSize   int         `json:"page_size"`   // 每页条数
	TotalPages int64       `json:"total_pages"` // 总页数
}

// GetOffset 计算Offset值
func (p *PageReq) GetOffset() int {
	if p.Page <= 0 {
		p.Page = 1 // 默认第1页
	}
	if p.PageSize <= 0 || p.PageSize > 100 {
		p.PageSize = 10 // 默认10条，最大100条
	}
	return (p.Page - 1) * p.PageSize
}

func (p *PageReq) GetLimit() int {
	return p.PageSize
}

// CalcTotalPages 计算总页数
func (p *PageReq) CalcTotalPages(total int64) int64 {
	if total == 0 {
		return 0
	}
	return (total + int64(p.PageSize) - 1) / int64(p.PageSize)
}

func Paginate(model interface{}, pageQuery *PageReq, query any, args any) (PageRes, error) {
	// 1.处理分页参数
	offset := pageQuery.GetOffset()
	limit := pageQuery.GetLimit()

	// 2.查询总条数（仅查询count，不查数据）
	var total int64
	var err error
	if query != nil && args != nil {
		err = DB.Model(model).Where(query, args).Count(&total).Error
	} else {
		err = DB.Model(model).Count(&total).Error
	}
	if err != nil {
		return PageRes{}, fmt.Errorf("查询总条数失败：%w", err)
	}

	// 3.分页查询数据
	if query != nil && args != nil {
		err = DB.Offset(offset).Limit(limit).Where(query, args).Find(model).Error
	} else {
		err = DB.Offset(offset).Limit(limit).Find(model).Error
	}
	if err != nil {
		return PageRes{}, fmt.Errorf("分页查询数据失败：%w", err)
	}

	// 4.组装分页结果
	return PageRes{
		List:       model,
		Total:      total,
		Page:       pageQuery.Page,
		PageSize:   pageQuery.PageSize,
		TotalPages: pageQuery.CalcTotalPages(total),
	}, nil
}
