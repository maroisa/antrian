<script setup>
import AntrianHeader from "../components/AntrianHeader.vue";
import Cetak from "../components/Cetak.vue";
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

    const [res, err] = await get("loket", loket, "baru");
    if (err) {
        console.log(err);
        return;
    }

    const json = await res.json();

    Object.assign(newAntrian, json);
}
</script>

<template>
    <AntrianHeader />
    <template v-if="Object.keys(newAntrian).length">
        <Cetak :urut="newAntrian.Urut" :tanggal="newAntrian.Tanggal" />
    </template>
    <main class="flex justify-center items-center">
        <form class="w-full max-w-sm" v-on:submit="register">
            <button class="btn btn-primary w-full p-4" @click="">
                <PrinterIcon class="size-6" />
                <span>Cetak Nomor</span>
            </button>
        </form>
    </main>
</template>
