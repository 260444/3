<template>
  <div class="operation-log-container">
    <el-card>
      <!-- 搜索和操作区域 -->
      <div class="header-section">
        <el-form :inline="true" :model="searchForm">
          <el-form-item label="操作用户">
            <el-input v-model="searchForm.username" placeholder="请输入操作用户" />
          </el-form-item>
          <el-form-item label="操作内容">
            <el-input v-model="searchForm.operation" placeholder="请输入操作内容" />
          </el-form-item>
          <el-form-item label="操作时间">
            <el-date-picker
              v-model="dateRange"
              type="datetimerange"
              range-separator="至"
              start-placeholder="开始时间"
              end-placeholder="结束时间"
              format="YYYY-MM-DD HH:mm:ss"
              value-format="YYYY-MM-DD HH:mm:ss"
            />
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="handleSearch">搜索</el-button>
            <el-button @click="resetSearch">重置</el-button>
          </el-form-item>
        </el-form>
      </div>

      <!-- 操作日志列表 -->
      <el-table :data="logList" stripe style="width: 100%" v-loading="loading">
        <el-table-column prop="id" label="ID" width="80" />
        <el-table-column prop="username" label="操作用户" width="120" />
        <el-table-column prop="operation" label="操作内容" show-overflow-tooltip />
        <el-table-column prop="request_method" label="请求方法" width="100" />
        <el-table-column prop="request_path" label="请求路径" width="150" />
        <el-table-column prop="ip" label="IP地址" width="130" />
        <el-table-column label="状态" width="80">
          <template #default="{ row }">
            <el-tag :type="row.status === 1 ? 'success' : 'danger'">
              {{ row.status === 1 ? '成功' : '失败' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="response_time" label="响应时间(ms)" width="120" />
        <el-table-column prop="created_at" label="操作时间" width="180" />
        <el-table-column label="操作" width="100">
          <template #default="{ row }">
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
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { getOperationLogs, deleteOperationLog } from '@/api/operationLog'

// 搜索表单
const searchForm = reactive({
  username: '',
  operation: ''
})

// 日期范围
const dateRange = ref<[string, string] | null>(null)

// 分页信息
const pagination = reactive({
  page: 1,
  page_size: 10,
  total: 0
})

// 日志列表
const logList = ref<any[]>([])
const loading = ref(false)

// 获取操作日志列表
const getOperationLogList = async () => {
  loading.value = true
  try {
    const params: any = {
      page: pagination.page,
      page_size: pagination.page_size,
      ...searchForm
    }
    
    // 处理日期范围
    if (dateRange.value) {
      params.start_time = dateRange.value[0]
      params.end_time = dateRange.value[1]
    }
    
    const response = await getOperationLogs(params)
    logList.value = response.data.list
    pagination.total = response.data.total
  } catch (error) {
    console.error('获取操作日志列表失败:', error)
    ElMessage.error('获取操作日志列表失败')
  } finally {
    loading.value = false
  }
}

// 搜索
const handleSearch = () => {
  pagination.page = 1
  getOperationLogList()
}

// 重置搜索
const resetSearch = () => {
  searchForm.username = ''
  searchForm.operation = ''
  dateRange.value = null
  pagination.page = 1
  getOperationLogList()
}

// 处理分页大小变化
const handleSizeChange = (size: number) => {
  pagination.page_size = size
  pagination.page = 1
  getOperationLogList()
}

// 处理当前页变化
const handleCurrentChange = (page: number) => {
  pagination.page = page
  getOperationLogList()
}

// 处理删除操作日志
const handleDelete = async (row: any) => {
  try {
    await ElMessageBox.confirm(`确定要删除该操作日志吗？`, '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    })
    
    await deleteOperationLog(row.id)
    ElMessage.success('删除成功')
    getOperationLogList()
  } catch (error) {
    if (error !== 'cancel') {
      console.error('删除操作日志失败:', error)
      ElMessage.error('删除失败')
    }
  }
}

onMounted(() => {
  getOperationLogList()
})
</script>

<style scoped>
.operation-log-container {
  background-color: #fff;
  padding: 20px;
  border-radius: 4px;
}

.header-section {
  margin-bottom: 20px;
}

.header-section .el-form {
  display: flex;
  flex-wrap: wrap;
}

.header-section .el-form-item {
  margin-right: 10px;
  margin-bottom: 10px;
}

.pagination {
  margin-top: 20px;
  display: flex;
  justify-content: flex-end;
}
</style>