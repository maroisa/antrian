<script setup>
import { onMounted, ref, reactive, nextTick } from "vue";
import useWebSocket from "../utils/socket.js";
import { useRouter } from "vue-router";
import { get } from "../utils/api.js";

let synth = SpeechSynthesis;

const socket = ref(null);

const allLoket = reactive([]);

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

    if (typeof window !== "undefined" && "speechSynthesis" in window) {
        synth = window.speechSynthesis;
    }

    const [wsRes, wsErr] = await useWebSocket(socket, onMessage)
        .then((res) => [res, null])
        .catch((err) => [null, err]);

    if (wsErr) {
        console.log(wsErr.message);
        return;
    }

    socket.value = wsRes;
});

function onMessage(message) {
    console.log(message);
    Object.assign(allLoket, message.data ? message.data : []);
    // speak("Nomor urut " + message.Urut + " di loket " + message.Loket);
}

function speak(text) {
    if (!synth || !text) return;
    synth.cancel();
    const utterance = new SpeechSynthesisUtterance(text);
    const idVoice = synth
        .getVoices()
        .find((voice) => voice.lang.startsWith("id"));
    if (idVoice) {
        utterance.voice = idVoice;
    } else {
        utterance.lang = "id-ID";
    }

    utterance.onstart = () => setIsSpeaking(true);
    utterance.onend = () => setIsSpeaking(false);
    utterance.onerror = () => setIsSpeaking(false);

    console.log("Synth:", synth);
    console.log("Voice:", idVoice);
    console.log("Text:", text);
    synth.speak(utterance);
}
</script>

<template>
    <main class="flex justify-center items-center" v-if="isAuth">
        <div class="grid grid-cols-3 gap-4">
            <div class="border p-4" v-for="item in allLoket">
                <p class="text-center font-semibold mb-2">
                    Loket {{ item.Loket }}
                </p>
                <p>No. Urut {{ item.Urut }}</p>
            </div>
        </div>
    </main>
</template>
