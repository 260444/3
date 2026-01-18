<template>
  <div class="dashboard-container">
    <el-row :gutter="20">
      <el-col :span="6">
        <el-card class="stat-card">
          <div class="stat-item">
            <div class="stat-icon user-icon">
              <el-icon><User /></el-icon>
            </div>
            <div class="stat-content">
              <div class="stat-title">用户总数</div>
              <div class="stat-number">1,234</div>
            </div>
          </div>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card class="stat-card">
          <div class="stat-item">
            <div class="stat-icon role-icon">
              <el-icon><Avatar /></el-icon>
            </div>
            <div class="stat-content">
              <div class="stat-title">角色总数</div>
              <div class="stat-number">12</div>
            </div>
          </div>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card class="stat-card">
          <div class="stat-item">
            <div class="stat-icon menu-icon">
              <el-icon><Menu /></el-icon>
            </div>
            <div class="stat-content">
              <div class="stat-title">菜单总数</div>
              <div class="stat-number">45</div>
            </div>
          </div>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card class="stat-card">
          <div class="stat-item">
            <div class="stat-icon log-icon">
              <el-icon><Document /></el-icon>
            </div>
            <div class="stat-content">
              <div class="stat-title">今日操作</div>
              <div class="stat-number">128</div>
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>

    <el-row :gutter="20" style="margin-top: 20px;">
      <el-col :span="16">
        <el-card>
          <template #header>
            <div class="card-header">
              <span>数据概览</span>
            </div>
          </template>
          <div id="chart" style="height: 300px;"></div>
        </el-card>
      </el-col>
      <el-col :span="8">
        <el-card>
          <template #header>
            <div class="card-header">
              <span>快速操作</span>
            </div>
          </template>
          <div class="quick-actions">
            <el-button type="primary" size="large" @click="goTo('users')">
              <el-icon><User /></el-icon>
              用户管理
            </el-button>
            <el-button type="success" size="large" @click="goTo('roles')">
              <el-icon><Avatar /></el-icon>
              角色管理
            </el-button>
            <el-button type="warning" size="large" @click="goTo('menus')">
              <el-icon><Menu /></el-icon>
              菜单管理
            </el-button>
            <el-button type="info" size="large" @click="goTo('operation-logs')">
              <el-icon><Document /></el-icon>
              操作日志
            </el-button>
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router'
import { User, Avatar, Menu, Document } from '@element-plus/icons-vue'

const router = useRouter()

const goTo = (path: string) => {
  router.push({ path: `/${path}` })
}
</script>

<style scoped>
.dashboard-container {
  background-color: #fff;
  padding: 20px;
  border-radius: 4px;
}

.stat-card {
  height: 100px;
}

.stat-item {
  display: flex;
  align-items: center;
}

.stat-icon {
  width: 60px;
  height: 60px;
  border-radius: 10px;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-right: 15px;
}

.user-icon {
  background-color: #409EFF;
}

.role-icon {
  background-color: #67C23A;
}

.menu-icon {
  background-color: #E6A23C;
}

.log-icon {
  background-color: #909399;
}

.stat-icon i {
  color: white;
  font-size: 24px;
}

.stat-content .stat-title {
  font-size: 14px;
  color: #909399;
  margin-bottom: 5px;
}

.stat-content .stat-number {
  font-size: 24px;
  font-weight: bold;
  color: #303133;
}

.card-header {
  font-weight: bold;
}

.quick-actions {
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.quick-actions .el-button {
  width: 100%;
  justify-content: flex-start;
}

.quick-actions .el-button i {
  margin-right: 8px;
}
</style>