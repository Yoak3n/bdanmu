<template>
  <n-list hoverable clickable class="menu-wrapper">
    <n-list-item class="menu-item" @click="JumpToLiveRoom">
      跳转直播间
    </n-list-item>
    <n-list-item class="menu-item">
      直播间信息
    </n-list-item>
    <n-list-item class="menu-item">
      跳转直播间
    </n-list-item>
    <n-list-item class="menu-item">
      退      出
    </n-list-item>
  </n-list>

</template>
<script setup lang="ts">
import { NList, NListItem } from 'naive-ui'

import {JumpToLiveRoom }from './mixin'


</script>

<style lang="less" scoped>
.menu-wrapper {
  width: 100%;
  background-color: transparent;
  .menu-item {
    width: 100%;
    height: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
  }
}
</style>