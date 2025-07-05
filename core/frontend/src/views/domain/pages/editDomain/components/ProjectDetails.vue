<template>
    <div class="content-wrapper">
        <n-card class="mb-5">
            <div class="page-tit">
                <div class="left-tit">
                    <div class="back-tool">
                        <i class="i-ri:apps-fill text-6"></i>
                    </div>
                    <span class="tit-content">
                        Project Details
                    </span>
                </div>
                <div class="right-tit">
                    <n-button type="primary" @click="updateProjectDetail(domain)">
                        <template #icon>
                            <i class="i-mingcute:save-2-line text-5"></i>
                        </template>
                        Save
                    </n-button>
                </div>
            </div>
        </n-card>
        <n-card class="mb-5">
            <n-form>
                <n-form-item>
                    <template #label><span class="form-label">Project Name</span></template>
                    <n-input v-model:value="project_name"></n-input>
                </n-form-item>
                <n-form-item>
                    <template #label><span class="form-label">Description</span></template>
                    <n-input type="textarea" :rows="7" v-model:value="description"></n-input>
                </n-form-item>
                <n-form-item>
                    <template #label><span class="form-label">Industry</span></template>
                    <n-select :options="industryOptions" v-model:value="industry"></n-select>
                </n-form-item>
            </n-form>
        </n-card>

        <n-card class="mb-5">
            <div class="page-tit mb-20px">
                <div class="left-tit">
                    <div class="back-tool">
                        <i class="i-ri:earth-fill text-6"></i>
                    </div>
                    <span class="tit-content">
                        Brand Logos
                    </span>
                </div>
            </div>

            <div class="sub-tit">
                <span class="tit">Logo System</span>
                <span>Upload your brand's logo system to maintain consistent branding across your project</span>
            </div>

            <div class="logo-operation">
                <div class="operation-item">
                    <div class="operation-tit">Primary Logo</div>
                    <div class="operation-sub-tit">Horizontal version for headers</div>
                    <div class="logo-area">
                        <div class="upload-logo-placeholder" v-if="!primary_logo">
                            <i class="i-carbon:cloud-upload text-8"></i>
                            <span>Primary Logo</span>
                        </div>
                        <n-image :src="primary_logo" class="w-100% h-100%" v-else>
                            <template #error>
                                <div class="w-100% h-100% flex justify-center items-center flex-col gap-1.25">
                                    <i class="i-mingcute:pic-line text-15"></i>
                                    <span>loading failed</span>
                                </div>
                            </template>
                        </n-image>
                    </div>
                </div>
                <div class="operation-item">
                    <div class="operation-tit">Secondary Logo</div>
                    <div class="operation-sub-tit">Vertical/stacked version</div>
                    <div class="logo-area">
                        <div class="upload-logo-placeholder" v-if="!secondary_logo">
                            <i class="i-carbon:cloud-upload text-8"></i>
                            <span>Primary Logo</span>
                        </div>
                        <n-image :src="secondary_logo" class="w-100% h-100%" v-else>
                            <template #error>
                                <div class="w-100% h-100% flex justify-center items-center flex-col gap-1.25">
                                    <i class="i-mingcute:pic-line text-15"></i>
                                    <span>loading failed</span>
                                </div>
                            </template>
                        </n-image>
                    </div>
                </div>
                <div class="operation-item">
                    <div class="operation-tit">Favicon</div>
                    <div class="operation-sub-tit">Icon for browser tabs</div>
                    <div class="logo-area">
                        <div class="upload-logo-placeholder" v-if="!favicon">
                            <i class="i-carbon:cloud-upload text-8"></i>
                            <span>Primary Logo</span>
                        </div>
                        <n-image :src="favicon" class="w-100% h-100%" v-else>
                            <template #error>
                                <div class="w-100% h-100% flex justify-center items-center flex-col gap-1.25">
                                    <i class="i-mingcute:pic-line text-15"></i>
                                    <span>loading failed</span>
                                </div>
                            </template>
                        </n-image>
                    </div>
                    <div class="operation-tools">
                        <div class="tool-item">
                            <i class="i-carbon:copy text-4"></i>
                        </div>
                        <div class="tool-item">
                            <i class="i-carbon:overflow-menu-vertical text-4"></i>
                        </div>
                        <div class="tool-item">
                            <i class="i-carbon:export text-4"></i>
                        </div>
                        <div class="tool-item">
                            <i class="i-material-symbols:delete-outline text-4"></i>
                        </div>
                    </div>
                </div>
            </div>
        </n-card>

        <n-card>
            <div class="page-tit">
                <div class="left-tit">
                    <div class="back-tool">
                        <i class="i-ri:earth-fill text-6"></i>
                    </div>
                    <span class="tit-content">
                        Knowledge Base
                    </span>
                </div>

                <div class="right-tit">
                    <n-button @click="knowledgeModalShow = true">New File</n-button>
                </div>
            </div>

            <n-card class="mt-5" v-for="(item, index) in knowledge_base" :key="index"
                style="background-color: var(--color-radio-1)">
                <div class="knowledge-box">
                    <div class="top-info">
                        <div class="info-left">
                            <span class="info-tit">{{ item.title }}</span>
                            <span class="info-sub-tit">11 days ago</span>
                        </div>
                        <div class="info-right" @click="handleDelete(item)">
                            <i class="i-material-symbols:close-rounded text-5"></i>
                        </div>
                    </div>
                    <div class="operation">
                        <n-button type="primary" ghost @click="openEditKnowledge(item)">Edit</n-button>
                        <n-button type="primary" ghost>Preiview</n-button>
                    </div>
                </div>
            </n-card>
        </n-card>
    </div>

    <!-- Add knowledge base -->
    <n-modal v-model:show="knowledgeModalShow" preset="card" draggable :close-on-esc="false" :mask-closable="false"
        title="New Knowledge File" class="w-150" :on-after-leave="closeKnowledgeModal">
        <n-form>
            <n-form-item label="File Name">
                <n-input v-model:value="knowledgeTitle"></n-input>
            </n-form-item>
            <n-form-item label="Content">
                <n-input type="textarea" :rows="8" v-model:value="knowledgeContent"></n-input>
            </n-form-item>
        </n-form>

        <template #footer>
            <div class="flex justify-end gap-5">
                <n-button @click="closeKnowledgeModal">Cancel</n-button>
                <n-button type="primary" @click="createOrUpdateKnowledge"
                    :disabled="!knowledgeTitle || !knowledgeContent">Create</n-button>
            </div>
        </template>
    </n-modal>

</template>

<script setup lang="ts">
    import { confirm } from '@/utils';
    import { getProjectDetail, updateProjectDetail, createKnowledgeBase, closeKnowledgeModal, openEditKnowledge, updateKnowledgeBase, deleteKnowledgeBase } from '../controller/projectDetail.controller';
    import { getEditDomainStoreData } from '../store';
    import { KnowledgeBase } from '../dto';
    const {
        project_name,
        description,
        industry,
        primary_logo,
        secondary_logo,
        favicon,
        knowledge_base,
        knowledgeModalShow,
        isEditKnowledge,
        knowledgeTitle,
        knowledgeContent
    } = getEditDomainStoreData()
    const route = useRoute()
    const industryOptions = ref([
        {
            label: "test",
            value: "test"
        }
    ])
    const domain = route.params.domain as string


    getProjectDetail(domain)
    /**
     * @description Create or update knowledge
     */
    function createOrUpdateKnowledge() {
        if (isEditKnowledge.value) {
            updateKnowledgeBase(domain)
        } else {
            createKnowledgeBase(domain)
        }
    }

    /**
     * @description Delete knowledge base
     */
    function handleDelete(knowledge: KnowledgeBase) {
        confirm({
            title: "Notice",
            content: "Are you sure you want to delete?",
            onConfirm() {
                deleteKnowledgeBase(domain, knowledge)
            }
        })
    }
</script>

<style scoped lang="scss">
    @use "@/styles/index" as base;
    @use "./mixin.scss" as mixin;

    .content-wrapper {
        @include mixin.content-wrapper;

        .page-tit {
            @include mixin.page-tit;
        }

        .sub-tit {
            @include base.col-flex-start;
            align-items: flex-start;
            gap: 10px;
            color: var(--color-text-2);
            margin-bottom: 40px;

            .tit {
                font-size: 14px;
                font-weight: bold;
            }
        }

        .logo-operation {
            display: grid;
            grid-template-columns: repeat(3, 1fr);
            gap: 20px;

            .operation-item {
                .operation-tit {
                    margin-bottom: 5px;
                    font-size: 14px;
                    font-weight: bold;
                }

                .operation-sub-tit {
                    margin-bottom: 15px;
                }

                .logo-area {
                    width: 100%;
                    height: 130px;
                    border: 1px solid var(--color-border-1);
                    border-radius: 16px;
                    margin-bottom: 10px;

                    .upload-logo-placeholder {
                        @include base.col-flex-center;
                        gap: 5px;
                        height: 100%;
                    }
                }

                .operation-tools {
                    @include base.row-flex-start;
                    gap: 10px;

                    .tool-item {
                        width: 34px;
                        height: 34px;
                        @include base.row-flex-center;
                        background: var(--color-bg-1);
                        border: 1px solid var(--color-border-1);
                        border-radius: 5px;
                        color: var(--color-text-2);
                        cursor: pointer;

                        &:hover {
                            background: var(--color-bg-4);
                        }
                    }
                }
            }
        }

        .form-label {
            @include mixin.form-label;
        }

        .knowledge-box {
            @include base.col-flex;
            justify-content: flex-start;
            align-items: flex-start;

            .top-info {
                width: 100%;
                @include base.row-flex;
                justify-content: space-between;
                align-items: flex-start;
                margin-bottom: 20px;

                .info-left {
                    @include base.col-flex;
                    justify-content: flex-start;
                    align-items: flex-start;

                    .info-tit {
                        font-size: 16px;
                        font-weight: bold;
                    }
                }

                .info-right {
                    @include base.row-flex-center;
                    align-items: center;
                    padding: 5px;
                    cursor: pointer;
                    transition: .15s all ease-in-out;

                    &:hover {
                        background: var(--color-border-2);
                    }
                }
            }

            .operation {
                @include base.row-flex-start;
                gap: 15px;
            }
        }
    }
</style>