<script setup>
import { onMounted, ref } from "vue";
import AntrianHeader from "../../components/AntrianHeader.vue";

const isLoading = ref(true);
const loketData = ref([]);
const loketID = ref("0");

async function getLoket(loketID) {
    const API_URL = import.meta.env.VITE_API_URL;
    const res = await fetch(API_URL + "loket/" + loketID);
    const json = await res.json();

    loketData.value = json ? json : [];
    isLoading.value = false;
}
</script>

<template>
    <AntrianHeader />
    <main class="p-4 max-w-6xl mx-auto">
        <div>
            <select
                class="join-item select"
                v-on:change="
                    (e) => {
                        getLoket(e.target.value);
                        loketID = e.target.value;
                    }
                "
            >
                <option disabled selected>Pilih loket...</option>
                <option value="1">Loket 1</option>
                <option value="2">Loket 2</option>
                <option value="3">Loket 3</option>
                <option value="4">Loket 4</option>
            </select>
        </div>
        <div class="overflow-auto">
            <template v-if="!isLoading && loketData.length <= 0">
                <p class="flex mb-4 justify-center items-center italic text-xl">
                    Tidak ada antrian di Loket {{ loketID }}
                </p>
            </template>
            <table v-else class="table table-zebra">
                <thead>
                    <tr>
                        <th>No. Urut</th>
                        <th>Aksi</th>
                    </tr>
                </thead>
                <tbody></tbody>
            </table>
        </div>
    </main>
</template>
