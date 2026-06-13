<script setup>
import { useRouter } from "vue-router";
import { post } from "../utils/api";

const router = useRouter();

async function handleForm(event) {
    event.preventDefault();
    const formData = new FormData(event.target);
    const [res, err] = await post("login", formData);
    if (err) {
        console.log(err.message);
        return;
    }

    router.replace("/");
}
</script>

<template>
    <main class="flex justify-center items-center">
        <form class="p-2" v-on:submit="handleForm">
            <fieldset
                class="fieldset bg-base-200 border-base-300 rounded-box w-xs border p-4"
            >
                <legend class="fieldset-legend">Login</legend>

                <label class="label">User</label>
                <input
                    required
                    name="nama"
                    type="text"
                    class="input"
                    placeholder="User"
                />

                <label class="label">Password</label>
                <input
                    required
                    name="password"
                    type="password"
                    class="input"
                    placeholder="Password"
                />

                <button type="submit" class="btn btn-primary mt-4">
                    Login
                </button>
            </fieldset>
        </form>
    </main>
</template>
