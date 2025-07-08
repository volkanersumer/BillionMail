<template>
    <div class="max-w-4xl mx-auto px-4 py-8">
        <n-card class="mb-5">
            <div class="page-tit" style="margin: 0;">
                <div class="left-tit">
                    <div class="back-tool">
                        <i class="i-ri:apps-fill text-6"></i>
                    </div>
                    <span class="tit-content">
                        AI Model
                    </span>
                </div>
                <div class="right-tit">
                    <n-button type="primary">
                        <template #icon>
                            <i class="i-ai:test w-5 h-5"></i>
                        </template>
                        Add model
                    </n-button>
                </div>
            </div>
        </n-card>

        <n-card>
            <div class="model-item">
                <div>

                    <span>Deep Seek</span>
                    <span></span>
                </div>
            </div>
        </n-card>
    </div>
</template>

<script setup lang="ts">

</script>

<style scoped lang="scss">
    @use "@/views/domain/pages/editDomain/components/mixin.scss";
    .page-tit {
        @include mixin.page-tit;
    }
</style>