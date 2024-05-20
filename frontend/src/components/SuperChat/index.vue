<script setup lang="ts">
import { ref, PropType } from 'vue';
import {NFlex} from 'naive-ui';
import {SuperChat,computeSuperChatBackground} from './super_chat'

const props = defineProps({
    data:{
        type:Object as PropType<SuperChat>,
        required:true,
    }
})
const superChatRef = ref<undefined | HTMLElement>(undefined)

let color = computeSuperChatBackground(props.data.price)

</script>

<template>
 <div class="super-chat-wrapper" ref="superChatRef">
    <n-flex justify="space-between" class="super-chat" >
        {{ props.data.content }}
    </n-flex>
 </div>
</template>

<style scoped lang="less">
.super-chat-wrapper{
    padding: 10px;
    width: 100%;
    border-radius: 10px;
    color: rgb(28, 28, 28);
    background-color: v-bind(color);
    .super-chat{
        width: 100%;
    }
}
</style>