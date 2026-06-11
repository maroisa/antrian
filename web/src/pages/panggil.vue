<script setup>
import { onMounted, reactive, ref } from "vue";
import AntrianHeader from "../components/AntrianHeader.vue";
import { hasDipanggil, lastLoket } from "../utils/state";
import { SpeakerWaveIcon } from "@heroicons/vue/24/outline";
import { CheckBadgeIcon } from "@heroicons/vue/24/outline";
import { get } from "../utils/api";

const isLoading = ref(true);

const loketData = reactive([]);

onMounted(() => {
    if (lastLoket.value != 0) {
        refresh(lastLoket.value);
    }
});

async function refresh(loketID) {
    const [res, err] = await get("loket", loketID);
    if (err) {
        console.log(err);
        return;
    }

    const json = await res.json();
    lastLoket.value = loketID;
    Object.assign(loketData, json ? json : []);
}

async function handleAntrian(action, antrianItem) {
    if (action == "panggil") {
        const [res, err] = await get("antrian", antrianItem.ID, "panggil");
        if (err) console.log(err);

        hasDipanggil.value = true;
    } else if (action == "selesai") {
        const [res, err] = await get("antrian", antrianItem.ID, "selesai");
        if (err) {
            console.log(err);
            return;
        }

        const json = await res.json();
        alert("Antrian no. urut " + antrianItem.Urut + " telah selesai");
        refresh(lastLoket.value);
        hasDipanggil.value = false;
    }
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
        <select
            class="select mb-10"
            v-on:change="
                (e) => {
                    getLoket(e.target.value);
                    loketID = e.target.value;
                }
            "
        >
            <option disabled :selected="lastLoket == 0">Pilih loket...</option>
            <template v-for="v in Array(1, 2, 3, 4)">
                <option :selected="lastLoket == v" :value="v">
                    Loket {{ v }}
                </option>
            </template>
        </select>
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
                    <template v-if="loketData.length">
                        <tr>
                            <td>{{ loketData[0].Urut }}</td>
                            <td class="flex gap-4">
                                <button
                                    class="btn btn-info"
                                    @click="
                                        handleAntrian('panggil', loketData[0])
                                    "
                                >
                                    <SpeakerWaveIcon class="size-6" />
                                    <span>Panggil</span>
                                </button>
                                <button
                                    @click="
                                        handleAntrian('selesai', loketData[0])
                                    "
                                    class="btn btn-success"
                                    :disabled="!hasDipanggil"
                                >
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
