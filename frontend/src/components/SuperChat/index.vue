<script setup lang="ts">
import { ref, PropType,onMounted } from 'vue';
import {NLoadingBarProvider,useLoadingBar} from 'naive-ui';
import {SuperChat,computeSuperChatBackground} from './super_chat'

const props = defineProps({
    data:{
        type:Object as PropType<SuperChat>,
        required:true,
    }
})
const superChatRef = ref<undefined | HTMLElement>(undefined)
const loadingBar = useLoadingBar()
let color = computeSuperChatBackground(props.data.price)

onMounted(()=>{
  loadingBar.finish()
})
</script>

<template>
  <n-loading-bar-provider
      :to="superChatRef"
      :loading-bar-style="{loading:{color:'red'}}"
      container-style="position: absolute;">
    <div class="super-chat-wrapper" ref="superChatRef">
      {{ props.data.content }}
      <button @click="loadingBar.start">加载</button>
    </div>
  </n-loading-bar-provider>

</template>

<style scoped lang="less">
.super-chat-wrapper{
  width: 100%;
  display: flex;
  align-items: center;
  min-height: 2rem;
  border-radius: 10px;
  color: rgb(28, 28, 28);
  background-color: v-bind(color);

}
</style>