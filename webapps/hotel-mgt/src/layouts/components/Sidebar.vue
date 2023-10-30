<template>
    <div>
        <el-menu :default-active="activeMenu" :collapse="store.isCollapse" router>
            <SidebarItem v-for="route in routes" :key="route.path" :route="route" />
        </el-menu>
    </div>
</template>

<script setup lang="ts">
import { computed } from "vue";
import SidebarItem from "./SidebarItem.vue";
import { useRoute } from "vue-router";
import { routes } from "@/router";
import { useAppStore } from "@/stores";

const route = useRoute();
const store = useAppStore();

const activeMenu = computed(() => {
    // console.log(route.path);
    return route.path;
});

</script>

<style scoped lang="scss">

</style>