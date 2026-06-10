<script setup>
import { onMounted, ref } from "vue";
import AntrianHeader from "../../components/AntrianHeader.vue";
import lastLoket from "../../utils/state";
import { SpeakerWaveIcon } from "@heroicons/vue/24/outline";
import { CheckBadgeIcon } from "@heroicons/vue/24/outline";

const isLoading = ref(true);
const loketData = ref([]);

onMounted(() => {
    if (lastLoket.value != 0) {
        refresh(lastLoket.value);
    }
});

async function refresh(loketID) {
    const API_URL = import.meta.env.VITE_API_URL;
    const res = await fetch(API_URL + "loket/" + loketID);
    const json = await res.json();
    lastLoket.value = loketID;
    loketData.value = json ? json : [];
}

async function getLoket(loketID) {
    refresh(loketID).then(() => {
        isLoading.value = false;
    });
}
</script>

<template>
    <AntrianHeader />
    <main class="p-4 w-full max-w-4xl mx-auto">
        <div>
            <select
                class="select"
                v-on:change="
                    (e) => {
                        getLoket(e.target.value);
                        loketID = e.target.value;
                    }
                "
            >
                <option disabled :selected="lastLoket == 0">
                    Pilih loket...
                </option>
                <template v-for="v in Array(1, 2, 3, 4)">
                    <option :selected="lastLoket == v" :value="v">
                        Loket {{ v }}
                    </option>
                </template>
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
                <tbody>
                    <template v-for="loketItem in loketData">
                        <tr>
                            <td>{{ loketItem.Urut }}</td>
                            <td class="flex gap-4">
                                <button class="btn btn-info">
                                    <SpeakerWaveIcon class="size-6" />
                                    <span>Panggil</span>
                                </button>
                                <button class="btn btn-success">
                                    <CheckBadgeIcon class="size-6" />
                                    <span>Selesai</span>
                                </button>
                            </td>
                        </tr>
                    </template>
                </tbody>
            </table>
        </div>
    </main>
</template>
