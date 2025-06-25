import type { RouteRecordRaw } from "vue-router";

export default {
    path: "/template",
    meta: {
        sort: 3,
        title:"template",
        key: 'template',
        titleKey: 'layout.menu.template',
    },
    component: () => import("@/views/template/index.vue")
} as RouteRecordRaw