<script setup>
import { onMounted, ref, reactive } from "vue";
import useWebSocket from "../utils/socket.js";

let synth = SpeechSynthesis;

const socket = ref(null);

const allLoket = reactive([]);

onMounted(async () => {
    if (typeof window !== "undefined" && "speechSynthesis" in window) {
        synth = window.speechSynthesis;
    }

    const [res, err] = await useWebSocket(socket, onMessage)
        .then((res) => [res, null])
        .catch((err) => [null, err]);

    if (err) {
        console.log(err);
        return;
    }

    socket.value = res;
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
    <main class="flex justify-center items-center">
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
