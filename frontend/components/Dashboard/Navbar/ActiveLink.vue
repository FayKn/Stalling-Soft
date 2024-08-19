<script setup lang="ts">
const isOnCurrentPage = (page: string) => {
    return useRoute().path === page;
}

defineProps({
    href: {
        type: String,
        required: true
    },
    linkText: {
        type: String,
        required: true
    }
})
</script>
<template>
    <a :class="{'underline underline-offset-4':isOnCurrentPage(href)}" :href=href>{{linkText}}</a>
</template>

<style scoped>

</style>