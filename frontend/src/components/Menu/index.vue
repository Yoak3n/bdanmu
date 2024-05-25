<template>
  <n-list hoverable clickable class="menu-wrapper">
    <n-list-item class="menu-item" @click="()=>{
      if ($route.path != '/dashboard'){
        $router.push({name: 'Dashboard', query: {from: $route.path.split('/')[1]}})
      }
      appStore.setDrawer(false)
      }">
      <template #prefix>
        <n-icon>
          <barcode />
        </n-icon>
      </template>
      弹幕流
    </n-list-item>
    <n-list-item class="menu-item" @click="(e)=>{
      JumpToLiveRoom(e)
      appStore.setDrawer(false)
    }">
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
    <n-list-item class="menu-item" @click="()=>{
      $router.push('/setting')
      appStore.setDrawer(false)
      }">
      <template #prefix>
        <n-icon>
          <settings />
        </n-icon>
      </template>
      设置
    </n-list-item>
    <n-list-item class="menu-item" @click="(e)=>{
      HideWindow(e)
      appStore.setDrawer(false)
    }">
      <template #prefix>
      <n-icon>
        <caret-down />
      </n-icon>
      </template>
      隐藏至托盘
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
  BarcodeOutline as Barcode,
  CaretDownCircleOutline as CaretDown
} from '@vicons/ionicons5'
import {useAppStore} from '@/store'
import {JumpToLiveRoom ,AppQuit,HideWindow}from './mixin'

const appStore = useAppStore()


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
    border-radius: 0;
  }
  .menu-item:hover{
    background-color: rgba(0,0,0, 0.5);
    box-shadow: 0 0 10px 2px #918f8f;
    transition: all 0.3s;
  }
      
}
</style>