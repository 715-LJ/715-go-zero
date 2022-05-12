package common

type (
	UserMens struct {
		Id int32 `json:"id"`
		Name string `json:"name"`
		Type string `json:"type"`
		ParentId int32 `json:"parent_id"`
		Path string `json:"path"`
		Key string `json:"key"`
		Icon string `json:"icon"`
	}
	Data struct {
		Id int32 `json:"id"`
		Name string `json:"name"`
		Type string `json:"type"`
		ParentId int32 `json:"parent_id"`
		Path string `json:"path"`
		Key string `json:"key"`
		Icon string `json:"icon"`
		Children []Data `json:"children"`
	}
	
	Permission struct {
		Id int32 `json:"id"`
		Name string `json:"label"`
		Type int16 `json:"type"`
		ItemType int16 `json:"item_type"`
		ParentId int32 `json:"parent_id"`
		Data map[string]string `json:"data"`
	}
	DataPermission struct {
		Id int32 `json:"id"`
		Name string `json:"label"`
		Type int16 `json:"type"`
		ItemType int16 `json:"item_type"`
		ParentId int32 `json:"parent_id"`
		Data map[string]string `json:"data"`
		Children []DataPermission `json:"children"`
	}
	Permissions struct {
		Id int32 `json:"id"`
		Name string `json:"label"`
		Disabled bool `json:"disabled"`
		Type int16 `json:"type"`
		Order int16 `json:"order"`
		ParentId int32 `json:"parent_id"`
	}
	DataPermissions struct {
		Id int32 `json:"id"`
		Name string `json:"label"`
		Disabled bool `json:"disabled"`
		Type int16 `json:"type"`
		Order int16 `json:"order"`
		ParentId int32 `json:"parent_id"`
		Children []DataPermissions `json:"children"`
	}
)
func MakeTree(list []*UserMens, pid int32)[]Data{
	var tree []Data
	for _, v := range list {
		if v.ParentId == pid {
			var children []Data
			children = append(children, MakeTree(list,v.Id)...)
			tree = append(tree, Data{
				Id: v.Id,
				Name: v.Name,
				Type: v.Type,
				ParentId: v.ParentId,
				Path: v.Path,
				Key: v.Key,
				Children: children,
			})
		}
	}
	return tree
}

func MakeTreePermission(list []*Permission, pid int32)[]DataPermission{
	var tree []DataPermission
	for _, v := range list {
		if v.ParentId == pid {
			var children []DataPermission
			children = append(children, MakeTreePermission(list,v.Id)...)
			tree = append(tree, DataPermission{
				Id: v.Id,
				Name: v.Name,
				Type: v.Type,
				ItemType: v.ItemType,
				ParentId: v.ParentId,
				Data: v.Data,
				Children: children,
			})
		}
	}
	return tree
}

func MakeTreePermissions(list map[int32]Permissions, pid int32)[]DataPermissions{
	var tree []DataPermissions
	for _, v := range list {
		if v.ParentId == pid {
			var children []DataPermissions
			children = append(children, MakeTreePermissions(list,v.Id)...)
			tree = append(tree, DataPermissions{
				Id: v.Id,
				Name: v.Name,
				Disabled: v.Disabled,
				Type: v.Type,
				ParentId: v.ParentId,
				Order: v.Order,
				Children: children,
			})
		}
	}
	return tree
}

/*
* 递归找父级
*/
func GetParent(data,list,arr map[int32]Permissions)map[int32]Permissions{
	for k,i := range data{
		arr[k] = i
		delete(data,k)
		if i.ParentId == 0{
			continue
		}
		arr[i.ParentId] = Permissions{
			Id: list[i.ParentId].Id,
			Name: list[i.ParentId].Name,
			ParentId: list[i.ParentId].ParentId,
			Disabled: list[i.ParentId].Disabled,
			Type: list[i.ParentId].Type,
			Order: list[i.ParentId].Order,
		}
		data[i.ParentId] = arr[i.ParentId]
		val := list[i.ParentId]
		if _,ok := arr[val.ParentId]; !ok{
			return GetParent(data,list,arr)
		}
	}
	return arr
}
/*
* 抛出错误码
*/
func ThrowException(code int32){
	

}