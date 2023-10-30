<template>
    <div class="layout-container">
        <header class="layout-header">
            <NavTopbar />
        </header>

        <section class="layout-body" :style="{ paddingLeft: store.sidebarWidth }">
            <div class="layout-body-sidebar">
                <el-scrollbar>
                    <Sidebar />
                </el-scrollbar>
            </div>

            <div class="layout-body-app-main">
                <AppMain />
                <el-button @click="handleClick">展开/收缩</el-button>
            </div>
        </section>
    </div>
</template>

<script setup lang="ts">
import NavTopbar from "./components/NavTopbar.vue";
import Sidebar from "./components/Sidebar.vue"; 
import AppMain from "./components/AppMain.vue";
import { useAppStore } from "@/stores";

const store = useAppStore();

function handleClick() {
    store.setMenuStatus();
}
</script>

<style scoped lang="scss">
.layout-container {
    .layout-header {
        display: flex;
        align-items: center;
        width: 100%;
        height: 48px;
        box-sizing: border-box;
        position: fixed;
        top: 0;
        right: 0;
        z-index: 9;
        padding: 0 16px;
        color: #fff;
        background-color: #000000;
    }

    .layout-body {
        display: grid;
        grid-template-columns: auto 1fr;
        position: relative;
        top: 48px;
        transition: all 0.2s ease;
    }

    .layout-body-sidebar {
        max-width: 210px;
        height: calc(100vh - 48px);
        max-height: calc(100vh - 48px);
        position: fixed;
        top: 48px;
        left: 0;
        bottom: 0;
        overflow: hidden;

        // el-menu border-right 取消，采用这个
        border-right: 1px solid var(--el-menu-border-color);
    }

    .layout-body-app-main {
        position: relative;
    }

    // 取消 el-menu border-right
    :deep(.el-menu) {
        border-right: none;
    }

    // 作用于 Sidebar 组件样式
    :deep(.el-menu:not(.el-menu--collapse)) {
        width: 210px;
    }
}
</style>