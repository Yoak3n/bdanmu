<template>
  <n-list hoverable clickable class="menu-wrapper">
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
    <n-list-item class="menu-item">
      <template #prefix>
        <n-icon>
          <settings />
        </n-icon>
      </template>

      修改
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
  SettingsOutline as Settings
} from '@vicons/ionicons5'

import {JumpToLiveRoom ,AppQuit}from './mixin'


</script>

<style lang="less" scoped>
.menu-wrapper {
  width: 100%;
  font-size: large;
  background-color: transparent;
  color: rgb(240,240,240);
  :hover+div{
    background-color: green;
  }
  .menu-item{
    width: 100%;
    height: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
    :hover+div{
        background-color: red;
        color: #000;
      }


  }
}
</style>