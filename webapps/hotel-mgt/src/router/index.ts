import { createRouter, createWebHistory } from "vue-router";
import MainLayout from "../layouts/MainLayout.vue";

// NOTE:
/**
 * meta 对象说明：
 * title        菜单名字/浏览器标签title
 * icon         菜单icon, 设置方式:（elemenetIcon: el-XXX; 本地svg: svg-XXX,必须存储在assets/svgIcons下）XXX 对应组件名或者svg名
 * hiddenInMenu 是否在菜单中隐藏
 */

/**
 * 设计菜单：
 * 总览页
 * 酒店中心
 *    酒店管理
 *    客房管理
 * 业务中心
 *    用户管理
 *    订单管理（预定管理）
 * 菜单中心
 *    菜单管理
 *    角色管理
 *    权限管理
 *
 */


export const routes = [
    {
        path: "/login",
        name: "login",
        component: import("../views/Login.vue"),
        meta: {
            title: "登录",
            hiddenInMenu: true,
        },
    },
    {
        path: "/",
        component: MainLayout,
        redirect: "/overview",
        children: [
            {
                path: "/overview",
                name: "overview",
                component: import("../views/Overview.vue"),
                meta: {
                    authRequired: true,
                    title: "总览页",
                    icon: "svg-apple",
                },
            },
        ],
    },
    {
        path: "/hotelcenter",
        component: MainLayout,
        redirect: "/hotelcenter/hotels",
        meta: {
            title: "酒店中心",
            icon: "el-Monitor",
        },
        children: [
            {
                path: "/hotelcenter/hotels",
                name: "hotels",
                component: import("../views/Hotels.vue"),
                meta: {
                    title: "酒店管理",
                    icon: "svg-apple",
                },
            },
            {
                path: "/hotelcenter/rooms",
                name: "rooms",
                component: import("../views/Rooms.vue"),
                meta: {
                    title: "客房管理",
                    icon: "el-Monitor"
                },
            },
        ],
    },
    {
        path: "/business",
        component: MainLayout,
        redirect: "/business/users",
        meta: {
            title: "业务中心",
            icon: "el-Monitor",
        },
        children: [
            {
                path: "/business/users",
                name: "users",
                component: import("../views/Users.vue"),
                meta: {
                    title: "用户管理",
                    icon: "el-Monitor"
                },
            },
            {
                path: "/business/bookings",
                name: "bookings",
                component: import("../views/Bookings.vue"),
                meta: {
                    title: "订单管理",
                    icon: "el-Monitor",
                },
            },
        ],
    },
    {
        path: "/menucenter",
        component: MainLayout,
        redirect: "/menucenter/menus",
        meta: {
            title: "菜单中心",
            icon: "el-Monitor",
        },
        children: [
            {
                path: "/menucenter/menus",
                name: "menus",
                component: import("../views/Menus.vue"),
                meta: {
                    title: "菜单管理",
                    icon: "el-Monitor"
                },
            },
            {
                path: "/menucenter/roles",
                name: "roles",
                component: import("../views/Roles.vue"),
                meta: {
                    title: "角色管理",
                    icon: "el-Monitor"
                },
            },
            {
                path: "/menucenter/permissions",
                name: "permissions",
                component: import("../views/Permissions.vue"),
                meta: {
                    title: "权限管理",
                    icon: "el-Monitor"
                },
            },
        ],
    },
];

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes,
    scrollBehavior(to, from, savedPosition) {
        return (
            savedPosition ||
            new Promise((resolve) => {
                setTimeout(() => resolve({ top: 0 }), 300);
            })
        );
    },
});

export default router;

