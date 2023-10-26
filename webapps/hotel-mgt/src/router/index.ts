import { createRouter, createWebHistory } from "vue-router";

// NOTE:
/**
 * meta 对象说明：
 * title        菜单名字/浏览器标签title
 * icon         菜单icon
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

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
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
            component: () => import("../layouts/Default.vue"),
            redirect: "/overview",
            children: [
                {
                    path: "/overview",
                    name: "overview",
                    component: import("../views/Overview.vue"),
                    meta: {
                        authRequired: true,
                        title: "总览页",
                        icon: "Monitor",
                    },
                },
            ],
        },
        {
            path: "/hotelcenter",
            component: () => import("../layouts/Default.vue"),
            redirect: "/hotelcenter/hotels",
            meta: {
                title: "酒店中心",
                icon: "Monitor",
            },
            children: [
                {
                    path: "/hotelcenter/hotels",
                    name: "hotels",
                    component: import("../views/Hotels.vue"),
                    meta: {
                        title: "酒店管理",
                        icon: "Monitor",
                    },
                },
                {
                    path: "/hotelcenter/rooms",
                    name: "rooms",
                    component: import("../views/Rooms.vue"),
                    meta: {
                        title: "客房管理",
                        icon: "Monitor"
                    },
                },
            ],
        },
        {
            path: "/business",
            component: () => import("../layouts/Default.vue"),
            redirect: "/business/users",
            meta: {
                title: "业务中心",
                icon: "Monitor",
            },
            children: [
                {
                    path: "/business/users",
                    name: "users",
                    component: import("../views/Users.vue"),
                    meta: {
                        title: "用户管理",
                        icon: "Monitor"
                    },
                },
                {
                    path: "/business/bookings",
                    name: "bookings",
                    component: import("../views/Bookings.vue"),
                    meta: {
                        title: "订单管理",
                        icon: "Monitor",
                    },
                },
            ],
        },
        {
            path: "/menucenter",
            component: () => import("../layouts/Default.vue"),
            redirect: "/menucenter/menus",
            meta: {
                title: "菜单中心",
                icon: "Monitor",
            },
            children: [
                {
                    path: "/menucenter/menus",
                    name: "menus",
                    component: import("../views/Menus.vue"),
                    meta: {
                        title: "菜单管理",
                        icon: "Monitor"
                    },
                },
                {
                    path: "/menucenter/roles",
                    name: "roles",
                    component: import("../views/Roles.vue"),
                    meta: {
                        title: "角色管理",
                        icon: "Monitor"
                    },
                },
                {
                    path: "/menucenter/permissions",
                    name: "permissions",
                    component: import("../views/Permissions.vue"),
                    meta: {
                        title: "权限管理",
                        icon: "Monitor"
                    },
                },
            ],
        },
    ],
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

