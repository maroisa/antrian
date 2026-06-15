<script setup>
import AntrianHeader from "../components/AntrianHeader.vue";
import { nextTick, onMounted, ref } from "vue";
import { get } from "../utils/api";
import { useRouter } from "vue-router";
import { SpeakerWaveIcon } from "@heroicons/vue/24/outline";
import { SpeakerXMarkIcon } from "@heroicons/vue/24/outline";

let loket = ref(1);
let synth = SpeechSynthesis;
let data = ref([]);

let isAuth = ref(false);
let speakerActive = ref(true);

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

    if (typeof window !== "undefined" && "speechSynthesis" in window) {
        synth = window.speechSynthesis;
    }

    refresh();
});

async function refresh() {
    const [listAntrian, listAntrianErr] = await get(`loket/${loket.value}`);
    const listAntrianJson = await listAntrian.json();
    data.value = listAntrianJson;
}

function speak(value1, value2) {
    if (!synth) return;
    synth.cancel();
    const utterance = new SpeechSynthesisUtterance(
        "Antrian nomor " + value1 + " silahkan ke loket " + value2,
    );
    const idVoice = synth
        .getVoices()
        .find((voice) => voice.lang.startsWith("id"));
    if (idVoice) {
        utterance.voice = idVoice;
    } else {
        utterance.lang = "id-ID";
    }

    console.log("Synth:", synth);
    console.log("Voice:", idVoice);
    if (speakerActive.value) {
        synth.speak(utterance);
    }
}

async function mintaAntrian() {
    const [res, err] = await get("antrian", "minta");

    const json = await res.json();
    if (json.ID == 0) {
        alert("Tidak ada antrian baru");
        return;
    }

    const [antrianRes, antrianErr] = await get(
        `antrian/ambil/${loket.value}/${json.ID}`,
    );

    const antrianJson = await antrianRes.json();
    speak(json.Urut, loket.value);
    refresh();
}

async function selesaiAntrian(id) {
    const [res, err] = await get(`antrian/${id}/selesai`);
    refresh();
    console.log(res);
}
</script>

<template>
    <AntrianHeader />
    <template v-if="isAuth">
        <main class="p-4 max-w-4xl mx-auto">
            <div class="flex items-center justify-between w-full gap-4">
                <button
                    v-if="speakerActive"
                    class="btn btn-primary"
                    v-on:click="
                        () => {
                            speakerActive = !speakerActive;
                            synth.cancel();
                        }
                    "
                >
                    <SpeakerWaveIcon class="size-6" />
                </button>
                <button
                    v-else
                    class="btn btn-secondary"
                    v-on:click="() => (speakerActive = !speakerActive)"
                >
                    <SpeakerXMarkIcon class="size-6" />
                </button>
                <div class="flex items-center w-3xs sm:w-md">
                    <select
                        class="select"
                        v-on:change="
                            (e) => {
                                loket = e.target.value;
                                refresh();
                            }
                        "
                    >
                        <option value="1">Loket 1</option>
                        <option value="2">Loket 2</option>
                        <option value="3">Loket 3</option>
                        <option value="4">Loket 4</option>
                        <option value="5">Loket 5</option>
                        <option value="6">Loket 6</option>
                        <option value="7">Loket 7</option>
                        <option value="8">Loket 8</option>
                    </select>
                    <button class="btn btn-primary m-4" @click="mintaAntrian">
                        Panggil
                    </button>
                </div>
            </div>
            <table
                class="table table-zebra border-collapse border-black text-lg"
            >
                <thead>
                    <tr>
                        <th>No. Urut</th>
                        <th>Aksi</th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="item in data">
                        <td>{{ item.Urut }}</td>
                        <td class="flex gap-2">
                            <button
                                @click="() => speak(item.Urut, loket)"
                                class="btn btn-primary"
                            >
                                Panggil Ulang
                            </button>
                            <button
                                @click="() => selesaiAntrian(item.ID)"
                                class="btn btn-success"
                            >
                                Selesai
                            </button>
                        </td>
                    </tr>
                </tbody>
            </table>
        </main>
    </template>
</template>
