<template>
    <div class="content-wrapper">
        <n-card class="mb-5">
            <!-- title -->
            <div class="page-tit">
                <div class="left-tit">
                    <div class="back-tool">
                        <i class="i-cuida:mail-outline text-7"></i>
                    </div>
                    <span class="tit-content">
                        Mail Domain Configuration
                    </span>
                </div>
                <div class="right-tit">
                    <n-button type="primary" @click="updateDomain">
                        <template #icon>
                            <i class="i-mingcute:save-2-line text-5"></i>
                        </template>
                        Save
                    </n-button>
                </div>
            </div>
        </n-card>
        <n-card>
            <!-- form data -->
            <n-form>
                <n-form-item>
                    <template #label><span class="form-label">Domain</span></template>
                    <n-input v-model:value="domainTit"></n-input>
                </n-form-item>
                <n-form-item label="">
                    <template #label><span class="form-label">Domain Quota</span></template>
                    <div class="flex justify-between gap-5 items-center w-100%">
                        <n-input v-model:value="quota as unknown as string"></n-input>
                        <n-select :options="uinitOptions" class="w-20" v-model:value="unit"></n-select>
                    </div>
                </n-form-item>
                <n-form-item>
                    <template #label><span class="form-label">MailBox Count</span></template>
                    <n-input v-model:value="mailboxes as unknown as string"></n-input>
                </n-form-item>
                <n-form-item>
                    <template #label><span class="form-label">Catch all</span></template>
                    <n-input v-model:value="catch_email"></n-input>
                </n-form-item>
            </n-form>
        </n-card>

        <!-- <n-card class="my-5">
            <div class="switch-settings">
                <div class="switch-item">
                    <div class="label">Track Email Opens</div>
                    <n-switch></n-switch>
                </div>
                <div class="switch-item" style="margin-bottom: 0;">
                    <div class="label">Track Link Clicks</div>
                    <n-switch></n-switch>
                </div>
            </div>
        </n-card> -->

        <!-- <n-card>
            <div class="page-tit mb-5">
                <div class="left-tit">
                    <div class="back-tool">
                        <i class="i-ri:rss-fill text-6"></i>
                    </div>
                    <span class="tit-content">
                        Subscription Management
                    </span>
                </div>
            </div>

            <div class="switch-settings">
                <div class="switch-item">
                    <div class="label">Include Unsubscribe Link</div>
                    <n-switch></n-switch>
                </div>
                <div class="switch-item" style="margin-bottom: 0;">
                    <div class="label">Inlcude Preferences Link</div>
                    <n-switch></n-switch>
                </div>
            </div>
        </n-card> -->
    </div>
</template>

<script setup lang="ts">
    
    import { getDomainDetail,updateDomain } from "../controller/domainConfiguration.controller"
    import { getEditDomainStoreData } from "../store"
    const {
        domainTit,
        quota,
        unit,
        mailboxes,
        catch_email
    } = getEditDomainStoreData()
    const route = useRoute()
    const domain = route.params.domain as string
    const uinitOptions = ref([
        {
            label: "B",
            value: "B"
        },
        {
            label: "KB",
            value: "KB"
        },
        {
            label: "MB",
            value: "MB"
        },
        {
            label: "GB",
            value: "GB"
        },
        {
            label: "TB",
            value: "TB"
        },
    ])

    getDomainDetail(domain)

</script>

<style scoped lang="scss">
    @use "@/styles/index" as base;
    @use "./mixin.scss" as mixin;

    .content-wrapper {
        @include mixin.content-wrapper;

        // title
        .page-tit {
            @include mixin.page-tit;
        }

        .form-label {
            @include mixin.form-label;
        }

        // Switch settings
        .switch-settings {
            @mixin settings-label {
                font-size: 14px;
                font-weight: bold;
            }

            .switch-item {
                @include base.row-flex;
                justify-content: space-between;
                margin-bottom: 15px;

                .label {
                    color: var(--color-text-2);
                    @include settings-label()
                }
            }
        }
    }


</style>