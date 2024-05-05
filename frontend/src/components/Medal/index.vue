<template>
    <div class="medal" :style="{backgroundColor:color}">
        <span class="medal-name">
            {{ medal_name }}
        </span>
        <span class="medal-level" :style="{ color: color }">
            {{ props.level }}
        </span>
    </div>
</template>
<script setup lang="ts">
import { ref } from 'vue';
import {danmuLevelColor,selectColor} from '../Danmu/danmu'
const props = defineProps({
    name: {
        type: String,
        required: true
    },
    level: {
        type: Number,
        required: true
    }
})
let color = ref<danmuLevelColor>()
let medal_name= ref(props.name)
color.value = selectColor(props.level)
if (props.name.length > 3) {
    medal_name.value = props.name.slice(0,3)
}
</script>

<style scoped lang="less">
.medal{
    min-width: 15%;

    text-align: center;
    padding: 2px;
    color: white;
    display:flex;
    font-size: 15px;
    height: 1rem;
    justify-content: space-between;
    margin: 0.1rem 0.5rem 0  0 ;
    border-radius: 5px;
    line-height: 1rem;
    .medal-level{
        border-radius:2px ;
        width:1.5rem ;
        margin: 0 auto;
        background-color: white;
    }
    .medal-name{
        margin: 0 auto;
        border-radius: 5px 0 0 5px;
        width: auto;

    }
}

</style>