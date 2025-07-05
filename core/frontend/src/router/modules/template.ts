import type { RouteRecordRaw } from "vue-router";
import { Layout } from '@/router/constant'
import { isDev } from "@/utils";
export default {
    path: "/template",
    component: Layout,
    meta: {
        sort: 3,
        title: "template",
        key: 'template',
        titleKey: 'layout.menu.template',
        hidden: !isDev
    },
    children: [
        {
            path: "/template",
            name: "template",
            component: () => import("@/views/template/index.vue")
        },
        {
            path: "ai-template",
            name: "ai-template",
            component: () => import("@/views/template/pages/AITemplate/index.vue")
        }
    ]
} as RouteRecordRaw
