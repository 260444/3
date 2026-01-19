<template>
  <div class="menu-manage-container">
    <el-card>
      <!-- 操作区域 -->
      <div class="header-section">
        <el-button type="primary" @click="handleAdd()">新增根菜单</el-button>
      </div>

      <!-- 菜单树 -->
      <el-table 
        :data="menuTree" 
        stripe 
        style="width: 100%" 
        row-key="id"
        default-expand-all
        v-loading="loading"
      >
        <el-table-column prop="title" label="菜单名称" width="200">
          <template #default="{ row }">
            <span :style="{ marginLeft: (row.level || 0) * 20 + 'px' }">
              {{ row.title }}
            </span>
          </template>
        </el-table-column>
        <el-table-column prop="name" label="标识" width="120" />
        <el-table-column prop="path" label="路径" width="150" />
        <el-table-column prop="component" label="组件" width="150" />
        <el-table-column prop="icon" label="图标" width="80">
          <template #default="{ row }">
            <el-icon v-if="row.icon">
              <component :is="row.icon" />
            </el-icon>
          </template>
        </el-table-column>
        <el-table-column label="状态" width="80">
          <template #default="{ row }">
            <el-tag :type="row.status === 1 ? 'success' : 'danger'">
              {{ row.status === 1 ? '显示' : '隐藏' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="排序" prop="sort" width="80" />
        <el-table-column label="操作" width="200">
          <template #default="{ row }">
            <el-button size="small" @click="handleAdd(row)">新增</el-button>
            <el-button size="small" @click="handleEdit(row)">编辑</el-button>
            <el-button size="small" type="danger" @click="handleDelete(row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>

    </el-card>

    <!-- 菜单编辑弹窗 -->
    <el-dialog v-model="dialogVisible" :title="dialogTitle" width="600px">
      <el-form :model="menuForm" :rules="menuRules" ref="menuFormRef" label-width="100px">
        <el-form-item label="父级菜单">
          <el-cascader
            v-model="menuForm.parent_id"
            :options="menuOptions"
            :props="{ value: 'id', label: 'title', checkStrictly: true, emitPath: false }"
            placeholder="请选择父级菜单（可选）"
            clearable
            style="width: 100%"
          />
        </el-form-item>
        <el-form-item label="菜单名称" prop="title">
          <el-input v-model="menuForm.title" placeholder="请输入菜单名称" />
        </el-form-item>
        <el-form-item label="标识" prop="name">
          <el-input v-model="menuForm.name" placeholder="请输入菜单标识" />
        </el-form-item>
        <el-form-item label="路径" prop="path">
          <el-input v-model="menuForm.path" placeholder="请输入菜单路径" />
        </el-form-item>
        <el-form-item label="组件" prop="component">
          <el-input v-model="menuForm.component" placeholder="请输入组件路径" />
        </el-form-item>
        <el-form-item label="重定向">
          <el-input v-model="menuForm.redirect" placeholder="请输入重定向路径" />
        </el-form-item>
        <el-form-item label="图标">
          <el-select v-model="menuForm.icon" placeholder="请选择图标" style="width: 100%">
            <el-option
              v-for="icon in iconList"
              :key="icon"
              :label="icon"
              :value="icon"
            >
              <div style="display: flex; align-items: center;">
                <el-icon>
                  <component :is="icon" />
                </el-icon>
                <span style="margin-left: 10px;">{{ icon }}</span>
              </div>
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="排序" prop="sort">
          <el-input-number v-model="menuForm.sort" :min="0" style="width: 100%" />
        </el-form-item>
        <el-form-item label="是否隐藏">
          <el-switch v-model="menuForm.is_hidden" />
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-radio-group v-model="menuForm.status">
            <el-radio :label="1">正常</el-radio>
            <el-radio :label="0">禁用</el-radio>
          </el-radio-group>
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="dialogVisible = false">取消</el-button>
          <el-button type="primary" @click="confirmMenu">确定</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { 
  getMenuTree as fetchMenuTree, 
  createMenu, 
  updateMenu, 
  deleteMenu,
  getAllMenus 
} from '@/api/menu'
import { 
  House, 
  User, 
  Avatar, 
  Menu as MenuIcon, 
  Document, 
  Setting, 
  Grid, 
  Link, 
  SwitchButton, 
  Postcard, 
  Tickets, 
  FolderOpened 
} from '@element-plus/icons-vue'

// 菜单列表
const menuTree = ref<any[]>([])
const loading = ref(false)

// 弹窗相关
const dialogVisible = ref(false)
const dialogTitle = ref('')
const menuFormRef = ref()
const menuForm = ref<any>({
  id: undefined,
  parent_id: null,
  name: '',
  title: '',
  path: '',
  component: '',
  redirect: '',
  icon: '',
  sort: 0,
  is_hidden: false,
  status: 1
})

// 图标列表
const iconList = [
  'House', 'User', 'Avatar', 'Menu', 'Document', 
  'Setting', 'Grid', 'Link', 'SwitchButton', 'Postcard',
  'Tickets', 'FolderOpened'
]

// 菜单选项（用于级联选择器）
const menuOptions = ref<any[]>([])

// 表单验证规则
const menuRules = {
  title: [
    { required: true, message: '请输入菜单名称', trigger: 'blur' }
  ],
  name: [
    { required: true, message: '请输入菜单标识', trigger: 'blur' }
  ],
  path: [
    { required: true, message: '请输入菜单路径', trigger: 'blur' }
  ]
}

// 获取菜单树
const getMenuTree = async () => {
  loading.value = true
  try {
    const response = await getAllMenus() // 使用 /api/v1/menus/all 接口获取所有菜单
    // 后端已返回完整的树形结构，直接使用即可
    menuTree.value = addLevelInfo(response.data, 0)
    buildMenuOptions()
  } catch (error) {
    console.error('获取菜单树失败:', error)
    ElMessage.error('获取菜单树失败')
  } finally {
    loading.value = false
  }
}



// 为菜单树添加层级信息（用于显示缩进）
const addLevelInfo = (menus: any[], level: number): any[] => {
  return menus.map(menu => ({
    ...menu,
    level,
    children: menu.children ? addLevelInfo(menu.children, level + 1) : []
  }))
}

// 构建菜单选项（用于级联选择器）
const buildMenuOptions = () => {
  const buildOptions = (menus: any[]): any[] => {
    return menus.map(menu => ({
      id: menu.id,
      title: menu.title,
      children: menu.children ? buildOptions(menu.children) : []
    }))
  }
  
  // 使用当前菜单树数据构建选项
  menuOptions.value = buildOptions(menuTree.value)
}

// 处理添加菜单
const handleAdd = (parentMenu?: any) => {
  menuForm.value = {
    id: undefined,
    parent_id: parentMenu ? parentMenu.id : null,
    name: '',
    title: '',
    path: '',
    component: '',
    redirect: '',
    icon: '',
    sort: 0,
    is_hidden: false,
    status: 1
  }
  dialogTitle.value = parentMenu ? `新增子菜单 - ${parentMenu.title}` : '新增根菜单'
  dialogVisible.value = true
}

// 处理编辑菜单
const handleEdit = (row: any) => {
  menuForm.value = { ...row }
  // 如果没有父菜单，设置为null而不是undefined
  if (!menuForm.value.parent_id) {
    menuForm.value.parent_id = null
  }
  dialogTitle.value = '编辑菜单'
  dialogVisible.value = true
}

// 确认菜单操作（新增或编辑）
const confirmMenu = async () => {
  try {
    await menuFormRef.value.validate()
    
    // 处理parent_id
    const menuData = { ...menuForm.value }
    if (!menuData.parent_id) {
      menuData.parent_id = null
    }
    
    if (menuData.id) {
      // 编辑菜单
      await updateMenu(menuData.id, menuData)
      ElMessage.success('菜单更新成功')
    } else {
      // 新增菜单
      await createMenu(menuData)
      ElMessage.success('菜单创建成功')
    }
    
    dialogVisible.value = false
    getMenuTree()
  } catch (error) {
    console.error('操作菜单失败:', error)
    ElMessage.error('操作失败')
  }
}

// 处理删除菜单
const handleDelete = async (row: any) => {
  try {
    await ElMessageBox.confirm(`确定要删除菜单 "${row.title}" 吗？`, '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    })
    
    await deleteMenu(row.id)
    ElMessage.success('删除成功')
    getMenuTree()
  } catch (error) {
    if (error !== 'cancel') {
      console.error('删除菜单失败:', error)
      ElMessage.error('删除失败')
    }
  }
}

onMounted(() => {
  getMenuTree()
})
</script>

<style scoped>
.menu-manage-container {
  background-color: #fff;
  padding: 20px;
  border-radius: 4px;
}

.header-section {
  margin-bottom: 20px;
}

:deep(.el-table .el-table__row) {
  height: 50px;
}

:deep(.el-table .el-table__cell) {
  padding: 8px 0;
}
</style>