<script setup lang="ts">
import {ref,onMounted} from 'vue'
import {NFloatButton,NIcon,NDrawer,NDrawerContent} from 'naive-ui'
import {MenuOutline} from '@vicons/ionicons5'
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
    <n-float-button :right="0" :bottom="0" shape="square"  @click="quit" class="menu">
        <n-icon >
          <menu-outline />
        </n-icon>
    </n-float-button>
    <n-drawer v-model:show="drawer_open"
    style="background-color: rgba(240,240,240,0.5);margin: 1rem 0"  width="50%">
      <n-drawer-content     body-content-style="padding: 0;">
        <Menu />
      </n-drawer-content>
    </n-drawer>
  </div>

</template>

<style scoped lang="less">
.menu{
  color:rgba(28,28,28,1);
  background-color:rgba(28,28,28,1);
}
.menu:hover{
  animation-duration: .5s;
  animation-name: fadeIn;
  background-color: rgba(240,240,240,1);
}
.menu:hover div{
  animation-duration: .5s;
  animation-name: fadeIn;
  background-color: rgba(240,240,240,1);
}
@keyframes  fadeIn{
  from {
    background-color:rgba(28,28,28,0);
  }
  to{
    background-color: rgba(240,240,240,1);
  }
}


</style>