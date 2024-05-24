<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { NFloatButton, NIcon, NDrawer, NDrawerContent} from 'naive-ui'
import { MenuOutline } from '@vicons/ionicons5'
import NaiveProvider from './components/NaiveProvider/index.vue'
import Menu from './components/Menu/index.vue'
import { useRoute } from 'vue-router'
import { useRoomsStore } from '@/store/room'
import { storeToRefs } from 'pinia'

const roomsStore = useRoomsStore()
const $route = useRoute()
let drawer_open = ref(false)

let { room_title:title } = storeToRefs(roomsStore)
onMounted(() => {
  roomsStore.syncAuth()
})

</script>

<template>
  <naive-provider>
    <router-view v-slot="{ Component }">
      <keep-alive>
        <component :is="Component" v-if="$route.meta.keepAlive" />
      </keep-alive>
      <component :is="Component" v-if="!$route.meta.keepAlive"/>
    </router-view>
  </naive-provider>
  <div>
    <n-float-button :right="0" :bottom="0" shape="square" @click="drawer_open = !drawer_open" class="menu">
      <n-icon>
        <menu-outline />
      </n-icon>
    </n-float-button>
    <n-drawer v-model:show="drawer_open" show-mask="transparent"
      style="background-color: rgba(30,30,30,0.6);margin: 3rem 0;border: 1px solid rgba(240,240,240,0.6);border-radius:  10px 0 0 10px"
      width="40%">
      <n-drawer-content body-content-style="padding: 0;">
        <template #header >
          <div style="color:rgb(240,240,240);cursor: default;user-select: none" onselectstart="return false" unselectable="on">
          {{ title }}
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