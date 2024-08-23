<script setup lang="ts">
const props = defineProps<{
    date: string,
    data: Array<{ id: number, start: string, end: string, title: string, description: string }>
}>();

function getEventsForDay() {
    return props.data?.map((event: { id: number, start: string, end: string, title: string }) => {
        const eventStart = new Date(event.start);
        const eventEnd = new Date(event.end);
        const startHour = eventStart.getHours();
        const startMinute = eventStart.getMinutes();
        const endHour = eventEnd.getHours();
        const endMinute = eventEnd.getMinutes();
        const durationInMinutes = (endHour * 60 + endMinute) - (startHour * 60 + startMinute);
        return {
            ...event,
            startHour,
            startMinute,
            durationInMinutes
        };
    });
}
</script>

<template>
    <div class="flex flex-col h-full">
        <CalendarRowHead :date="date"/>
        <div class="relative flex flex-col bg-neutral-500 border-b border-neutral-300 h-auto md:h-[1296px]"> <!-- 54px * 24 hours -->
            <template v-for="hour in 24" :key="hour">
                <div class="absolute w-full border-t border-neutral-300 min-h-[54px] min-w-[200px]" :style="{ top: `${hour * 54}px` }"
                     @click="$emit('createEvent', hour, date)"></div>
            </template>
            <template v-for="event in getEventsForDay()" :key="event.id">
                <div
                    class="absolute bg-green-700 text-white p-[1%] w-[98%] rounded-xl"
                    :style="{ top: `${event.startHour * 54 + (event.startMinute / 60) * 54}px`, height: `${(event.durationInMinutes / 60) * 54}px` }"
                >
                    <CalendarRowEvent :hour="event.startHour" :data="event"/>
                    <CalendarRowEventDetails :date="date" :data="event"/>
                </div>
            </template>
        </div>
    </div>
</template>

<style scoped>

</style>