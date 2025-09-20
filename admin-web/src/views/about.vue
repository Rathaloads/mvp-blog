<template>
  <h1>About View</h1>
  <RouterLink to="/">Home</RouterLink>
  <span>{{ globalStore.count }}</span>
  /
  <span>{{ globalStore.getCount }}</span>
  <ElButton type="primary" @click="globalStore.updateCount">Click Me</ElButton>
</template>

<script lang="ts" setup>
import { ElButton } from "element-plus";
import { useGlobalStore } from "../store/global";

const globalStore = useGlobalStore();
</script>
