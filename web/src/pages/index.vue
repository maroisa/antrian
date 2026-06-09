<script setup>
import { PlusCircleIcon } from "@heroicons/vue/20/solid";
import { ComputerDesktopIcon } from "@heroicons/vue/24/solid";
import { SpeakerWaveIcon } from "@heroicons/vue/24/solid";
import NAV from "../nav";
</script>

<template>
    <main class="flex justify-center items-center">
        <div class="grid grid-cols-2 *:flex *:flex-col *:gap-2 gap-4">
            <template v-for="item in NAV">
                <RouterLink
                    :class="`last:col-span-2 btn h-max p-6 btn-${item.color}`"
                    :to="item.to"
                >
                    <component class="size-10" :is="item.icon"></component>
                    <span>{{ item.name }}</span>
                </RouterLink>
            </template>
        </div>
    </main>
</template>
