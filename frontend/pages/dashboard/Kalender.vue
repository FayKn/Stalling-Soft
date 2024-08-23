<script setup lang="ts">
definePageMeta({
    layout: 'dashboard'
})

const calendarData = [
    {
        id: 1,
        date: "2024-08-18",
        data: [
            {
                id: 1,
                start: "2024-08-18T01:00:00",
                end: "2024-08-18T02:30:00",
                title: "Meeting 1",
                description: "Meeting with the team"
            },
            {
                id: 2,
                start: "2024-08-18T10:00:00",
                end: "2024-08-18T11:00:00",
                title: "Meeting 2",
                description: "Meeting with the team"
            }
        ]
    },
    {
        id: 2,
        date: "2024-08-19",
        data: null
    },
    {
        id: 3,
        date: "2024-08-20",
        data: null
    }
];



function formatHour(hour: number): string {
    return `${hour.toString().padStart(2, '0')}:00`;
}

function createEvent(hour: number, date: string) {
    console.log('Create event for ', date, ' at ', hour);
}
</script>

<template>
    <NuxtLayout>
        <div class="w-full mt-10 pb-10 min-h-[85vh]">
            <button class="bg-neutral-600 m-3 p-2 rounded text-white">
                + New Event
            </button>
            <div class="flex justify-center items-center">
                <div class="overflow-y-auto w-[96vw] max-h-[85vh]
                        grid grid-cols-[auto,repeat(3,1fr)]">
                    <!-- Hour Column -->
                    <div class="flex flex-col h-fit mt-[84px]">
                        <div class="flex flex-col h-[54px] mr-2" v-for="hour in 23" :key="hour">
                            <span class="text-white">{{ formatHour(hour) }}</span>
                        </div>
                    </div>
                    <!-- Calendar Rows -->
                    <template v-for="(date, count) in calendarData" :key="date.id">
                        <CalendarRow :class="{'border-l border-white': count !== 0  }" :date="date.date"
                                     :data="date.data" @createEvent="createEvent"/>
                    </template>
                </div>
            </div>
        </div>
    </NuxtLayout>
</template>