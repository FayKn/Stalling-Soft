<script setup lang="ts">
const props = defineProps<{
    date: string,
    data: Array<{ id: number, start: string, end: string, title: string, description: string }>
}>();
</script>

<template>
    <div class="bg-pink-500">
        {{date}}
    </div>
</template>

<style scoped>

</style>