<script setup lang="ts">
import {ref,onMounted} from 'vue'
import {NFloatButton,NIcon,NDrawer,NDrawerContent} from 'naive-ui'
import {CloseOutline} from '@vicons/ionicons5'
import NaiveProvider from './components/NaiveProvider/index.vue'
import Menu from './components/Menu/index.vue'
import {EventsOnce} from '../wailsjs/runtime'

let drawer_open = ref(false)

const quit = ()=> {
  // Quit()
  drawer_open.value = !drawer_open.value
}

onMounted(()=>{
  EventsOnce('started',function (room) {
    localStorage.setItem('room_id',room.short_id)
  })
})

</script>

<template>
  <naive-provider>
    <router-view/>
  </naive-provider>
  <div >
    <n-float-button :right="0" :bottom="0" shape="square"  @click="quit" class="quit">
      <n-icon >
        <close-outline/>
      </n-icon>
    </n-float-button>
    <n-drawer v-model:show="drawer_open" 

    style="background-color: rgba(240,240,240,0.8);"  width="50%">
      <n-drawer-content     body-content-style="padding: 0;">
        <Menu />
      </n-drawer-content>
    </n-drawer>
  </div>

</template>

<style scoped lang="less">
.quit{
  background-color: transparent;
  color:transparent;
}
.quit:hover{
  visibility: visible;
  color:rgb(28,28,28);
  background-color: white;
}

</style>