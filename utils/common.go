package utils

import "gin_mall/v1/model/response"

//生成菜单树结构

func MenuTree(list []*response.Menu, RuleId int64) []*response.Menu {
	var treeList []*response.Menu
	for _, r := range list {
		if r.RuleId == RuleId {
			r.Child = MenuTree(list, r.Id)
			treeList = append(treeList, r)
		}
	}
	return treeList
}

//生成分类树结构

func CategoryTree(list []*response.CategoryList, categoryId int64) []*response.CategoryList {
	var treeList []*response.CategoryList
	for _, r := range list {
		if r.CategoryId == categoryId {
			r.Child = CategoryTree(list, r.Id)
			treeList = append(treeList, r)
		}
	}
	return treeList
}
