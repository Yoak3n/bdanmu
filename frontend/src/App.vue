<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { NFloatButton, NIcon, NDrawer, NDrawerContent } from 'naive-ui'
import { MenuOutline } from '@vicons/ionicons5'
import NaiveProvider from './components/NaiveProvider/index.vue'
import Menu from './components/Menu/index.vue'
import { EventsOnce } from '../wailsjs/runtime'

let drawer_open = ref(false)

const open_drawer = () => {
  // Quit()
  drawer_open.value = !drawer_open.value
}

onMounted(() => {
  EventsOnce('started', function (room) {
    localStorage.setItem('room_id', room.short_id)
  })
})

</script>

<template>
  <naive-provider>
    <router-view />
  </naive-provider>
  <div>
    <n-float-button :right="0" :bottom="0" shape="square" @click="open_drawer" class="menu">
      <n-icon>
        <menu-outline />
      </n-icon>
    </n-float-button>
    <n-drawer v-model:show="drawer_open" show-mask="transparent"
      style="background-color: rgba(30,30,30,0.6);margin: 3rem 0;border: 1px solid rgba(240,240,240,0.6);border-radius:  10px 0 0 10px"
      width="40%">
      <n-drawer-content body-content-style="padding: 0;">
        <template #header>
          <div style="color:rgb(240,240,240);">
            直播间标题  
          </div>
        </template>
        <Menu />
      </n-drawer-content>
    </n-drawer>
  </div>

</template>

<style scoped lang="less">
.menu {
  color: rgba(28, 28, 28, 1);
  background-color: rgba(28, 28, 28, 1);
}

.menu:hover {
  animation-duration: .5s;
  animation-name: fadeIn;
  background-color: rgba(240, 240, 240, 1);
}

.menu:hover div {
  animation-duration: .5s;
  animation-name: fadeIn;
  background-color: rgba(240, 240, 240, 1);
}

@keyframes fadeIn {
  from {
    background-color: rgba(28, 28, 28, 0);
  }

  to {
    background-color: rgba(240, 240, 240, 1);
  }
}
</style>