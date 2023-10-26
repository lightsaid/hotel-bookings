<template>
    <div>
        <h1>Default Layout</h1>
        <router-view v-slot="{ Component, route }">
            <keep-alive>
                <component :is="Component" :key="route.path"></component>
            </keep-alive>
        </router-view>
    </div>
</template>

<script setup lang="ts"></script>

<style scoped lang="scss"></style>
