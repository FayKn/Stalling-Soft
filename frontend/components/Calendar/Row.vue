<script setup lang="ts">
const props = defineProps<{
    date: {
        type: String,
        required: true
    },
    data: {
        type: Array<{ id: number, start: string, end: string, title: string }>,
        required: true
    }
}>();

function getEventsForDay() {
    return props.data?.map((event: { id: number, start: string, end: string, title: string }) => {
        const eventStart = new Date(event.start);
        const eventEnd = new Date(event.end);
        const startHour = eventStart.getHours();
        const endHour = eventEnd.getHours();
        const duration = endHour - startHour;
        return {
            ...event,
            startHour,
            duration
        };
    });
}
</script>

<template>
    <div class="flex flex-col h-full">
        <CalendarRowHead :date="date"/>
        <div class="relative flex flex-col bg-neutral-500 border-b border-neutral-300 h-auto md:h-[1296px]"> <!-- 54px * 24 hours -->
            <template v-for="hour in 24" :key="hour">
                <div class="absolute w-full border-t border-neutral-300" :style="{ top: `${hour * 54}px` }"></div>
            </template>
            <template v-for="event in getEventsForDay()" :key="event.id">
                <div
                    class="absolute bg-green-700 text-white p-2"
                    :style="{ top: `${event.startHour * 54}px`, height: `${event.duration * 54}px` }"
                >
                    <CalendarRowEvent :hour="event.startHour" :data="event"/>
                </div>
            </template>
        </div>
    </div>
</template>

<style scoped>

</style>