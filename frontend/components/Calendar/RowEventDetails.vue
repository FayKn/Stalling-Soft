<script lang="ts" setup>
const props = defineProps<{
    date: string,
    data: Array<{ id: number, start: string, end: string, title: string, description: string }>
}>();
</script>

<template>
    <div class="bg-pink-500 absolute p-2">
        {{ date }}<br>
        {{ data.title }}<br>
        {{ data.description }}
    </div>
</template>

<style scoped>

</style>