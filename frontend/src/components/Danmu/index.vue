<template>
    <div class="danmu-wrapper" v-show="props.danmu != null">
        <div class="danmu" >
            <div class="avatar">
                <n-avatar  class="avatar-img" round size="large"
                    :src="props.danmu.user.avatar != '' ? props.danmu.user.avatar:'https://i0.hdslb.com/bfs/face/member/noface.jpg'" fallback-src="https://i0.hdslb.com/bfs/face/member/noface.jpg"
                    :img-props="{alt: props.danmu.user.name, color: 'red' }"
                    :render-placeholder="renderLoading">
                    
                </n-avatar>
            </div>

            <div class="content">
                <div class="info">
                    <div v-if="props.danmu.user.guard" class="capiton"></div>
                    <Medal v-if="props.danmu.user.medal?.name" :name="props.danmu.user.medal.name"
                        :level="props.danmu.user.medal?.level" />
                    <span class="name">{{ props.danmu.user.name }}</span>
                </div>
                <div class="message" v-html="props.danmu.content">
                  </div>
            </div>
        </div>
    </div>
</template>
<script setup lang="ts">
import type { PropType, VNodeChild } from 'vue';
import { NAvatar } from 'naive-ui'

import { Danmu } from './danmu'
import Medal from '../Medal/index.vue'
import { h } from 'vue';

const props = defineProps(
    {
        danmu: {
            type: Object as PropType<Danmu>,
            required: true
        }
    }
)

const renderLoading = (): VNodeChild => {
    const vnode = h('div', { class: 'danmu', style: { width: '100%', height: '100%', background: 'red' } }, 123465)
    return vnode
}

</script>

<style scoped lang="less">
.fade-enter-active {
    transition: all 0.5s ease;
}

.fade-enter-from,
.fade-leave-to {
    opacity: 0;
    transform: translateY(50%);
}

.danmu-wrapper {
    .danmu {
        display: flex;
        overflow-wrap: break-word;
        width: 100%;

        .avatar {
            width: 15%;
            text-align: center;

            .avatar-img {
                // width: 80%;
                margin: 0 auto;
            }
        }

        .content {
            width: 85%;
            padding: 1%;

            .info {
                // background-color: bisque;
                border-radius: 5px;
                margin: 0 1%;
                height: 1.5rem;
                line-height: 1.5rem;
                display: flex;
                justify-content: left;

                .name {
                    color: rgb(189, 193, 197);
                    font-size: 1rem;
                    font: bold;
                    color: 36px;
                }
            }

        }
    }
}
</style>