<template>
    <n-modal preset="card" draggable :close-on-esc="false" :mask-closable="false" title="Add Model" class="w-100"
        v-model:show="show">
        <n-form :model="addModelFormData" :rules="rules" ref="addModelFormRef">
            <n-form-item label="模型ID" path="modelId">
                <n-input v-model:value="addModelFormData.modelId"></n-input>
            </n-form-item>
            <n-form-item label="模型别名" path="title">
                <n-input v-model:value="addModelFormData.title"></n-input>
            </n-form-item>
            <n-form-item label="模型功能" path="capability">
                <n-select :options="modelCapabilities" multiple v-model:value="addModelFormData.capability"></n-select>
            </n-form-item>
            <n-form-item label="最大上下文长度" path="max_tokens">
                <n-input-number v-model:value="addModelFormData.max_tokens" class="w-100%"></n-input-number>
            </n-form-item>
        </n-form>
        <template #footer>
            <div class="flex justify-end gap-5">
                <n-button @click="close">Cancel</n-button>
                <n-button type="primary" @click="confirmAddModel(store)">Confirm</n-button>
            </div>
        </template>
    </n-modal>
</template>

<script setup lang="ts">
    import type { SelectOption } from "naive-ui"
    import { ModelStore } from "../dto"
    import { confirmAddModel } from "../controller"
    const store = inject<ModelStore>("modelStore")!
    const { addModelFormData, addModelFormRef } = store
    const show = ref(false)
    const modelCapabilities: SelectOption[] = [
        {
            label: "llm",
            value: "llm"
        },
        {
            label: "vision",
            value: "vision"
        },
        {
            label: "tools",
            value: "tools"
        },
        {
            label: "text-to-image",
            value: "text-to-image"
        },
    ]
    const rules = {
        modelId: [{ required: true, message: "modelId is required", trigger: "blur" }],
        title: [{ required: true, message: "title is required", trigger: "blur" }],
        capability: [{
            validator(_: any, value: any, callback: any) {
                if (value.length) {
                    return true
                } else {
                    return new Error("capability is required")
                }
            }, trigger: "blur"
        }],
        max_tokens: [{
            validator(_: any, value: any, callback: any) {
                if (value) {
                    return true
                } else {
                    return new Error("max_tokens is required")
                }
            }, trigger: "blur"
        }],
    }

    /***
     * @description Open modal
     */
    function open() {
        show.value = true
    }

    /**
     * @description Close  modal
     */
    function close() {
        show.value = false
        addModelFormData.value.capability = []
        addModelFormData.value.max_tokens = 0
        addModelFormData.value.modelId = ""
        addModelFormData.value.title = ""
    }

    defineExpose({
        open,
        close
    })
</script>

<style scoped></style>