<template>
    <el-menu>
        <el-sub-menu index="1">
          <template #title>
            <span>Navigator One</span>
          </template>
          <el-menu-item index="1-1">item one</el-menu-item>
          <el-menu-item index="1-2">item two</el-menu-item>
        </el-sub-menu>
        <el-sub-menu index="2">
          <template #title>
            <span>Navigator One</span>
          </template>
          <el-menu-item index="2-1">item one</el-menu-item>
          <el-menu-item index="2-2">item two</el-menu-item>
        </el-sub-menu>
      </el-menu>
</template>

<script lang="ts" setup>
import { ElMenu, ElSubMenu, ElMenuItem } from 'element-plus';
</script>

<style lang="scss" scoped>
</style>
