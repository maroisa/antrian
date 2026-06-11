<script setup>
import NAV from "../nav.js";
</script>

<template>
    <header class="w-full flex justify-center gap-4 p-4">
        <template v-for="item in NAV">
            <RouterLink class="flex gap-2 btn nav-link text-" :to="item.to">
                <component class="size-6" :is="item.icon"></component>
                <span>{{ item.name }}</span>
            </RouterLink>
        </template>
    </header>
</template>
