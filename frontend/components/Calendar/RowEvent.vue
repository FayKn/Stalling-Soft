<script setup lang="ts">
const props = defineProps<{
    hour: number,
    data: {
        start: string,
        end: string,
        title: string
    }
}>();

const startFormatted = new Date(props.data.start).toLocaleTimeString('nl-NL', {
    hour: '2-digit',
    minute: '2-digit'
});

const endFormatted = new Date(props.data.end).toLocaleTimeString('nl-NL', {
    hour: '2-digit',
    minute: '2-digit'
});
</script>

<template>
    <div class="flex flex-row h-full">
        <div class="flex mr-2">
            <div>{{ startFormatted }}</div>
            -
            <div>{{ endFormatted }}</div>
        </div>
        <div>{{ props.data.title }}</div>
    </div>
</template>

<style scoped>

</style>