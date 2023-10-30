<template>
    <el-icon v-show="iconComp">
        <component :is="iconComp" />
    </el-icon>
</template>

<script setup lang="ts">
import { defineAsyncComponent, onBeforeMount, ref } from "vue";

let { icon } = defineProps({
    icon: {
        type: String,
    },
});

// icon 组件
const iconComp = ref<any>(null);

onBeforeMount(() => {
    if (!icon) {
        return;
    }
    // 以 "el-" 开头, element ui icon
    if (icon?.indexOf("el-") === 0) {
        iconComp.value = icon.replace("el-", "");
    } else {
        // 本地 svg icon
        icon = icon?.replace("svg-", "");

        // TODO: 在菜单搜索折叠切换下，异步组件的ICON会有闪烁问题
        iconComp.value = defineAsyncComponent(() => {
            return new Promise((resolve, reject) => {
                resolve(import(`../../assets/svgIcons/${icon}.svg?component`))
            })
        })
    }
});
</script>

<style scoped lang="scss"></style>
