<template>
    <div class="p-6">
        <!-- Page tit -->
        <div class="page-tit">
            <div class="back-tool">
                <i class="i-ci:arrow-left-lg text-4"></i>
                <div class="back-btn">back</div>
            </div>
            <span class="tit-content">
                Domain / Edit domain
            </span>
        </div>

        <!-- Content tabs -->
         <div class="content-tabs">
            <div class="tab-item"><span>Project Details</span></div>
            <div class="tab-item"><span>Company Profile</span></div>
            <div class="tab-item"><span>Styling</span></div>
            <div class="tab-item"><span>Typography</span></div>
            <div class="tab-item"><span>Sitemap</span></div>
            <div class="tab-item"><span>Footer Settings</span></div>
            <div class="tab-item"><span>AI Settings</span></div>
         </div>
    </div>
</template>

<script setup lang="ts">

</script>

<style scoped lang="scss">
    @use "@/styles/index" as base;

    // page title
    .page-tit {
        @include base.row-flex-start;
        margin-bottom: 20px;
        gap: 15px;
        
        .back-tool {
            color: var(--color-primary-1);
            @include base.row-flex-start;
            gap: 5px;
            cursor: pointer;
            @include base.transition-all;
            &:hover{
                color: var(--color-primary-3);
            }

            .back-btn {
                font-size: 18px;

            }
        }

        .tit-content {
            font-size: 18px;
            color: #333;
            font-weight: bold;
        }
    }

    // Content tabs
    .content-tabs {
        width: 100%;
        height: 46px;
        border-bottom: 1px solid var(--color-border-1);
        @include base.row-flex-start;
        .tab-item {
            height: 100%;
            @include base.col-flex-center;
            font-size: 14px;
            padding: 0 15px;
            cursor: pointer;
            @include base.transition-all;

            &:hover {
                color: var(--color-primary-1);
            }
        }
    }
</style>