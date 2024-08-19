<script setup lang="ts">
const props = defineProps<{
    date: Date
}>();

const formattedDate = new Date(props.date).toLocaleDateString('nl-NL', {
    weekday: 'short',
    day: '2-digit'
}).replace(/^\w/, c => c.toUpperCase());

const [weekday, day] = formattedDate.split(' ');
const formattedOutput = `${day} ${weekday}`;

const isCurrentDay = () => {
    const currentDate = new Date();
    const curDateFormatted = currentDate.toLocaleDateString('nl-NL', {
        weekday: 'short',
        day: '2-digit'
    }).replace(/^\w/, c => c.toUpperCase());
    return curDateFormatted === formattedDate;
}
</script>

<template>
    <div class="p-2 text-white bg-neutral-600 text-center" :class="{'!bg-green-700':isCurrentDay()}">
        {{formattedOutput}}
    </div>
</template>

<style scoped>

</style>