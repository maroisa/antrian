<script setup>
import { onMounted, ref } from 'vue';

let synth = SpeechSynthesis;

let angka = ref("")
let loket = ref()

onMounted(() => {
    if (typeof window !== "undefined" && "speechSynthesis" in window) {
        synth = window.speechSynthesis;
    }
})

function panggil() {
    let a = angka.value.replace(/[^0-9]/g, '')
    if (!a) return
    speak(a, loket.value)
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
    synth.speak(utterance);
}

</script>

<template>
    <main class="flex justify-center items-center">
        <div class="p-4 flex gap-4" >
            <input v-model="angka" class="input" placeholder="Masukkan angka"></input>
                <select
                v-model="loket"
                    class="select"
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
                <button @click="panggil" class="btn btn-primary">Panggil</button>
        </div>
    </main>
</template>
