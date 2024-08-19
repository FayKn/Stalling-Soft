<script setup lang="ts">
definePageMeta({
    layout: 'dashboard'
})

const klantenData = [
    {
        id: 1,
        name: "Klant 1",
        dateSent: "01-01-2021",
        amount: "€100",
        paymentDate: "01-02-2021",
        paymentStatus: "Betaald"
    },
    {
        id: 2,
        name: "Klant 2",
        dateSent: "01-01-2021",
        amount: "€100",
        paymentDate: "01-02-2021",
        paymentStatus: "Betaald"
    },
    {
        id: 3,
        name: "Klant 3",
        dateSent: "01-01-2021",
        amount: "€100",
        paymentDate: "01-02-2021",
        paymentStatus: "Betaald"
    },
    {
        id: 4,
        name: "Klant 4",
        dateSent: "01-01-2021",
        amount: "€100",
        paymentDate: "01-02-2021",
        paymentStatus: "Betaald"
    },
]
</script>

<template>
    <NuxtLayout>
        <div class="grid grid-cols-5 mx-[10%] text-white border rounded-md mt-10">
            <div class="tableCol"> Klant</div>
            <div class="tableCol"> Datum verzonden</div>
            <div class="tableCol"> Bedrag</div>
            <div class="tableCol"> Uiterlijke datum betaling</div>
            <div class="tableCol"> Status betaling</div>

            <template v-for="klant in klantenData">
                <a :href="'/dashboard/klanten/'+klant.id" class="tableCol"> {{ klant.name }} </a>
                <a :href="'/dashboard/klanten/'+klant.id" class="tableCol"> {{ klant.dateSent }} </a>
                <a :href="'/dashboard/klanten/'+klant.id" class="tableCol"> {{ klant.amount }} </a>
                <a :href="'/dashboard/klanten/'+klant.id" class="tableCol"> {{ klant.paymentDate }} </a>
                <a :href="'/dashboard/klanten/'+klant.id" class="tableCol"> {{ klant.paymentStatus }} </a>
            </template>
        </div>
    </NuxtLayout>
</template>

<style scoped>
.tableCol {
    @apply border-l border-b pl-2;
}
</style>