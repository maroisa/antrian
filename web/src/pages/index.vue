<script setup>
import { PlusCircleIcon } from "@heroicons/vue/20/solid";
import { ComputerDesktopIcon } from "@heroicons/vue/24/solid";
import { SpeakerWaveIcon } from "@heroicons/vue/24/solid";
import { PlusIcon } from "@heroicons/vue/24/solid";
</script>

<template>
    <main class="flex justify-center items-center">
        <div class="grid grid-cols-2 *:flex *:flex-col *:gap-2 gap-4">
            <RouterLink class="btn h-max p-6 btn-primary" to="/daftar">
                <PlusCircleIcon class="size-10" />
                <span>Daftar</span>
            </RouterLink>
            <RouterLink class="btn h-max p-6 btn-secondary" to="/panggil">
                <SpeakerWaveIcon class="size-10" />
                <span>Panggil</span>
            </RouterLink>
            <RouterLink class="btn h-max p-6 btn-neutral col-span-2" to="/view">
                <ComputerDesktopIcon class="size-10" />
                <span>Tampilan</span>
            </RouterLink>
        </div>
    </main>
</template>
