<template>
    <n-modal v-model:show="show" preset="card" draggable :close-on-esc="false" :mask-closable="false" segmented class="w-110"
        title="">
        <div class="wrapper">
            <div class="creation-methods">
                <div v-for="(item, index) in methodsList" :key="index"
                    :class="['choose-item', { active: choosedMethod == item }]" @click="choosedMethod = item">
                    <!-- <i
                        :class="[choosedMethod == item ? 'i-material-symbols:check-circle' : 'i-material-symbols:check-circle-outline', 'text-6']"></i> -->
                    <span class="item-label">{{ item }}</span>
                </div>
            </div>
            <div :class="['url-source', { hidden: choosedMethod !== 'AI' }]">
                <span class="label">Source Url</span>
                <n-input style="flex: 1;"></n-input>
            </div>
            <div class="desc">
                Use AI tools to help you automatically generate email content and improve your work efficiency.
            </div>
        </div>

        <template #footer>
            <div class="flex justify-end">
                <n-button type="primary" @click="jumpAndCreat">Create</n-button>
            </div>
        </template>
    </n-modal>
</template>

<script setup lang="ts">
    const router = useRouter()
    const methodsList = ref(["Drag", "AI", "HTML"])
    const choosedMethod = ref("")
    const show = ref(false)
    /**
     * @description open modal
     */
    function open() {
        show.value = true
    }

    /**
     * @description close modal
     */
    function close() {
        show.value = false
    }

    /**
     * @description jump
     */
    function jumpAndCreat() {
        router.push({ name: "ai-template" })
    }
    defineExpose({
        open,
        close
    })
</script>

<style scoped lang="scss">
    @use "@/styles/index" as base;

    .wrapper {
        @include base.col-flex-center;
        height: 153px;
        gap: 20px;

        .creation-methods {
            @include base.row-flex-center;
            border-bottom: 1px solid var(--color-border-1);
            width: 90%;
            gap: 25px;

            .choose-item {
                @include base.row-flex-start;
                gap: 5px;
                padding: 6px 15px;
                align-items: center;
                cursor: pointer;
                color: var(--color-text-2);
                transition: .1s all ease-in-out;
                box-sizing: border-box;

                .item-label {
                    font-size: 16px;
                    transition: .2s all ease-in-out;
                }

                &:hover {
                    color: var(--color-primary-1);
                    // transform: scale(1.3) translate(0, -5px);
                }

                &.active {
                    color: var(--color-primary-1);
                    transform: scale(1.3) translate(0, -5px);
                    font-weight: bold;
                    border-bottom: 2px solid var(--color-primary-1);
                }
            }
        }

        .url-source {
            width: 90%;
            @include base.row-flex-start;
            gap: 10px;

            &.hidden {
                visibility: hidden;
            }
        }

        .desc {
            width: 90%;
            font-size: 14px;
            color: var(--color-text-3);
        }
    }

</style>