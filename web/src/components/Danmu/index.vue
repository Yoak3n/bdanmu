<template>
    <div class="danmu-wrapper">
        <div class="danmu">
            <a class="avatar" :href="`https://space.bilibili.com/${danmu.user.uid}`" target="_blank">
                <el-avatar round :size="45"
                    :src="props.danmu.user.avatar != '' ? props.danmu.user.avatar : 'https://i0.hdslb.com/bfs/face/member/noface.jpg'"
                    fallback-src="https://i0.hdslb.com/bfs/face/member/noface.jpg"
                    :img-props="{ class: 'avatar-img', alt: props.danmu.user.name }">
                </el-avatar>
            </a>
            <div class="content">
                <div class="info">
                    <div v-if="props.danmu.user.guard" class="fleet"></div>
                    <Medal v-if="props.danmu.user.medal?.name" :name="props.danmu.user.medal.name"
                        :level="props.danmu.user.medal?.level" />
                    <span class="name">{{ props.danmu.user.name }}</span>
                </div>
                <div class="message-box">
                    <div class="message" v-html="props.danmu.content">
                    </div>

                </div>
            </div>
        </div>
    </div>
</template>
<script setup lang="ts">
import type { PropType } from 'vue';
import { Danmu } from './danmu'
import Medal from '../Medal/index.vue'
const props = defineProps(
    {
        danmu: {
            type: Object as PropType<Danmu>,
            required: true
        }
    }
)


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
        padding: 0 .5rem;
        display: flex;
        overflow-wrap: break-word;
        width: 100%;

        .avatar {
            display: flex;
            width: 45px;
            height: auto;
            line-height: 100%;
            text-align: center;
            align-items: center;
            justify-content: center;

        }

        .content {
            width: 85%;
            padding: 1%;

            .message-box {
                display: flex;
                font-size: 16px;
                width: 100%;

                .message {
                    position: relative;
                    max-width: 95%;
                    background-color: #6ed7db;
                    font-size: 20px;
                    color: #fff;
                    word-wrap: break-word;
                    overflow-wrap: break-word;
                    border-radius: 10px 20px 20px 10px;
                    padding: 0 15px 0 10px;
                    align-items: center;
                    font-weight: bold;

                    &::before {
                        position: absolute;
                        content: "";
                        bottom: -5px;
                        /* 根据气泡大小调整 */
                        left: 50%;
                        margin-left: -10px;
                        /* 用于水平居中箭头 */
                        border-width: 10px 10px 0;
                        /* 根据气泡大小调整 */
                        border-style: solid;
                        border-color:  #6ed7db  transparent transparent ;
                    }
                }

            }

            .info {
                // background-color: bisque;
                border-radius: 5px;
                margin: 0;
                height: 1.5rem;
                line-height: 1.5rem;
                display: flex;
                justify-content: left;
                width: 400px;

                .name {
                    color: rgb(189, 193, 197);
                    font-size: 1rem;
                    font-weight: bold;
                }
            }

        }
    }
}
</style>