<script setup lang="ts">
import { ref, PropType, onMounted } from 'vue';
import { NAvatar, NProgress } from 'naive-ui';
import { SuperChat, computeSuperChatBackground } from './super_chat'
const props = defineProps({
  data: {
    type: Object as PropType<SuperChat>,
    required: true,
  }
})
const superChatRef = ref<undefined | HTMLElement>(undefined)
const color = computeSuperChatBackground(props.data.price)
const long = props.data.end_time - props.data.timestamp
const height = Math.floor(props.data.content.length / 28 + 1) * 30

let percentage = ref(100)
onMounted(() => {
  let count = 0
  let bar = setInterval(() => {
    count += 1
    // percentage.value = 100 - (count / (long*100) *100)
    percentage.value = 100 - (count / long)
    if (count >= long*100) {
      clearInterval(bar)
    }
  }, 10)
})

</script>

<template>
  <div class="super-chat-wrapper" ref="superChatRef">
    <div class="super-chat-box">
      <div class="info">
        <n-avatar round :size="45" :src="props.data.user.avatar"
          fallback-src="https://i0.hdslb.com/bfs/face/member/noface.jpg"
          :img-props="{ class: 'avatar-img', alt: props.data.user.name }">
        </n-avatar>
        <div class="name">{{ props.data.user.name }}</div>
      </div>
      <div class="progress-bar">
        <n-progress :percentage="percentage" :fill-border-radius="0" :border-radius="0" :height="height" type="line"
          :color="color" rail-color="rgba(28, 28, 28,0.8)" :show-indicator="false">
        </n-progress>
        <div class="content">{{ props.data.content }}</div>
      </div>


    </div>

  </div>

</template>

<style scoped lang="less">
.super-chat-wrapper {
  width: 100%;
  display: flex;
  align-items: center;
  min-height: 2rem;
  color: rgb(28, 28, 28);

  .super-chat-box {
    display: flex;
    flex-flow: row wrap;
    position: relative;

    .info {
      padding-left: 1rem;
      display: inline-flex;
      width: 512px;
      align-items: center;
      background-color: azure;

      .name {
        font-size: 1.2rem;
        margin-left: 1rem;
      }
    }

    .progress-bar {
      width: 100%;
      display: inline-flex;
      flex-flow: row wrap;
      justify-content: space-between;
      padding-right: 1rem;

      .content {
        padding: 0 1rem;
        position: absolute;
        overflow-wrap: break-word;
        font-size: 1rem;
        color: #fff;
        word-break: break-word;
        width: 480px;
        background-color: transparent;
      }

    }
  }
}
</style>