package system_manager

import (
	sysModel "backend/internal/model/system_manager"
	"fmt"

	"gorm.io/gorm"
)

// RoleMenuRepository 规则与菜单关系仓库
type RoleMenuRepository struct {
	DB *gorm.DB
}

// NewRoleMenuRepository 规则与菜单关系仓库
func NewRoleMenuRepository(db *gorm.DB) *RoleMenuRepository {
	return &RoleMenuRepository{DB: db}
}

// CreateRoleMenus 创建记录RoleMenu
func (r *RoleMenuRepository) CreateRoleMenus(roleMeans []sysModel.RoleMenu) error {
	return r.DB.Create(&roleMeans).Error
}

// GetRoleMenus 获取记录RoleMenu
//func (r *RoleMenuRepository) GetRoleMenus(roleId uint) (roleMeans []*sysModel.RoleMenuRequest, err error) {
//	err = r.DB.Raw(`
//		SELECT a.role_id, a.menu_id, b.title
//		FROM role_menus AS a
//		LEFT JOIN menus AS b ON a.menu_id = b.id
//		WHERE a.role_id = ?
//    `).Scan(&roleMeans).Error
//	return roleMeans, err
//}

// DeleteRoleMenus 删除记录RoleMenu
func (r *RoleMenuRepository) DeleteRoleMenus(roleId uint, roleMeans []uint) error {
	return r.DB.Where("role_id = ? AND menu_id IN ?", roleId, roleMeans).Delete(&sysModel.RoleMenu{}).Error
}

// GetRoleMenuByID 根据ID获取记录RoleMenu
func (r *RoleMenuRepository) GetRoleMenuByID(roleId uint) (roleMeans []*sysModel.RoleMenuRequest, err error) {

	var p []uint
	//查询出有权限的父菜单
	//select  menu_id from role_menus where role_id=2 and menu_id in (select id from menus where parent_id = 0);
	err = r.DB.Raw(`SELECT menu_id FROM role_menus WHERE role_id = ? AND menu_id IN (SELECT id FROM menus WHERE parent_id = 0)`, roleId).Scan(&p).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	// 初始化切片，设置正确的容量
	roleMeans = make([]*sysModel.RoleMenuRequest, len(p))
	// 为每个元素分配内存
	for i := range roleMeans {
		roleMeans[i] = &sysModel.RoleMenuRequest{}
	}

	//递归查询出所有子菜单
	for i, pid := range p {
		roleMeans[i].PId = pid
		err = r.DB.Raw(`select  menu_id from role_menus where  role_id= ? and menu_id in (select id from menus where parent_id = ?)`, roleId, pid).Scan(&roleMeans[i].MId).Error
		if err != nil {
			return nil, err
		}
		fmt.Println("pid", roleMeans[i])
		fmt.Println("mids", roleMeans[i].MId)
	}

	return roleMeans, err
}
