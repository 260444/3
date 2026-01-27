<template>
  <div class="role-manage-container">
    <el-card>
      <!-- 操作区域 -->
      <div class="header-section">
        <el-button type="primary" @click="handleAdd">新增角色</el-button>
      </div>

      <!-- 角色列表 -->
      <el-table :data="roleList" stripe style="width: 100%" v-loading="loading">
        <el-table-column prop="id" label="ID" width="80" />
        <el-table-column prop="ident" label="角色标识" />
        <el-table-column prop="name" label="角色名称" />
        <el-table-column prop="description" label="描述" />
        <el-table-column label="状态" width="100">
          <template #default="{ row }">
            <el-tag :type="row.status === 1 ? 'success' : 'danger'">
              {{ row.status === 1 ? '正常' : '禁用' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="created_at" label="创建时间" width="180" />
        <el-table-column label="操作" width="420">
          <template #default="{ row }">
            <el-button size="small" @click="handleEdit(row)">编辑</el-button>
            <el-button size="small" type="primary" @click="handleAssignMenus(row)">分配菜单</el-button>
            <el-button size="small" type="warning" @click="handleAssignPermissions(row)">分配权限</el-button>
            <el-button size="small" type="danger" @click="handleDelete(row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>

      <!-- 分页 -->
      <div class="pagination">
        <el-pagination
          v-model:current-page="pagination.page"
          v-model:page-size="pagination.page_size"
          :page-sizes="[10, 20, 50, 100]"
          :total="pagination.total"
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange"
        />
      </div>
    </el-card>

    <!-- 角色编辑弹窗 -->
    <el-dialog v-model="dialogVisible" :title="dialogTitle" width="500px">
      <el-form :model="roleForm" :rules="roleRules" ref="roleFormRef" label-width="80px">
        <el-form-item label="角色名称" prop="name">
          <el-input v-model="roleForm.name" placeholder="请输入角色名称" />
        </el-form-item>
        <el-form-item label="角色标识" prop="ident">
          <el-input v-model="roleForm.ident" placeholder="请输入角色标识" />
        </el-form-item>
        <el-form-item label="描述" prop="description">
          <el-input v-model="roleForm.description" type="textarea" placeholder="请输入角色描述" />
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-radio-group v-model="roleForm.status">
            <el-radio :label="1">正常</el-radio>
            <el-radio :label="0">禁用</el-radio>
          </el-radio-group>
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="dialogVisible = false">取消</el-button>
          <el-button type="primary" @click="confirmRole">确定</el-button>
        </span>
      </template>
    </el-dialog>

    <!-- 分配菜单弹窗 -->
    <el-dialog v-model="menuDialogVisible" :title="`为【${currentRoleName}】分配菜单`" width="600px">
      <el-tree
        :data="allMenus"
        :props="{ children: 'children', label: 'title', value: 'id' }"
        node-key="id"
        show-checkbox
        ref="menuTreeRef"
        :loading="menuTreeLoading"
        @check="onMenuCheckChange"
        :check-strictly="true"
      >
        <template #default="{ node, data }">
          <span class="custom-tree-node">
            <span>{{ data.title }}</span>
            <span v-if="data.name" class="menu-name">({{ data.name }})</span>
          </span>
        </template>
      </el-tree>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="menuDialogVisible = false">取消</el-button>
          <el-button type="primary" @click="submitMenuAssignment">确定分配</el-button>
        </span>
      </template>
    </el-dialog>

    <!-- 分配权限弹窗 -->
    <el-dialog v-model="permissionDialogVisible" :title="`为【${currentRoleName}】分配权限`" width="800px">
      <div class="permission-filter">
        <el-input
          v-model="permissionSearch"
          placeholder="搜索权限路径或描述"
          clearable
          style="width: 300px; margin-bottom: 20px;"
        />
        <el-button type="primary" @click="loadPermissions">刷新</el-button>
      </div>
      
      <el-table
        :data="filteredPermissions"
        v-loading="permissionLoading"
        ref="permissionTableRef"
        @selection-change="handlePermissionSelectionChange"
        max-height="400"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column prop="id" label="ID" width="80" />
        <el-table-column prop="path" label="路径" min-width="150" />
        <el-table-column prop="method" label="方法" width="100">
          <template #default="{ row }">
            <el-tag :type="getMethodType(row.method)">{{ row.method }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="description" label="描述" min-width="150" />
        <el-table-column label="状态" width="100">
          <template #default="{ row }">
            <el-tag :type="row.status === 1 ? 'success' : 'danger'">
              {{ row.status === 1 ? '正常' : '禁用' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="created_at" label="创建时间" width="180" />
      </el-table>
      
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="permissionDialogVisible = false">取消</el-button>
          <el-button type="primary" @click="submitPermissionAssignment" :disabled="selectedPermissions.length === 0">确定分配</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted, nextTick, computed } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage, ElMessageBox } from 'element-plus'
import { getRoles, createRole, updateRole, deleteRole, getRoleMenus, assignRoleMenus, removeRoleMenus } from '@/api/role'
import { getAllMenus } from '@/api/menu'
import { getPermissions, getPolicies, addPolicy, removePolicy } from '@/api/permission'

// 从API响应中提取菜单ID的辅助函数
// 根据API文档，响应格式为: [{"p_id":1,"m_id":null},{"p_id":2,"m_id":[11,12,14]}]
// p_id代表父菜单ID，m_id代表该父菜单下的子菜单ID数组
// 在业务逻辑中，当p_id有分配的子菜单(m_id数组)时，p_id代表的父菜单也应该被选中
const extractMenuIdsFromResponse = (responseData: any[]): number[] => {
  let menuIds: number[] = []
  
  if (Array.isArray(responseData)) {
    for (const item of responseData) {
      // 将父菜单ID添加到选中列表（无论m_id是数组、单个ID还是null）
      menuIds.push(item.p_id);
      
      if (item.m_id && Array.isArray(item.m_id)) {
        // 如果m_id是数组，则添加所有子菜单ID
        menuIds = menuIds.concat(item.m_id)
      } else if (item.m_id !== null && item.m_id !== undefined) {
        // 如果m_id是单个ID，则添加
        menuIds.push(item.m_id)
      }
      // 注意：即使item.m_id为null，我们已经添加了item.p_id（父菜单）
    }
  }
  
  return menuIds
}

const router = useRouter()

// 分页信息
const pagination = reactive({
  page: 1,
  page_size: 10,
  total: 0
})

// 角色列表
const roleList = ref<any[]>([])
const loading = ref(false)

// 角色编辑弹窗相关
const dialogVisible = ref(false)
const dialogTitle = ref('')
const roleFormRef = ref()
const roleForm = ref<any>({
  id: undefined,
  name: '',
  description: '',
  status: 1
})

// 分配菜单弹窗相关
const menuDialogVisible = ref(false)
const currentRoleId = ref(0)
const currentRoleName = ref('')
const allMenus = ref<any[]>([])
const checkedMenuIds = ref<number[]>([])
const menuTreeLoading = ref(false)
const menuTreeRef = ref()

// 分配权限弹窗相关
const permissionDialogVisible = ref(false)
const permissionLoading = ref(false)
const permissions = ref<any[]>([])
const selectedPermissions = ref<any[]>([])
const permissionSearch = ref('')
const permissionTableRef = ref()

// 获取方法类型用于标签显示
const getMethodType = (method: string) => {
  switch (method.toUpperCase()) {
    case 'GET':
      return 'success'
    case 'POST':
      return 'primary'
    case 'PUT':
      return 'warning'
    case 'DELETE':
      return 'danger'
    default:
      return 'info'
  }
}

// 表单验证规则
const roleRules = {
  name: [
    { required: true, message: '请输入角色名称', trigger: 'blur' },
    { min: 2, max: 20, message: '长度在 2 到 20 个字符', trigger: 'blur' }
  ],
  ident: [
    { required: true, message: '请输入角色标识', trigger: 'blur' },
    { min: 2, max: 50, message: '长度在 2 到 50 个字符', trigger: 'blur' }
  ],
  description: [
    { required: true, message: '请输入角色描述', trigger: 'blur' }
  ]
}

// 获取角色列表
const getRoleList = async () => {
  loading.value = true
  try {
    const params = {
      page: pagination.page,
      page_size: pagination.page_size
    }
    const response = await getRoles(params)
    roleList.value = response.data.list
    pagination.total = response.data.total
  } catch (error) {
    console.error('获取角色列表失败:', error)
    ElMessage.error('获取角色列表失败')
  } finally {
    loading.value = false
  }
}

// 处理分页大小变化
const handleSizeChange = (size: number) => {
  pagination.page_size = size
  pagination.page = 1
  getRoleList()
}

// 处理当前页变化
const handleCurrentChange = (page: number) => {
  pagination.page = page
  getRoleList()
}

// 处理添加角色
const handleAdd = () => {
  roleForm.value = {
    id: undefined,
    name: '',
    description: '',
    status: 1
  }
  dialogTitle.value = '新增角色'
  dialogVisible.value = true
}

// 处理编辑角色
const handleEdit = (row: any) => {
  roleForm.value = { ...row }
  dialogTitle.value = '编辑角色'
  dialogVisible.value = true
}

// 确认角色操作（新增或编辑）
const confirmRole = async () => {
  try {
    await roleFormRef.value.validate()
    
    if (roleForm.value.id) {
      // 编辑角色
      await updateRole(roleForm.value.id, roleForm.value)
      ElMessage.success('角色更新成功')
    } else {
      // 新增角色
      await createRole(roleForm.value)
      ElMessage.success('角色创建成功')
    }
    
    dialogVisible.value = false
    getRoleList()
  } catch (error) {
    console.error('操作角色失败:', error)
    ElMessage.error('操作失败')
  }
}

// 处理删除角色
const handleDelete = async (row: any) => {
  try {
    await ElMessageBox.confirm(`确定要删除角色 "${row.name}" 吗？`, '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    })

    await deleteRole(row.id)
    ElMessage.success('删除成功')
    getRoleList()
  } catch (error) {
    if (error !== 'cancel') {
      console.error('删除角色失败:', error)
      ElMessage.error('删除失败')
    }
  }
}

// 处理分配菜单
const handleAssignMenus = async (row: any) => {
  currentRoleId.value = row.id
  currentRoleName.value = row.name
  menuDialogVisible.value = true
  await loadMenuData()
}

// 加载菜单数据
const loadMenuData = async () => {
  menuTreeLoading.value = true
  try {
    // 获取所有菜单
    const allMenusResponse = await getAllMenus()
    allMenus.value = allMenusResponse.data

    // 获取当前角色已分配的菜单
    const roleMenusResponse = await getRoleMenus(currentRoleId.value)
    
    // 使用辅助函数从API响应中提取菜单ID
    const menuIds = extractMenuIdsFromResponse(roleMenusResponse.data)
    checkedMenuIds.value = menuIds

    // 等待DOM更新后设置默认选中状态
    await nextTick()
    if (menuTreeRef.value) {
      // 使用setCheckedKeys设置选中状态，Tree组件会自动处理父子关系
      menuTreeRef.value.setCheckedKeys(checkedMenuIds.value)
    }
  } catch (error) {
    console.error('加载菜单数据失败:', error)
    ElMessage.error('加载菜单数据失败')
  } finally {
    menuTreeLoading.value = false
  }
}

// 提交菜单分配
const submitMenuAssignment = async () => {
  try {
    // 获取当前完全选中的菜单ID（包括因子菜单选中而自动选中的父菜单）
    const currentCheckedIds = menuTreeRef.value.getCheckedKeys()
    
    // 获取当前半选中的菜单ID（父菜单，其部分子菜单被选中）
    // const currentHalfCheckedIds = menuTreeRef.value.getHalfCheckedKeys() || []
    
    // 获取当前角色已有的菜单权限
    const roleMenusResponse = await getRoleMenus(currentRoleId.value)
    
    // 使用辅助函数从API响应中提取菜单ID
    const existingMenuIds = extractMenuIdsFromResponse(roleMenusResponse.data)

    // 计算需要新增的菜单ID（用户新选择的）
    const menusToAdd = currentCheckedIds.filter((id: number) => !existingMenuIds.includes(id))

    // 计算需要移除的菜单ID（用户取消选择的）
    const menusToRemove = existingMenuIds.filter((id: number) => !currentCheckedIds.includes(id))

    // 执行添加操作
    if (menusToAdd.length > 0) {
      await assignRoleMenus(currentRoleId.value, menusToAdd)
    }

    // 执行移除操作
    if (menusToRemove.length > 0) {
      await removeRoleMenus(currentRoleId.value, menusToRemove)
    }

    ElMessage.success('菜单权限分配成功')
    menuDialogVisible.value = false
  } catch (error) {
    console.error('分配菜单权限失败:', error)
    ElMessage.error('分配菜单权限失败')
  }
}

// 菜单选中状态变化处理
const onMenuCheckChange = () => {
  checkedMenuIds.value = menuTreeRef.value.getCheckedKeys()
}

// 处理分配权限
const handleAssignPermissions = async (row: any) => {
  currentRoleId.value = row.id
  currentRoleName.value = row.name
  permissionDialogVisible.value = true
  await loadPermissions()
}

// 加载权限列表
const loadPermissions = async () => {
  permissionLoading.value = true
  try {
    const params = {
      page: 1,
      page_size: 1000 // 获取所有权限
    }
    const response = await getPermissions(params)
    permissions.value = response.data.list || []
    
    // 获取当前角色已有的权限策略
    const policiesResponse = await getPolicies(currentRoleId.value)
    const existingPolicies = policiesResponse.data || []
    
    // 设置已选中的权限
    await nextTick()
    if (permissionTableRef.value) {
      const selectedIds = existingPolicies.map((policy: any) => policy.obj + policy.act)
      // 根据已有策略设置选中状态
      permissions.value.forEach(permission => {
        const permissionKey = permission.path + permission.method
        if (selectedIds.includes(permissionKey)) {
          permissionTableRef.value.toggleRowSelection(permission, true)
        }
      })
    }
  } catch (error) {
    console.error('加载权限列表失败:', error)
    ElMessage.error('加载权限列表失败')
  } finally {
    permissionLoading.value = false
  }
}

// 过滤后的权限列表
const filteredPermissions = computed(() => {
  if (!permissionSearch.value) {
    return permissions.value
  }
  return permissions.value.filter(permission => 
    permission.path.toLowerCase().includes(permissionSearch.value.toLowerCase()) ||
    (permission.description && permission.description.toLowerCase().includes(permissionSearch.value.toLowerCase()))
  )
})

// 处理权限表格选中状态变化
const handlePermissionSelectionChange = (selection: any[]) => {
  selectedPermissions.value = selection
}

// 提交权限分配
const submitPermissionAssignment = async () => {
  try {
    // 获取当前角色已有的权限策略
    const policiesResponse = await getPolicies(currentRoleId.value)
    const existingPolicies = policiesResponse.data || []
    
    // 当前选中的权限
    const currentSelected = permissionTableRef.value.getSelectionRows()
    
    // 计算需要添加的权限
    const permissionsToAdd = currentSelected.filter((current: any) => 
      !existingPolicies.some((existing: any) => 
        existing.obj === current.path && existing.act === current.method
      )
    )
    
    // 计算需要移除的权限
    const permissionsToRemove = existingPolicies.filter((existing: any) => 
      !currentSelected.some((current: any) => 
        existing.obj === current.path && existing.act === current.method
      )
    )
    
    // 执行添加操作
    for (const permission of permissionsToAdd) {
      await addPolicy(currentRoleId.value, permission.path, permission.method)
    }
    
    // 执行移除操作
    for (const policy of permissionsToRemove) {
      await removePolicy(currentRoleId.value, policy.obj, policy.act)
    }
    
    ElMessage.success('权限分配成功')
    permissionDialogVisible.value = false
  } catch (error) {
    console.error('分配权限失败:', error)
    ElMessage.error('分配权限失败')
  }
}

onMounted(() => {
  getRoleList()
})
</script>

<style scoped>
.role-manage-container {
  background-color: #fff;
  padding: 20px;
  border-radius: 4px;
}

.header-section {
  margin-bottom: 20px;
}

.pagination {
  margin-top: 20px;
  display: flex;
  justify-content: flex-end;
}

.custom-tree-node {
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: space-between;
  font-size: 14px;
  padding-right: 8px;
}

.menu-name {
  color: #999;
  font-size: 12px;
  margin-left: 8px;
}

.permission-filter {
  margin-bottom: 20px;
  display: flex;
  gap: 10px;
  align-items: center;
}
</style>