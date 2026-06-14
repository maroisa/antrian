<script setup>
import { nextTick, onMounted, ref } from "vue";
import NAV from "../nav";
import { get } from "../utils/api";
import { useRoute, useRouter } from "vue-router";

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
</script>

<template>
    <main v-if="isAuth" class="flex justify-center items-center">
        <div class="flex *:flex *:flex-col *:gap-2 gap-4">
            <template v-for="item in NAV">
                <RouterLink
                    :class="`last:col-span-2 btn h-max p-6 btn-${item.color}`"
                    :to="item.to"
                >
                    <component class="size-10" :is="item.icon"></component>
                    <span>{{ item.name }}</span>
                </RouterLink>
            </template>
        </div>
    </main>
</template>
