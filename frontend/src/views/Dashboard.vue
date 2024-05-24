<template>
  <div class="home-wrapper" ref="containerRef">
    <n-affix :trigger-top="0" :listen-to="() => containerRef" class="super-chat-box">
      <Transition name="sc">
        <div v-if="superChats.length > 0">
          <SuperChatbox v-for="superChat in superChats" :key="superChat.message_id" :data="superChat" />
        </div>
      </Transition>

      <!-- <button @click="testSuperChat">test</button> -->

    </n-affix>
    <n-infinite-scroll class="danmu-box">

      <div v-for="(danmu, index) in danmus" :id="index == danmus.length - 1 ? 'bottom' : ''" :key="danmu.message_id">
        <Danmubox :danmu="danmu" />
      </div>

    </n-infinite-scroll>
  </div>
</template>
<script setup lang="ts">
import { ref, onMounted, nextTick } from 'vue';
import { NAffix, NInfiniteScroll, useMessage } from 'naive-ui'
import { useRoute } from 'vue-router';
import { useRoomsStore} from '@/store/room'
import Danmubox from '../components/Danmu/index.vue'
import SuperChatbox from '../components/SuperChat/index.vue'
import type { Danmu } from '../components/Danmu/danmu'
import type { SuperChat } from '../components/SuperChat/super_chat';
import type { User } from '../components/types'
import { EventsOn, EventsOff, EventsEmit } from '../../wailsjs/runtime'

const roomsStore = useRoomsStore()
const containerRef = ref<HTMLElement | undefined>(undefined)
const $message = useMessage()
const $route = useRoute()
let danmus = ref<Array<Danmu>>([])
let superChats = ref<Array<SuperChat>>([
  // {
  //   room_id: 42062,
  //   message_id: "123123",
  //   content: "哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈",
  //   timestamp: 0,
  //   end_time: 0,
  //   price: 30,
  //   user: {
  //     name: "hello",
  //     uid: 0,
  //     avatar: "https://i0.hdslb.com/bfs/face/member/noface.jpg",
  //     sex: 0,
  //     guard: false
  //   }
  // }
])
onMounted(async () => {
  if ($route.query.from == "login") {
    EventsOn('started', function (room) {
      roomsStore.setRoomTitle(room.title)
      roomsStore.setRoomId(room.short_id)
      $message.create("已连接房间：" + room.short_id, { duration: 5000 })
      EventsOff("danmu", "user", "superChat")
      EventsOn("danmu", pushDanmu)
      EventsOn("user", updateDanmu)
      EventsOn("superChat", pushSuperChat)
    })
  } else if ($route.query.from == "setting") {
    danmus.value = []
    EventsOn('started', function (room) {
      roomsStore.setRoomTitle(room.title)
      roomsStore.setRoomId(room.short_id)
      $message.create("已连接房间：" + room.short_id, { duration: 5000 })
      EventsOff("danmu", "user", "superChat")
      EventsOn("danmu", pushDanmu)
      EventsOn("user", updateDanmu)
      EventsOn("superChat", pushSuperChat)
    })
  } else {
    EventsOn('started', function (room) {
      roomsStore.setRoomTitle(room.title)
      roomsStore.setRoomId(room.short_id)
      $message.create("已连接房间：" + room.short_id, { duration: 5000 })
      EventsOn("danmu", pushDanmu)
      EventsOn("user", updateDanmu)
      EventsOn("superChat", pushSuperChat)
    })
    EventsEmit("start")
  }



})

// const testSuperChat = () => {
//   let index = 0
//   index += 1


//   return pushSuperChat({
//     room_id: 6154037,
//     message_id: '6522809',
//     content: '猪播完美预测自己第一个死，这就是鹅鸭杀高玩吗',
//     timestamp: 1677069035 + index,
//     end_time: 1677069095 + index,
//     price: 30,
//     user: {
//       name: '界原虚',
//       uid: 294094150,
//       avatar: 'https://i1.hdslb.com/bfs/face/7a11b48e0a3055e220fa8b4c7d938cd4bcac2577.jpg',
//       sex: -1,
//       guard: true,
//     }
//   })
// }

const pushSuperChat = (super_chat: SuperChat) => {
  superChats.value.push(super_chat)
  setTimeout(() => {
    superChats.value.splice(superChats.value.findIndex((sc) => sc.end_time == super_chat.end_time), 1)
    console.log("delete")
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
.home-wrapper {
  height: 100%;
  width: 100%;

  // margin: 0 2rem;
  .super-chat-box {
    width: 100%;
    z-index: 9;
  }

  .danmu-box {
    height: 100%;
    width: 100%;
    overflow-y: scroll;
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