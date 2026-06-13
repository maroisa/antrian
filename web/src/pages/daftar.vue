<script setup>
import AntrianHeader from "../components/AntrianHeader.vue";
import { PrinterIcon } from "@heroicons/vue/24/outline";
import { lastLoket } from "../utils/state.js";
import { get } from "../utils/api.js";
import { reactive, ref } from "vue";

let newAntrian = reactive({});

async function register(event) {
    event.preventDefault();
    const formData = new FormData(event.target);
    const loket = formData.get("loket");
    lastLoket.value = loket;

    const [res, err] = await get("loket", "baru");
    if (err) {
        console.log(err);
        return;
    }

    const json = await res.json();

    Object.assign(newAntrian, json);

    setTimeout(() => {
        window.print();
    }, 100);
}
</script>

<template>
    <AntrianHeader />
    <div class="p-2" id="printable-area">
        <p>SMK Negeri 7 Surakarta</p>
        <p>Tanggal {{ newAntrian.Tanggal }}</p>
        <h1 class="font-bold text-7xl">{{ newAntrian.Urut }}</h1>
    </div>
    <main class="flex justify-center items-center">
        <form class="w-full max-w-sm" v-on:submit="register">
            <button class="btn btn-primary w-full p-4" @click="">
                <PrinterIcon class="size-6" />
                <span>Cetak Nomor</span>
            </button>
        </form>
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
