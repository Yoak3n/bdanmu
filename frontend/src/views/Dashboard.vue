<template>
  <div class="dashboard-wrapper" ref="containerRef">
    <n-affix :trigger-top="0" :listen-to="() => containerRef" class="super-chat-box">
      <Transition name="sc">
        <div v-if="superChats.length > 0">
          <SuperChatbox v-for="superChat in superChats" :key="superChat.message_id" :data="superChat" />
        </div>
      </Transition>
      <!-- <button @click="testSuperChat">test</button> -->
    </n-affix>
    <n-infinite-scroll class="danmu-box" style="z-index: -1;">
      <transition-group name="fade" tag="div">
          <Danmubox v-for="(danmu, index) in danmus" :id="index == danmus.length - 1 ? 'bottom' : ''" :key="danmu.message_id" class="danmu-item" :danmu="danmu" />
      </transition-group>
    </n-infinite-scroll>
  </div>
</template>
<script setup lang="ts">
import { ref, onMounted, nextTick,onActivated} from 'vue';
import { NAffix, NInfiniteScroll } from 'naive-ui'
import { useRoute,useRouter } from 'vue-router';
import { useRoomStore} from '@/store'
import Danmubox from '../components/Danmu/index.vue'
import SuperChatbox from '../components/SuperChat/index.vue'
import type { Danmu } from '../components/Danmu/danmu'
import type { SuperChat } from '../components/SuperChat/super_chat';
import type { User,Room } from '../components/types'
import { EventsOn, EventsOff, EventsEmit } from '../../wailsjs/runtime'

const roomsStore = useRoomStore()
const containerRef = ref<HTMLElement | undefined>(undefined)
const $route = useRoute()
const $router = useRouter()
let danmus = ref<Array<Danmu>>([])
let superChats = ref<Array<SuperChat>>([])

onActivated(() => {
  if ($route.query.from == "login") {
  } else if ($route.query.from == "setting") {
    danmus.value = []
  }else{
    if (roomsStore.room_id == 0 || roomsStore.room_id == null) {
      window.$message.error("未找到直播间",{keepAliveOnHover: true})
      $router.push({name: 'Setting', query: {from: 'dashboard'}})
    }else{
      EventsEmit("change",roomsStore.room_id)
    }
  }
})

onMounted(() => {
    EventsOn('started', function (room:Room) {
      roomsStore.setRoomTitle(room.title)
      roomsStore.setRoomId(room.short_id)
      window.$message.create("已连接房间：" + room.short_id, { duration: 5000 })
      EventsOff("danmu", "user", "superChat")
      EventsOn("danmu", pushDanmu)
      EventsOn("user", updateDanmu)
      EventsOn("superChat", pushSuperChat)
    })
    EventsOn('error', function (err:string) {
      window.$message.error(err,{keepAliveOnHover: true,duration: 5000})
      $router.push({name: 'Setting', query: {from: 'dashboard'}})
    })
})


const pushSuperChat = (super_chat: SuperChat) => {
  superChats.value.push(super_chat)
  setTimeout(() => {
    superChats.value.splice(superChats.value.findIndex((sc) => sc.end_time == super_chat.end_time), 1)
  }, (super_chat.end_time - super_chat.timestamp) * 1000)
}

const pushDanmu = (danmu: Danmu) => {
  if (danmus.value.length > 200) {
    danmus.value.shift()
  }
  danmus.value.push(danmu)
  nextTick(() => {
    const bottom = document.getElementById("bottom")
    bottom?.scrollIntoView({ behavior: "smooth", block: "center", inline: "end" });
  })

}

const updateDanmu = (user: User) => {
  danmus.value.forEach(danmu => {
    if (danmu.user.uid == user.uid) {
      danmu.user.avatar = user.avatar
      danmu.user.fans_count = user.fans_count
      danmu.user.sex = user.sex
    }
  })

}

</script>

<style scoped lang="less">
.fade-move{
transition: all .5s ease;
}
.fade-leave-active,
.fade-enter-active {
  transition: all .5s ease;
}
.fade-leave-to,
.fade-enter-from{
  opacity: 0;
  transform: translateX(-30px);
}
.fade-leave-active{
  position: absolute;
}
.dashboard-wrapper {
  height: 100%;
  width: 100%;

  // margin: 0 2rem;
  .super-chat-box {
    width: 100%;
  }

  .danmu-box {
    height: 100%;
    width: 100%;
    overflow-y: scroll;
    position: relative;
  }

}

.sc-leave-active {
  transition: opacity 1s ease;
}

.sc-leave-from {
  opacity: 1;
}

.sc-leave-to {
  opacity: 0;
}
</style>../components/types