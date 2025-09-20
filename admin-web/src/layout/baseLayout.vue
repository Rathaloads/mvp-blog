<template>
  <div class="base-layout">
    <div class="base-layout__slider">
      <Slider></Slider>
    </div>
    <div class="base-layout__content">
      12312312
    </div>
  </div>
</template>

<script lang="ts" setup>
import Slider from './components/Slider.vue';
// import TopNavigation from "./components/TopNavigation.vue";
</script>

<style lang="scss" scoped>
.base-layout {
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: row;
  justify-content: start;
  position: relative;
  overflow: hidden;

  &__slider {
    height: 100%;
    background-color: var(--content-color);
  }

  &__content {
    width: 100%;
    height: 100%;
    padding: 0px 10px;
    background-color: #EAF5FE;
  }
}
</style>
