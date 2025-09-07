<template>
    <div class="base-layout">
        <div class="base-layout-header">
          <TopNavigation></TopNavigation>
        </div>
        <div class="base-layout-content">
            <RouterView></RouterView>
        </div>
    </div>
</template>

<script lang="ts" setup>
import TopNavigation from "./components/TopNavigation.vue";
</script>

<style lang="scss" scoped>
.base-layout {
    width: 100%;
    height: 100%;
    display: flex;
    flex-direction: column;
    justify-content: start;
    position: relative;
    overflow: hidden;
    .base-layout-header {
        width: 100%;
        height: 60px;
        position: relative;
        overflow: hidden;
    }
    .base-layout-content {
        width: 100%;
        height: calc(100vh - 60px);
        overflow: hidden;
    }
}
</style>