<template>
    <n-modal v-model:show="show" :close-on-esc="false" :mask-closable="false">
        <div class="flex justify-start gap-2.5 bg-transparent items-center" style="box-shadow: none;">
            <i class="i-svg-spinners:bars-scale text-[var(--color-warning-1)] text-6"></i>
            <span class="text-white flex items-center text-4">{{ $t("domain.waitAndCheckDomainStatus.loading") }}</span>
        </div>
    </n-modal>
</template>

<script setup lang="ts">
    import { instance } from '@/api';
    import { Message } from '@/utils';
    const { t } = useI18n()
    const show = ref(false)

    function open(domain: string) {
        checkStatus(domain)
        show.value = true
    }
    function close() {
        show.value = false
    }
    async function checkStatus(domain: string) {
        let timer = setInterval(async () => {
            const statusFlag = await getBaseInfoAndCheckStatus(domain)
            if (statusFlag) {
                close()
                clearInterval(timer)
                Message.success(t("domain.waitAndCheckDomainStatus.success"))
            }
        }, 1000)
    }

    /**
     * @description Get base info and check status
     */
    async function getBaseInfoAndCheckStatus(domain: string) {
        try {
            const res = await instance.post("/askai/project/get_base_info", { domain }) as Record<string, any>
            if (res.update_time) {
                return true
            } else {
                return false
            }
        } catch (error) {
            console.warn(error)
            Message.error(t("domain.waitAndCheckDomainStatus.error"))
            return false
        }
    }
    defineExpose({
        open,
        close,
        checkStatus
    })

</script>

<style scoped></style>