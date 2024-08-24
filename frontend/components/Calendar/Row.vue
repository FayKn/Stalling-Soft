<script lang="ts" setup>
const props = defineProps<{
    selectedEvent: number,
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

function getWidthOfElementById(id: number) {
    // parse the id to a string
    const idStr = id.toString();
    const element = document.getElementById(idStr);
    if (element) {
        return element.offsetWidth;
    }
    return 0;
}
</script>

<template>
    <div class="flex flex-col h-full">
        <CalendarRowHead :date="date"/>
        <div class="relative flex flex-col bg-neutral-500 border-b border-neutral-300 h-auto md:h-[1296px]">
            <!-- 54px * 24 hours -->
            <template v-for="hour in 24" :key="hour">
                <div :style="{ top: `${hour * 54}px` }"
                     class="absolute w-full border-t border-neutral-300 min-h-[54px] min-w-[200px]"
                     @click="$emit('createEvent', hour, date)"/>
            </template>
            <template v-for="event in getEventsForDay()" :key="event.id">
                <div
                    :id="event.id"
                    :style="{ top: `${event.startHour * 54 + (event.startMinute / 60) * 54}px`, height: `${(event.durationInMinutes / 60) * 54}px` }"
                    class="absolute bg-green-700 text-white p-[1%] w-[98%] rounded-xl"
                    @click="$emit('setSelectedEvent', event.id)"
                >
                    <CalendarRowEvent :data="event" :hour="event.startHour"/>
                </div>
                <ClientOnly>
                    <!-- Move the event left of the event by its whole width -->
                    <CalendarRowEventDetails
                        v-if="selectedEvent === event.id"
                        ref="openEvent"
                        :data="event" :date="date"
                        :style="{ top: `${event.startHour * 54 + (event.startMinute / 60) * 54}px`, transform: `translateX(${getWidthOfElementById(event.id)+12}px)`}"
                        class="z-10"
                    />
                </ClientOnly>
            </template>
        </div>
    </div>
</template>

<style scoped>

</style>