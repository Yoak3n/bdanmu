<template>
  <n-list hoverable clickable class="menu-wrapper">
    <n-list-item class="menu-item" @click="()=>{$router.push('/')}">
      <template #prefix>
        <n-icon>
          <barcode />
        </n-icon>
      </template>
      弹幕流
    </n-list-item>
    <n-list-item class="menu-item" @click="JumpToLiveRoom">
      <template #prefix>
        <n-icon>
          <book-icon />
        </n-icon>
      </template>
      跳转直播间
    </n-list-item>
    <n-list-item class="menu-item">

      <template #prefix>
        <n-icon>
          <person-icon />
        </n-icon>
      </template>
      直播间信息
    </n-list-item>
    <n-list-item class="menu-item" @click="()=>{$router.push('/setting')}">
      <template #prefix>
        <n-icon>
          <settings />
        </n-icon>
      </template>
      设置
    </n-list-item>
    <n-list-item class="menu-item" @click="AppQuit">
      <template #prefix>
      <n-icon>
        <logout />
      </n-icon>
      </template>
      退出
    </n-list-item>
  </n-list>

</template>
<script setup lang="ts">
import { NList, NListItem ,NIcon} from 'naive-ui'
import {
  BookOutline as BookIcon,
  PersonOutline as PersonIcon,
  LogOutOutline as Logout,
  SettingsOutline as Settings,
  BarcodeOutline as Barcode
} from '@vicons/ionicons5'

import {JumpToLiveRoom ,AppQuit}from './mixin'


</script>

<style lang="less" scoped>
.menu-wrapper {
  width: 100%;
  font-size: large;
  background-color: transparent;
  color: rgb(240,240,240);
  .menu-item{
    width: 100%;
    height: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
    border-radius: 0;
  }
  .menu-item:hover{
    background-color: rgba(240, 240, 240, 0.5);
    // color: rgb(28,28,28);
      
  }
}
</style>