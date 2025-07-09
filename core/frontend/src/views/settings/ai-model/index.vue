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
                    <n-button type="primary" @click="openModelManager">
                        <template #icon>
                            <i class="i-ic:baseline-add-circle"></i>
                        </template>
                        Add model
                    </n-button>
                </div>
            </div>
        </n-card>

        <n-card>
            <div class="model-item">
                <div class="item-header">
                    <div class="header-left">
                        <i class="i-ai:deep-seek w-10 h-10"></i>
                        <span class="model-tit">DeepSeek</span>
                    </div>
                    <div class="header-right">
                        <div class="remove-model">
                            <i class="i-material-symbols:close-rounded text-5"></i>
                        </div>
                    </div>
                </div>
                <n-divider style="margin: 10px 0;"></n-divider>
                <div class="item-desc">
                    <div class="desc">
                        <span class="label">调用次数：</span>
                        <span>20次</span>
                    </div>
                    <div class="desc">
                        <span class="label">Token使用量：</span>
                        <span>2018</span>
                    </div>
                </div>
            </div>
        </n-card>
    </div>

    <!-- Model maneger modal -->
    <ModelManager ref="modelManagerRef" />
</template>

<script setup lang="tsx">
    import ModelManager from './components/ModelManager.vue';
    import { ModelStore } from './dto';
    import { useModelManagerStore } from './store/index,';
    const modelManagerStore = useModelManagerStore()
    provide<ModelStore>("modelStore", modelManagerStore)
    const modelManagerRef = ref()

    /**
     * @description Open model manager modal
     */
    function openModelManager() {
        modelManagerRef.value.open()
    }
</script>

<style scoped lang="scss">
    @use "@/styles/index.scss" as base;
    @use "@/views/domain/pages/editDomain/components/mixin.scss";

    .page-tit {
        @include mixin.page-tit;
    }

    .model-item {
        .item-header {
            @include base.row-flex;
            justify-content: space-between;

            .header-left {
                @include base.row-flex-start;
                gap: 10px;

                .model-tit {
                    font-size: 16px;
                    font-weight: bold;
                }
            }

            .header-right {
                .remove-model {
                    padding: 5px;
                    cursor: pointer;
                    @include base.row-flex-center;
                    transition: .15s all ease-in-out;

                    &:hover {
                        background: var(--color-bg-4)
                    }
                }
            }
        }

        .item-desc {
            @include base.row-flex-start;
            gap: 15px;

            .label {
                font-weight: bold;
            }
        }
    }
</style>