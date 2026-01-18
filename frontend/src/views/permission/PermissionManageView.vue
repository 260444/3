<template>
  <div class="permission-manage">
    <el-card class="box-card">
      <template #header>
        <div class="card-header">
          <span>权限管理</span>
        </div>
      </template>

      <el-tabs v-model="activeTab" @tab-change="handleTabChange">
        <!-- 菜单权限管理 -->
        <el-tab-pane label="菜单权限" name="menu">
          <div class="menu-permission">
            <el-form :inline="true" class="demo-form-inline">
              <el-form-item label="选择角色">
                <el-select v-model="selectedRoleId" placeholder="请选择角色" @change="handleRoleChange">
                  <el-option
                    v-for="role in roleList"
                    :key="role.id"
                    :label="role.name"
                    :value="role.id"
                  />
                </el-select>
              </el-form-item>
              <el-form-item>
                <el-button type="primary" @click="handleSaveMenuPermission" :disabled="!selectedRoleId">
                  保存菜单权限
                </el-button>
              </el-form-item>
            </el-form>

            <el-table
              ref="menuTableRef"
              :data="menuList"
              row-key="id"
              border
              default-expand-all
              :tree-props="{ children: 'children' }"
              style="width: 100%"
            >
              <el-table-column prop="title" label="菜单名称" width="200" />
              <el-table-column prop="path" label="路径" />
              <el-table-column prop="component" label="组件" />
              <el-table-column label="权限" width="100">
                <template #default="{ row }">
                  <el-checkbox v-model="row.checked" :disabled="!selectedRoleId">
                    授权
                  </el-checkbox>
                </template>
              </el-table-column>
            </el-table>
          </div>
        </el-tab-pane>

        <!-- API权限管理 -->
        <el-tab-pane label="API权限" name="api">
          <div class="api-permission">
            <el-form :inline="true" class="demo-form-inline">
              <el-form-item label="选择角色">
                <el-select v-model="selectedRoleId" placeholder="请选择角色" @change="handleRoleChange">
                  <el-option
                    v-for="role in roleList"
                    :key="role.id"
                    :label="role.name"
                    :value="role.id"
                  />
                </el-select>
              </el-form-item>
              <el-form-item>
                <el-button type="primary" @click="handleAddPolicy" :disabled="!selectedRoleId">
                  添加策略
                </el-button>
              </el-form-item>
            </el-form>

            <el-table :data="policyList" border style="width: 100%">
              <el-table-column prop="path" label="路径" />
              <el-table-column prop="method" label="方法" width="100">
                <template #default="{ row }">
                  <el-tag :type="getMethodType(row.method)">{{ row.method }}</el-tag>
                </template>
              </el-table-column>
              <el-table-column label="操作" width="100">
                <template #default="{ row }">
                  <el-button type="danger" size="small" @click="handleRemovePolicy(row)">
                    删除
                  </el-button>
                </template>
              </el-table-column>
            </el-table>
          </div>
        </el-tab-pane>

        <!-- 所有权限策略 -->
        <el-tab-pane label="所有策略" name="all">
          <el-table :data="allPolicyList" border style="width: 100%">
            <el-table-column prop="0" label="主体" width="150" />
            <el-table-column prop="1" label="对象" />
            <el-table-column prop="2" label="动作" width="100">
              <template #default="{ row }">
                <el-tag :type="getMethodType(row[2])">{{ row[2] }}</el-tag>
              </template>
            </el-table-column>
          </el-table>
        </el-tab-pane>
      </el-tabs>
    </el-card>

    <!-- 添加策略对话框 -->
    <el-dialog v-model="policyDialogVisible" title="添加API权限策略" width="500px">
      <el-form :model="policyForm" label-width="80px">
        <el-form-item label="路径">
          <el-input v-model="policyForm.path" placeholder="例如: /api/v1/users" />
        </el-form-item>
        <el-form-item label="方法">
          <el-select v-model="policyForm.method" placeholder="请选择方法">
            <el-option label="GET" value="GET" />
            <el-option label="POST" value="POST" />
            <el-option label="PUT" value="PUT" />
            <el-option label="DELETE" value="DELETE" />
          </el-select>
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="policyDialogVisible = false">取消</el-button>
          <el-button type="primary" @click="handleConfirmAddPolicy">确定</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { ElMessage } from 'element-plus'
import * as roleApi from '@/api/role'
import * as menuApi from '@/api/menu'
import * as permissionApi from '@/api/permission'

const route = useRoute()

// 角色列表
const roleList = ref<any[]>([])
// 菜单列表
const menuList = ref<any[]>([])
// 策略列表
const policyList = ref<any[]>([])
// 所有策略列表
const allPolicyList = ref<any[]>([])
// 选中的角色ID
const selectedRoleId = ref<number | null>(null)
// 当前激活的标签页
const activeTab = ref('menu')
// 策略对话框可见性
const policyDialogVisible = ref(false)
// 策略表单
const policyForm = ref({
  path: '',
  method: 'GET'
})

// 获取角色列表
const fetchRoleList = async () => {
  try {
    const response: any = await roleApi.getRoles()
    roleList.value = response.data.list
  } catch (error) {
    ElMessage.error('获取角色列表失败')
  }
}

// 获取菜单列表
const fetchMenuList = async () => {
  try {
    const response: any = await menuApi.getMenuTree()
    menuList.value = response.data
  } catch (error) {
    ElMessage.error('获取菜单列表失败')
  }
}

// 获取角色的菜单权限
const fetchRoleMenus = async (roleId: number) => {
  try {
    const response: any = await permissionApi.getRoleMenus(roleId)
    const roleMenus = response.data

    // 重置所有菜单的选中状态
    const resetMenuChecked = (menus: any[]) => {
      menus.forEach(menu => {
        menu.checked = false
        if (menu.children && menu.children.length > 0) {
          resetMenuChecked(menu.children)
        }
      })
    }
    resetMenuChecked(menuList.value)

    // 设置角色拥有的菜单为选中状态
    const setMenuChecked = (menus: any[]) => {
      menus.forEach(menu => {
        if (roleMenus.find((m: any) => m.id === menu.id)) {
          menu.checked = true
        }
        if (menu.children && menu.children.length > 0) {
          setMenuChecked(menu.children)
        }
      })
    }
    setMenuChecked(menuList.value)
  } catch (error) {
    ElMessage.error('获取角色菜单权限失败')
  }
}

// 获取角色的策略列表
const fetchRolePolicies = async (roleId: number) => {
  try {
    const response: any = await permissionApi.getPolicies(roleId)
    policyList.value = response.data.map((policy: string[]) => ({
      path: policy[1],
      method: policy[2]
    }))
  } catch (error) {
    ElMessage.error('获取角色策略失败')
  }
}

// 获取所有策略
const fetchAllPolicies = async () => {
  try {
    const response: any = await permissionApi.getAllPolicies()
    allPolicyList.value = response.data
  } catch (error) {
    ElMessage.error('获取所有策略失败')
  }
}

// 角色改变
const handleRoleChange = (roleId: number) => {
  if (activeTab.value === 'menu') {
    fetchRoleMenus(roleId)
  } else if (activeTab.value === 'api') {
    fetchRolePolicies(roleId)
  }
}

// 标签页改变
const handleTabChange = (tabName: string) => {
  if (tabName === 'all') {
    fetchAllPolicies()
  } else if (selectedRoleId.value) {
    handleRoleChange(selectedRoleId.value)
  }
}

// 保存菜单权限
const handleSaveMenuPermission = async () => {
  if (!selectedRoleId.value) {
    ElMessage.warning('请先选择角色')
    return
  }

  const getCheckedMenuIds = (menus: any[]): number[] => {
    let ids: number[] = []
    menus.forEach(menu => {
      if (menu.checked) {
        ids.push(menu.id)
      }
      if (menu.children && menu.children.length > 0) {
        ids = ids.concat(getCheckedMenuIds(menu.children))
      }
    })
    return ids
  }

  const menuIds = getCheckedMenuIds(menuList.value)

  try {
    await permissionApi.assignMenuToRole(selectedRoleId.value, menuIds)
    ElMessage.success('菜单权限保存成功')
  } catch (error) {
    ElMessage.error('保存菜单权限失败')
  }
}

// 添加策略
const handleAddPolicy = () => {
  policyForm.value = {
    path: '',
    method: 'GET'
  }
  policyDialogVisible.value = true
}

// 确认添加策略
const handleConfirmAddPolicy = async () => {
  if (!policyForm.value.path || !policyForm.value.method) {
    ElMessage.warning('请填写完整的策略信息')
    return
  }

  try {
    await permissionApi.addPolicy(selectedRoleId.value!, policyForm.value.path, policyForm.value.method)
    ElMessage.success('策略添加成功')
    policyDialogVisible.value = false
    fetchRolePolicies(selectedRoleId.value!)
  } catch (error) {
    ElMessage.error('添加策略失败')
  }
}

// 移除策略
const handleRemovePolicy = async (policy: any) => {
  try {
    await permissionApi.removePolicy(selectedRoleId.value!, policy.path, policy.method)
    ElMessage.success('策略删除成功')
    fetchRolePolicies(selectedRoleId.value!)
  } catch (error) {
    ElMessage.error('删除策略失败')
  }
}

// 获取方法类型
const getMethodType = (method: string) => {
  const typeMap: any = {
    GET: 'success',
    POST: 'primary',
    PUT: 'warning',
    DELETE: 'danger'
  }
  return typeMap[method] || 'info'
}

onMounted(() => {
  fetchRoleList()
  fetchMenuList()
  fetchAllPolicies()

  // 检查是否有路由参数（从角色管理页面跳转过来）
  const roleId = route.query.roleId
  if (roleId) {
    selectedRoleId.value = Number(roleId)
    activeTab.value = 'menu'
    fetchRoleMenus(Number(roleId))
  }
})
</script>

<style scoped>
.permission-manage {
  padding: 20px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.menu-permission,
.api-permission {
  padding: 20px 0;
}

.demo-form-inline {
  margin-bottom: 20px;
}
</style>