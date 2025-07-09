<template>
    <n-modal preset="card" draggable :close-on-esc="false" :mask-closable="false" title="Model manager"
        class="w-180 customer-modal" segmented v-model:show="show">
        <div class="manager-wrapper">
            <!-- Provider list -->
            <div class="left-provider">
                <div class="provider-list">
                    <div :class="['provider-item', { active: item.supplierName == currentProvider?.supplierName }]"
                        v-for="item in providerList" :key="item.supplierName"
                        @click="getModelListBySupplier(item)">
                        <i class="i-ai:deep-seek w-8 h-8"></i>
                        <span class="tit">{{ item.supplierTitle }}</span>
                    </div>
                </div>
                <div class="add-privider">
                    <n-button @click="handleAddProvider">
                        <template #icon>
                            <i class="i-ic:baseline-add-circle"></i>
                        </template>
                        Add model service provider
                    </n-button>
                </div>
            </div>

            <!-- Configuration -->
            <n-spin :show="configurationLoading">
                <div class="right-configuration">
                    <div class="top-switch">
                        <div class="switch-info">
                            <i class="i-ic:baseline-error text-5"></i>
                            <span class="tit">DeepSeek</span>
                        </div>
                        <n-switch></n-switch>
                    </div>
                    <div class="center-api">
                        <n-form class="mt-15px">
                            <n-form-item label="API密钥">
                                <div class="w-100%">
                                    <n-input-group>
                                        <n-input v-model:value="currentProvider.apiKey"></n-input>
                                        <n-button>检查</n-button>
                                    </n-input-group>
                                    <n-button text type="info" class="mt-5px">点击获取密钥</n-button>
                                </div>
                            </n-form-item>
                            <n-form-item label="API地址">
                                <div class="w-100%">
                                    <n-input v-model:value="currentProvider.baseUrl"></n-input>
                                    <span class="text-[var(--color-text-3)] mt-5">示例: https://api.deepseek.com/v1</span>
                                </div>
                            </n-form-item>
                        </n-form>
                        <n-button type="primary">保存API</n-button>
                    </div>
                    <div class="bottom-model-list">
                        <div class="header-tit">
                            <div>
                                <span class="tit">模型</span>
                                <span class="sub-tit">默认从/models获取所有模型</span>
                            </div>
                            <div class="add-model">
                                <i class="i-ic:baseline-control-point text-4"></i>
                                <span class="tit">新建模型</span>
                            </div>
                        </div>
                        <!--  <n-scrollbar style="height: 100px;">
                        <div class="model-list">
                            <div class="model-item">
                                <i class="i-ic:round-check-circle-outline text-5"></i>
                                <span class="model-name">DeepSeek/deepseek-reasoner</span>
                            </div>
                        </div>
                    </n-scrollbar> -->
                        <n-select :options="modelList" label-field="title" value-field="title"
                            v-model:value="modelName"></n-select>
                    </div>
                </div>
            </n-spin>
        </div>
        <template #footer>
            <div class="flex justify-end gap-5">
                <n-button>Cancel</n-button>
                <n-button type="primary">Create</n-button>
            </div>
        </template>
    </n-modal>

    <!-- Add Provider -->
    <AddProvider ref="addProviderRef" />

</template>

<script setup lang="ts">
    import AddProvider from './AddProvider.vue';
    import { ModelStore, Provider } from '../dto';
    import { getModelList, getProviderList } from "../controller"
    const modelStore = inject<ModelStore>("modelStore")!
    const {
        providerList,
        modelList,
        modelName,
        currentProvider,
        configurationLoading,
        addProviderRef
    } = modelStore
    const show = ref(false)
    /**
     * @description open modal
     */
    function open() {
        /**
         * @description Get provider list
         */
        getProviderList(modelStore)
        show.value = true
    }

    /**
     * @description close modal
     */
    function close() {
        show.value = false
    }
    defineExpose({
        open,
        close
    })

    /**
     * @description Get model list
     */
    function getModelListBySupplier(provider: Provider) {
        currentProvider.value = provider
        getModelList(provider.supplierName, modelStore)
    }

    /**
     * @description Handle add provider
     */
    function handleAddProvider() {
        addProviderRef.value.open()
    }

</script>

<style scoped lang="scss">
    @use "@/styles/index.scss" as base;

    .manager-wrapper {
        height: 480px;
        display: grid;
        grid-template-columns: 230px 1fr;
        border-top: 1px solid var(--color-border-2);
        border-bottom: 1px solid var(--color-border-2);

        .left-provider {
            height: 100%;
            display: grid;
            grid-template-rows: 1fr 60px;
            border-right: 1px solid var(--color-border-2);

            .provider-list {
                @include base.col-flex;
                gap: 0;

                .provider-item {
                    height: 50px;
                    width: 100%;
                    box-sizing: border-box;
                    padding-left: 10px;
                    border-bottom: 1px solid var(--color-border-1);
                    @include base.row-flex-start;
                    gap: 5px;
                    align-items: center;
                    transition: .15s all ease-in-out;
                    cursor: pointer;

                    .tit {
                        font-size: 14px;
                        font-weight: bold;
                    }

                    &:hover {
                        background: #eef9ee;
                    }

                    &.active {
                        background: #eef9ee;
                    }
                }

            }

            .add-privider {
                @include base.row-flex-center;

            }
        }

        .right-configuration {
            padding: 0 15px 15px;

            .top-switch {
                height: 50px;
                @include base.row-flex;
                align-items: center;
                justify-content: space-between;
                border-bottom: 1px solid var(--color-border-2);

                .switch-info {
                    @include base.row-flex-start;
                    align-items: center;
                    gap: 5px;

                    .tit {
                        font-size: 14px;
                        font-weight: bold;
                    }
                }
            }

            .center-api {
                padding-bottom: 15px;
                border-bottom: 1px solid var(--color-border-2);
            }

            .bottom-model-list {
                margin-top: 15px;

                .header-tit {
                    @include base.row-flex;
                    justify-content: space-between;
                    margin-bottom: 15px;

                    .tit {
                        font-weight: bold;
                        font-size: 14px;
                        margin-right: 5px;
                    }

                    .sub-tit {
                        color: var(--color-text-3);
                    }

                    .add-model {
                        @include base.row-flex-start;
                        gap: 5px;
                        align-items: center;
                        cursor: pointer;

                        .tit {
                            font-size: 14px;
                        }
                    }
                }

                .model-list {
                    @include base.col-flex;
                    justify-content: flex-start;


                    .model-item {
                        @include base.row-flex-start;
                        width: 100%;
                        align-items: center;
                        gap: 5px;
                        padding: 10px 5px;
                        transition: .15s all ease-in-out;
                        cursor: pointer;

                        i.active {
                            color: var(--color-primary-1);
                        }

                        .model-name {
                            font-weight: bold;
                            font-size: 14px;
                        }

                        &:hover {
                            background: #eef9ee;
                        }

                    }
                }
            }
        }
    }




</style>

<style>
    .customer-modal {
        .n-card__content {
            padding: 0 !important;
        }
    }
</style>