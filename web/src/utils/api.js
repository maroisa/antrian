import { handleError } from "vue";
import { useRouter } from "vue-router";

const host = import.meta.env.VITE_API_HOST
    ? import.meta.env.VITE_API_HOST
    : window.location.host;

const URL = window.location.protocol + "//" + host + "/";

export async function get(...segments) {
    const [response, error] = await fetch(URL + segments.join("/"), {
        method: "GET",
        credentials: "include",
    })
        .then((res) => [res, null])
        .catch((err) => [null, err]);

    if (error) {
        return [null, error];
    }
    if (!response.ok) {
        return [null, new Error(response.statusText)];
    }
    return [response, error];
}

export async function post(path, body) {
    const [response, error] = await fetch(URL + path, {
        method: "POST",
        body: body,
        credentials: "include",
    })
        .then((res) => [res, null])
        .catch((err) => [null, err]);

    if (error) {
        return [null, error];
    }

    if (response && !response.ok) {
        return [null, new Error(response.statusText)];
    }
    return [response, error];
}
