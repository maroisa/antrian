<script setup>
import { onMounted } from "vue";

const { tanggal, urut } = defineProps(["tanggal", "urut"]);

onMounted(() => {
    window.print();
});
</script>

<template>
    <main class="p-2" id="printable-area">
        <p>SMK Negeri 7 Surakarta</p>
        <p>Tanggal {{ tanggal }}</p>
        <h1 class="font-bold text-7xl">{{ urut }}</h1>
    </main>
</template>

<style>
#printable-area,
#printable-area * {
    display: none;
    visibility: hidden;
}

@media print {
    body * {
        visibility: hidden;
    }

    #printable-area,
    #printable-area * {
        display: block;
        visibility: visible;
    }

    #printable-area {
        position: fixed;
        left: 0;
        top: 0;
    }
}
</style>
