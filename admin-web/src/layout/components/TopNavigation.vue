

<template>
    <div class="top-navigation w-full h-full flex flex-row justify-start items-center gap-[10px]">
      <div class="top-navigation-logo w-[160px] h-full flex flex-row justify-center items-center">
        <span class="font-bold text-[#fff] text-[18px]">TMAN</span>
      </div>
      <div class="top-navigation-main flex-1 flex flex-row justify-between items-center pr-[10px] gap-[10px]">
        <div class="navigation flex-1 h-full">
          <ElMenu
            class="w-full"
            mode="horizontal"
          >
            <template v-for="item in menuList" :key="item.id">
              <ElMenuItem :index="item.id+''" >
                {{ item.name }}
              </ElMenuItem>
            </template>
            
          </ElMenu>
        </div>
        <div class="profile w-[200px]">22222222222222</div>
      </div>
    </div>
</template>

<script setup lang="ts">
import { ElMenu, ElMenuItem } from 'element-plus';

const menuList = [
  {
    id: 1,
    icon: "",
    path: "",
    text: "菜单1",
    name: "menu1"
  },
  {
    id: 2,
    icon: "",
    path: "",
    text: "菜单2",
    name: "menu2"
  },
  {
    id: 3,
    icon: "",
    path: "",
    text: "菜单3",
    name: "menu3"
  }
]

</script>

<style lang="scss">
.top-navigation {
  background-color: var(--primary2);

  .top-navigation-logo {
    background-color: var(--primary1);
  }
}
</style>