<script setup>
import { PrinterIcon } from "@heroicons/vue/24/outline";
import AntrianHeader from "../components/AntrianHeader.vue";
import { lastLoket } from "../utils/state.js";
import { get } from "../utils/api.js";

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

    alert("Antrian baru telah terdaftar pada loket " + loket);
}
</script>

<template>
    <AntrianHeader />
    <main class="flex justify-center items-center">
        <form class="w-full max-w-sm" v-on:submit="register">
            <select name="loket" class="select w-full mb-4">
                <option disabled :selected="lastLoket == 0">
                    Pilih loket...
                </option>
                <template v-for="v in Array(1, 2, 3, 4)">
                    <option :selected="lastLoket == v" :value="v">
                        Loket {{ v }}
                    </option>
                </template>
            </select>
            <button class="btn btn-primary w-full p-4">
                <PrinterIcon class="size-6" />
                <span>Cetak Nomor</span>
            </button>
        </form>
    </main>
</template>
