<script setup>
import AntrianHeader from "../components/AntrianHeader.vue";
import { PrinterIcon } from "@heroicons/vue/24/outline";
import { lastLoket } from "../utils/state.js";
import { get } from "../utils/api.js";
import { reactive, ref, onMounted, nextTick } from "vue";
import { useRouter } from "vue-router";

let newAntrian = reactive({});

let isAuth = ref(false);
const router = useRouter();

onMounted(async () => {
    const [res, err] = await get("auth");
    if (err) {
        console.log(err.message);
        router.replace({ path: "/login" });
        await nextTick();
        return;
    }

    isAuth.value = true;
});

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
        <p class="text-center text-2xl font-bold">SPMB</p>
        <p class="text-center text-xl font-bold">SMK Negeri 7 Surakarta</p>
        <p class="text-center w-full">Tanggal {{ newAntrian.Tanggal }}</p>
        <h1 class="text-center font-bold text-7xl">
            {{ newAntrian.Urut }}
        </h1>
    </div>
    <template v-if="isAuth">
        <main class="flex justify-center items-center">
            <form class="w-full max-w-sm" v-on:submit="register">
                <button
                    class="btn btn-primary flex flex-col h-max w-max mx-auto text-xl gap-4 p-8"
                    @click=""
                >
                    <PrinterIcon class="size-8" />
                    <span>Cetak Nomor</span>
                </button>
            </form>
        </main>
    </template>
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
