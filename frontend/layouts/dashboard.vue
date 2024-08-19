<script setup lang="ts">

</script>

<template>
    <div>
        <div class="w-full bg-green-600 h-14 flex flex-row">
            <a href="/dashboard" class="h-full flex items-center text-white font-bold text-2xl mx-4">
                <img src="/logo.png" alt="Logo" class="h-full mx-4" />
            </a>
            <div class="flex flex-row items-center justify-between w-full
                        border-l border-black px-4 text-white font-medium">
                <div class="flex gap-4">
                    <DashboardNavbarActiveLink href="/dashboard/klanten" linkText="Klanten" />
                    <DashboardNavbarActiveLink href="/dashboard/facturen" linkText="Facturen" />
                    <DashboardNavbarActiveLink href="/dashboard/kalender" linkText="Agenda" />
                </div>
                <div class="flex gap-4">
                    <a href="dashboard/klanten">Instellingen</a>
                    <a href="dashboard/klanten">Uitloggen</a>
                </div>
            </div>
        </div>
        <slot />
    </div>
</template>

<style scoped>

</style>