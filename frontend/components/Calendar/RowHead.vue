<script lang="ts" setup>
const props = defineProps<{
    date: string
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
    <div :class="{'!bg-green-700':isCurrentDay()}" class="p-2 text-white bg-neutral-600 text-center">
        {{ formattedOutput }}
    </div>
</template>

<style scoped>

</style>