<template>
    <!-- 拥有子路由字 > 1-->
    <el-sub-menu :index="route.path" v-if="route.children && route.children.length > 1 && !route?.meta?.hiddenInMenu">
        <template #title>
            <SidebarIcon :icon="route?.meta?.icon" />
            <span>{{ route?.meta?.title }}</span>
        </template>
        <!-- 循环渲染 -->
        <SidebarItem v-for="ch in route.children" :key="ch.path" :route="ch" />
    </el-sub-menu>

    <!-- 子路由==1 -->
    <el-menu-item :index="route.children[0].path" v-else-if="route.children && route.children.length == 1 && !route?.meta?.hiddenInMenu">
        <SidebarIcon :icon="route.children[0]?.meta?.icon" />
        <template #title>{{ route.children[0]?.meta?.title }}</template>
    </el-menu-item>

    <!-- 没有子路由 -->
    <el-menu-item v-else :index="route.path" v-if="!route?.meta?.hiddenInMenu">
        <SidebarIcon :icon="route?.meta?.icon" />
        <template #title>{{ route?.meta?.title }}</template>
    </el-menu-item>
</template>

<script setup lang="ts">
import SidebarIcon from "./SidebarIcon.vue"
defineProps({
    route: {
        type: Object,
        required: true,
    },
});


</script>

<style scoped lang="scss"></style>
