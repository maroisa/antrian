<script setup>
import { onMounted, ref } from "vue";
import useWebSocket from "../utils/socket.js";

const socket = ref(null);

onMounted(async () => {
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
}
</script>

<template>
    <main class="flex justify-center items-center">
        <div class="grid grid-cols-3 gap-4">
            <div class="border p-4" v-for="item in Array(1, 2, 3, 4, 5, 6)">
                <p>Loket {{ item }}</p>
            </div>
        </div>
    </main>
</template>
